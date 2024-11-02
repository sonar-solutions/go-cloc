package scanner

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go-cloc/logger"
	"io"
	"os"
)

type LanguageInfo struct {
	LineComments      []string   `json:"LineComments"`
	MultiLineComments [][]string `json:"MultiLineComments"`
	Extensions        []string   `json:"Extensions"`
	FileNames         []string   `json:"FileNames"`
}

var Languages = map[string]LanguageInfo{
	"ActionScript": {
		LineComments:      []string{"//"},
		MultiLineComments: [][]string{{"/*", "*/"}},
		Extensions:        []string{".as"},
		FileNames:         []string{},
	},
	"Abap": {
		LineComments:      []string{"\""},
		MultiLineComments: [][]string{{"/*", "*/"}},
		Extensions:        []string{".abap", ".ab4", ".flow"},
		FileNames:         []string{},
	},
	"Apex": {
		LineComments:      []string{"//"},
		MultiLineComments: [][]string{{"/*", "*/"}},
		Extensions:        []string{".cls", ".trigger"},
		FileNames:         []string{},
	},
	"C": {
		LineComments:      []string{"//"},
		MultiLineComments: [][]string{{"/*", "*/"}},
		Extensions:        []string{".c"},
		FileNames:         []string{},
	},
	"C Header": {
		LineComments:      []string{"//"},
		MultiLineComments: [][]string{{"/*", "*/"}},
		Extensions:        []string{".h"},
		FileNames:         []string{},
	},
	"C++": {
		LineComments:      []string{"//"},
		MultiLineComments: [][]string{{"/*", "*/"}},
		Extensions:        []string{".cpp", ".cc", ".cxx", ".c++"},
		FileNames:         []string{},
	},
	"C++ Header": {
		LineComments:      []string{"//"},
		MultiLineComments: [][]string{{"/*", "*/"}},
		Extensions:        []string{".hh", ".hpp", ".hxx", ".h++", ".ipp"},
		FileNames:         []string{},
	},
	"COBOL": {
		LineComments:      []string{"*", "/"},
		MultiLineComments: [][]string{},
		Extensions:        []string{".cbl", ".ccp", ".cob", ".cobol", ".cpy"},
		FileNames:         []string{},
	},
	"C#": {
		LineComments:      []string{"//"},
		MultiLineComments: [][]string{{"/*", "*/"}},
		Extensions:        []string{".cs"},
		FileNames:         []string{},
	},
	"CSS": {
		LineComments:      []string{"//"},
		MultiLineComments: [][]string{{"/*", "*/"}},
		Extensions:        []string{".css"},
		FileNames:         []string{},
	},
	"Golang": {
		LineComments:      []string{"//"},
		MultiLineComments: [][]string{{"/*", "*/"}},
		Extensions:        []string{".go"},
		FileNames:         []string{},
	},
	"HTML": {
		LineComments:      []string{},
		MultiLineComments: [][]string{{"<!--", "-->"}},
		Extensions:        []string{".html", ".htm", ".cshtml", ".vbhtml", ".aspx", ".ascx", ".rhtml", ".erb", ".shtml", ".shtm", ".cmp"},
		FileNames:         []string{},
	},
	"Java": {
		LineComments:      []string{"//"},
		MultiLineComments: [][]string{{"/*", "*/"}},
		Extensions:        []string{".java", ".jav"},
		FileNames:         []string{},
	},
	"JavaScript": {
		LineComments:      []string{"//"},
		MultiLineComments: [][]string{{"/*", "*/"}},
		Extensions:        []string{".js", ".jsx", ".jsp", ".jspx", ".jspf", ".mjs"},
		FileNames:         []string{},
	},
	"Kotlin": {
		LineComments:      []string{"//"},
		MultiLineComments: [][]string{{"/*", "*/"}},
		Extensions:        []string{".kt", ".kts"},
		FileNames:         []string{},
	},
	"Flex": {
		LineComments:      []string{"//"},
		MultiLineComments: [][]string{{"/*", "*/"}},
		Extensions:        []string{".as"},
		FileNames:         []string{},
	},
	"PHP": {
		LineComments:      []string{"//", "#"},
		MultiLineComments: [][]string{{"/*", "*/"}},
		Extensions:        []string{".php", ".php3", ".php4", ".php5", ".phtml", ".inc"},
		FileNames:         []string{},
	},
	"Objective-C": {
		LineComments:      []string{"//"},
		MultiLineComments: [][]string{{"/*", "*/"}},
		Extensions:        []string{".m"},
		FileNames:         []string{},
	},
	"Oracle PL/SQL": {
		LineComments:      []string{"--"},
		MultiLineComments: [][]string{{"/*", "*/"}},
		Extensions:        []string{".pkb"},
		FileNames:         []string{},
	},
	"PL/I": {
		LineComments:      []string{"--"},
		MultiLineComments: [][]string{{"/*", "*/"}},
		Extensions:        []string{".pl1"},
		FileNames:         []string{},
	},
	"Python": {
		LineComments:      []string{"#"},
		MultiLineComments: [][]string{{"\"\"\"", "\"\"\""}},
		Extensions:        []string{".py", ".python", ".ipynb"},
		FileNames:         []string{},
	},

	"RPG": {
		LineComments:      []string{"#"},
		MultiLineComments: [][]string{},
		Extensions:        []string{".rpg"},
		FileNames:         []string{},
	},
	"Ruby": {
		LineComments:      []string{"#"},
		MultiLineComments: [][]string{{"=begin", "=end"}},
		Extensions:        []string{".rb"},
		FileNames:         []string{},
	},
	"Scala": {
		LineComments:      []string{"//"},
		MultiLineComments: [][]string{{"/*", "*/"}},
		Extensions:        []string{".scala"},
		FileNames:         []string{},
	},
	"Scss": {
		LineComments:      []string{"//"},
		MultiLineComments: [][]string{{"/*", "*/"}},
		Extensions:        []string{".scss"},
		FileNames:         []string{},
	},
	"SQL": {
		LineComments:      []string{"--"},
		MultiLineComments: [][]string{{"/*", "*/"}},
		Extensions:        []string{".sql"},
		FileNames:         []string{},
	},
	"Swift": {
		LineComments:      []string{"//"},
		MultiLineComments: [][]string{{"/*", "*/"}},
		Extensions:        []string{".swift"},
		FileNames:         []string{},
	},
	"TypeScript": {
		LineComments:      []string{"//"},
		MultiLineComments: [][]string{{"/*", "*/"}},
		Extensions:        []string{".ts", ".tsx"},
		FileNames:         []string{},
	},
	"T-SQL": {
		LineComments:      []string{"--"},
		MultiLineComments: [][]string{},
		Extensions:        []string{".tsql"},
		FileNames:         []string{},
	},
	"Vue": {
		LineComments:      []string{"<!--"},
		MultiLineComments: [][]string{{"<!--", "-->"}},
		Extensions:        []string{".vue"},
		FileNames:         []string{},
	},
	"Visual Basic .NET": {
		LineComments:      []string{"'"},
		MultiLineComments: [][]string{},
		Extensions:        []string{".vb"},
		FileNames:         []string{},
	},
	"XML": {
		LineComments:      []string{"<!--"},
		MultiLineComments: [][]string{{"<!--", "-->"}},
		Extensions:        []string{".xml", ".XML", ".xsd", ".xsl"},
		FileNames:         []string{},
	},
	"XHTML": {
		LineComments:      []string{"<!--"},
		MultiLineComments: [][]string{{"<!--", "-->"}},
		Extensions:        []string{".xhtml"},
		FileNames:         []string{},
	},
	"YAML": {
		LineComments:      []string{"#"},
		MultiLineComments: [][]string{},
		Extensions:        []string{".yaml", ".yml"},
		FileNames:         []string{},
	},
	"Terraform": {
		LineComments:      []string{},
		MultiLineComments: [][]string{},
		Extensions:        []string{".tf"},
		FileNames:         []string{},
	},
	"JCL": {
		LineComments:      []string{"//"},
		MultiLineComments: [][]string{{"/*", "*/"}},
		Extensions:        []string{".jcl", ".JCL"},
		FileNames:         []string{},
	},
	"Docker": {
		LineComments:      []string{"#"},
		MultiLineComments: [][]string{},
		Extensions:        []string{".dockerfile"},
		FileNames:         []string{"Dockerfile"},
	},
}

// Function to look up file information based on its extension
/*
@ext should match exactly as above, ".java" etc.
*/
func LookupByExtension(ext string) (string, LanguageInfo, bool) {
	for lang, info := range Languages {
		for _, languageExt := range info.Extensions {
			if languageExt == ext {
				return lang, info, true
			}
		}
	}
	return "", LanguageInfo{}, false
}

func LookupByFileName(fileName string) (string, LanguageInfo, bool) {
	for lang, info := range Languages {
		for _, languageFileName := range info.FileNames {
			if languageFileName == fileName {
				return lang, info, true
			}
		}
	}
	return "", LanguageInfo{}, false
}

func PrintLanguages() {
	logger.Info("Supported Languages:")
	// Create a buffer to hold the JSON data
	var buf bytes.Buffer

	// Create a new JSON encoder and set SetEscapeHTML to false
	encoder := json.NewEncoder(&buf)
	encoder.SetEscapeHTML(false)
	encoder.SetIndent("", "  ")
	// Encode the map to JSON
	if err := encoder.Encode(Languages); err != nil {
		logger.Error("Error encoding JSON: ", err)
		logger.LogStackTraceAndExit(err)
	}

	// Print the JSON string
	fmt.Println(buf.String())
}

// LoadLanguages reads the JSON file and overrides the default Languages map
func LoadLanguages(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		logger.LogStackTraceAndExit(err)
	}
	defer file.Close()

	byteValue, err := io.ReadAll(file)
	if err != nil {
		logger.LogStackTraceAndExit(err)
	}

	err = json.Unmarshal(byteValue, &Languages)
	if err != nil {
		logger.LogStackTraceAndExit(err)
	}
}
