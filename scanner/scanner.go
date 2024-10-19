package scanner

import (
	"bufio"
	"go-cloc/logger"
	"io"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

type FileScanResults struct {
	FilePath          string
	TotalLines        int
	CodeLineCount     int
	BlankLineCount    int
	CommentsLineCount int
}
type AnalyzeLineResult string

const (
	Code      AnalyzeLineResult = "code"
	Comment   AnalyzeLineResult = "comment"
	BlankLine AnalyzeLineResult = "blankline"
)

func AnalyzeLine(line string, languageInfo LanguageInfo, isInBlockComment bool) (AnalyzeLineResult, bool) {

	if isInBlockComment || beginsWithFirstMultiLineComment(line, languageInfo) {
		// last characters on the line end a multi-line comment
		if endsWithSecondMultiLineComment(line, languageInfo) {
			return Comment, false
		}
		// the multi-line comment continues
		if !hasSecondMultiLineComment(line, languageInfo) {
			return Comment, true
		} else {
			// end of a mult-line comment is within the same line, therefore lines of code could be after this
			splitLine := splitLineByFirstMultiLineComment(line, languageInfo)[1]
			return AnalyzeLine(splitLine, languageInfo, false)
		}

		// if not, then it could be code, need to recurse in

	} else if isBlankLine(line) {
		return BlankLine, false
	} else if hasSingleLineComment(line, languageInfo) {
		return Comment, false
	}
	// it must be code
	return Code, false

}

func ScanFile(filePath string) FileScanResults {
	commentsLineCount := 0
	codeLineCount := 0
	blankLineCount := 0
	totalLines := 0

	f, err := os.Open(filePath)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	// Get metadata about file
	fileName := f.Name()
	suffix := ParseFileSuffix(fileName)
	_, languageInfo, found := LookupByExtension(suffix)
	// If not supported return 0s, TODO should probably throw an error or report on it
	// TODO Dockerfile does not always have an extension, but we will count it
	if !found {
		logger.Debug("Skipping file: ", fileName, " suffix '", suffix, "' is not supported.")
		return FileScanResults{
			BlankLineCount:    0,
			CodeLineCount:     0,
			CommentsLineCount: 0,
			TotalLines:        0,
		}
	}

	// Scan file
	reader := bufio.NewReader(f)
	isInBlockComment := false
	debugLineNum := 1
	for {

		line, err := reader.ReadString('\n')
		line = strings.TrimSpace(line)

		lineResult, blockCommentContinuesToNexLine := AnalyzeLine(line, languageInfo, isInBlockComment)
		if lineResult == Code {
			isInBlockComment = false
			codeLineCount++
		} else if lineResult == BlankLine {
			isInBlockComment = false
			blankLineCount++
		} else if lineResult == Comment {
			isInBlockComment = blockCommentContinuesToNexLine
			commentsLineCount++
		}

		if err != nil {
			// reached end of file
			if err == io.EOF {
				break
			}
			log.Fatalln(err)
		}
		debugLineNum++
	}

	// return the totals
	result := FileScanResults{}
	result.TotalLines = totalLines
	result.CodeLineCount = codeLineCount
	result.BlankLineCount = blankLineCount
	result.CommentsLineCount = commentsLineCount
	result.FilePath = filePath
	return result

}

/*
*
@singleLineCommentPrefix is something "/" or "//" or "#"
*/
func hasSingleLineComment(line string, languageInfo LanguageInfo) bool {
	for _, singleLineCommentPrefix := range languageInfo.LineComments {
		return strings.HasPrefix(line, singleLineCommentPrefix)
	}
	return false
}

func beginsWithFirstMultiLineComment(line string, languageInfo LanguageInfo) bool {
	for _, pair := range languageInfo.MultiLineComments {
		firstMultiLineCommentToken := pair[0]
		return strings.HasPrefix(line, firstMultiLineCommentToken)
	}
	return false
}

func hasSecondMultiLineComment(line string, languageInfo LanguageInfo) bool {
	for _, pair := range languageInfo.MultiLineComments {
		secondMultiLineCommentToken := pair[1]
		return strings.Contains(line, secondMultiLineCommentToken)
	}
	return false
}
func endsWithSecondMultiLineComment(line string, languageInfo LanguageInfo) bool {
	for _, pair := range languageInfo.MultiLineComments {
		secondMultiLineCommentToken := pair[1]
		return strings.HasSuffix(line, secondMultiLineCommentToken)
	}
	return false
}
func splitLineByFirstMultiLineComment(line string, languageInfo LanguageInfo) []string {
	for _, pair := range languageInfo.MultiLineComments {
		secondMultiLineCommentToken := pair[1]
		return strings.SplitN(line, secondMultiLineCommentToken, 2)
	}
	return []string{}
}

func isBlankLine(line string) bool {
	return len(line) == 0
}

func loadIgnorePatterns(patterns []string) []*regexp.Regexp {
	var regexps []*regexp.Regexp
	for _, pattern := range patterns {
		// got *.json
		// "^(.*\.json)$"
		newPattern := "^" + strings.ReplaceAll(strings.ReplaceAll(pattern, ".", "\\."), "*", ".*") + "$"
		regexps = append(regexps, regexp.MustCompile(newPattern))
		logger.Debug("Adding pattern " + newPattern)
	}
	return regexps
}

// ReadIgnoreFile reads a file specified by the given path and returns a slice of strings
// containing non-empty, trimmed lines from the file. It logs the file path being read and
// exits the program if an error occurs while reading the file.
//
// Parameters:
//   - path: The file path to read.
//
// Returns:
//   - A slice of strings containing the non-empty, trimmed lines from the file.
func ReadIgnoreFile(path string) []string {
	logger.Debug("Reading ignore file ", path)
	data, err := os.ReadFile(path)
	if err != nil {
		logger.LogStackTraceAndExit(err)
	}

	// Split the file content by new lines and trim spaces
	lines := strings.Split(string(data), "\n")
	var ignoreList []string
	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if trimmed != "" { // Ignore empty lines
			ignoreList = append(ignoreList, trimmed)
		}
	}

	return ignoreList
}

func ParseFileSuffix(fileName string) string {
	splitArr := strings.Split(fileName, ".")
	// if file does not have a suffix
	if len(splitArr) != 0 {
		suffix := splitArr[len(splitArr)-1]
		return "." + strings.ToLower(suffix)
	}
	return ""
}

func WalkDirectory(targetPath string, ignorePatterns []string) []string {
	patterns := loadIgnorePatterns(ignorePatterns)

	// Store the current working directory
	originalDir, err := os.Getwd()
	if err != nil {
		logger.Error("Error getting current directory:", err)
	}

	logger.Debug("Target directory is ", originalDir)
	var fileNames []string
	err = filepath.WalkDir(targetPath, func(path string, info os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		for _, pattern := range patterns {
			if pattern.Match([]byte(path)) {
				if info.IsDir() {
					logger.Debug("SKipping DIR - " + info.Name())
					return filepath.SkipDir
				}
				return nil
			}
		}
		if !info.IsDir() {
			suffix := ParseFileSuffix(info.Name())
			_, _, found := LookupByExtension(suffix)

			if found {
				fileNames = append(fileNames, path)
			} else {
				logger.Debug("Skipping file - ", path, " suffix - ", suffix, " - not supported")
			}
			return nil
		}
		return err
	})
	if err != nil {
		log.Fatalln(err)
	}

	// Change back to the original directory
	err = os.Chdir(originalDir)
	if err != nil {
		logger.Debug("Error changing back to the original directory:", err)
	}

	return fileNames
}
