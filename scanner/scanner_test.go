package scanner

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_scanner_ScanFile_javascript_easy(t *testing.T) {
	result := ScanFile("test-files/js/easy.js")

	// Assert
	assert.Equal(t, 16, result.CodeLineCount)
	assert.Equal(t, 9, result.CommentsLineCount)
	assert.Equal(t, 7, result.BlankLineCount)
}
func Test_scanner_ScanFile_javascript_hard(t *testing.T) {
	result := ScanFile("test-files/js/hard.js")

	// Assert
	assert.Equal(t, 16, result.CodeLineCount)
	assert.Equal(t, 27, result.CommentsLineCount)
	assert.Equal(t, 8, result.BlankLineCount)
}

func Test_scanner_ScanFile_c_hard(t *testing.T) {
	result := ScanFile("test-files/c/hard.c")

	// Assert
	assert.Equal(t, 780, result.CodeLineCount)
	assert.Equal(t, 418, result.CommentsLineCount)
	assert.Equal(t, 135, result.BlankLineCount)

}
func Test_scanner_ScanFile_c_evil(t *testing.T) {
	result := ScanFile("test-files/c/evil.c")

	// Assert
	assert.Equal(t, 475, result.CodeLineCount)
	assert.Equal(t, 75, result.CommentsLineCount)
	assert.Equal(t, 21, result.BlankLineCount)

}

func Test_scanner_ScanFile_cpp_hard(t *testing.T) {
	result := ScanFile("test-files/cpp/hard.cpp")

	// Assert
	assert.Equal(t, 954, result.CodeLineCount)
	assert.Equal(t, 8, result.CommentsLineCount)
	assert.Equal(t, 38, result.BlankLineCount)
}

// this file is weird because some multi-line comments end with a \, these are counted as code
func Test_scanner_ScanFile_cpp_evil(t *testing.T) {
	result := ScanFile("test-files/cpp/evil.cpp")

	// Assert
	assert.Equal(t, 7158, result.CodeLineCount)
	assert.Equal(t, 2743, result.CommentsLineCount)
	assert.Equal(t, 1548, result.BlankLineCount)

}

func Test_scanner_AnalyzeLine_hard(t *testing.T) {
	testStr := "/* GFLOPS 3.398 x 20 = 67.956 */ {{7, 7}, {{1, 128, 46, 46}}, 128, 1, {1, 1}, {1, 1}, {3, 3}, {0, 0}, \"\", true, 3397788160.},"
	_, languageInfo, _ := LookupByExtension(".cpp")
	result, _ := AnalyzeLine(testStr, languageInfo, false)

	// Assert
	assert.Equal(t, Code, result)

}

func Test_scanner_ScanFile_binary(t *testing.T) {
	result := ScanFile("test-files/misc/test.bin")

	// Assert
	assert.Equal(t, 0, result.CodeLineCount)

}
func Test_scanner_ScanFile_blank_file(t *testing.T) {
	result := ScanFile("test-files/misc/blank-file.js")

	// Assert
	assert.Equal(t, 0, result.CodeLineCount)

}

func Test_scanner_ScanFile_pdf(t *testing.T) {
	result := ScanFile("test-files/misc/sample.pdf")

	// Assert
	assert.Equal(t, 0, result.CodeLineCount)

}

func Test_scanner_ScanFile_massive_line_yaml(t *testing.T) {
	result := ScanFile("test-files/misc/massive-line.yaml")

	// Assert
	assert.Equal(t, 1, result.CodeLineCount)
}
func Test_scanner_ScanFile_minified_line_txt(t *testing.T) {
	result := ScanFile("test-files/misc/minified.js")

	// Assert
	assert.Equal(t, 1, result.CodeLineCount)
}
func Test_scanner_ParseFileSuffix(t *testing.T) {
	suffix := ParseFileSuffix("main.js")

	// Assert
	assert.Equal(t, ".js", suffix)

	suffix = ParseFileSuffix("something.typescript.JCL")
	assert.Equal(t, ".jcl", suffix)
}
func Test_scanner_WalkDirectory_no_ignores(t *testing.T) {
	ignorePatterns := []string{}

	result := WalkDirectory("test-files/js", ignorePatterns)

	// Assert
	assert.Equal(t, 2, len(result))
}

func Test_scanner_WalkDirectory_with_ignores(t *testing.T) {
	ignorePatterns := []string{"*easy.js"}

	result := WalkDirectory("test-files/js", ignorePatterns)

	// Assert
	assert.Equal(t, 1, len(result))
}

func Test_scanner_ReadIgnoreFile(t *testing.T) {

	result := ReadIgnoreFile("test-files/test-ignore-file.txt")
	fmt.Println(result)
	// Assert
	assert.Equal(t, 2, len(result))
	assert.Equal(t, "*.js", result[0])
	assert.Equal(t, "misc/", result[1])
}
