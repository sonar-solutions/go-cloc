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
}

var Languages = map[string]LanguageInfo{
	"ActionScript": {
		LineComments:      []string{"//"},
		MultiLineComments: [][]string{{"/*", "*/"}},
		Extensions:        []string{".as"},
	},
	"Abap": {
		LineComments:      []string{"\""},
		MultiLineComments: [][]string{{"/*", "*/"}},
		Extensions:        []string{".abap", ".ab4", ".flow"},
	},
	"Apex": {
		LineComments:      []string{"//"},
		MultiLineComments: [][]string{{"/*", "*/"}},
		Extensions:        []string{".cls", ".trigger"},
	},
	"C": {
		LineComments:      []string{"//"},
		MultiLineComments: [][]string{{"/*", "*/"}},
		Extensions:        []string{".c"},
	},
	"C Header": {
		LineComments:      []string{"//"},
		MultiLineComments: [][]string{{"/*", "*/"}},
		Extensions:        []string{".h"},
	},
	"C++": {
		LineComments:      []string{"//"},
		MultiLineComments: [][]string{{"/*", "*/"}},
		Extensions:        []string{".cpp", ".cc", ".cxx", ".c++"},
	},
	"C++ Header": {
		LineComments:      []string{"//"},
		MultiLineComments: [][]string{{"/*", "*/"}},
		Extensions:        []string{".hh", ".hpp", ".hxx", ".h++", ".ipp"},
	},
	"COBOL": {
		LineComments:      []string{"*", "/"},
		MultiLineComments: [][]string{},
		Extensions:        []string{".cbl", ".ccp", ".cob", ".cobol", ".cpy"},
	},
	"C#": {
		LineComments:      []string{"//"},
		MultiLineComments: [][]string{{"/*", "*/"}},
		Extensions:        []string{".cs"},
	},
	"CSS": {
		LineComments:      []string{"//"},
		MultiLineComments: [][]string{{"/*", "*/"}},
		Extensions:        []string{".css"},
	},
	"Golang": {
		LineComments:      []string{"//"},
		MultiLineComments: [][]string{{"/*", "*/"}},
		Extensions:        []string{".go"},
	},
	"HTML": {
		LineComments:      []string{},
		MultiLineComments: [][]string{{"<!--", "-->"}},
		Extensions:        []string{".html", ".htm", ".cshtml", ".vbhtml", ".aspx", ".ascx", ".rhtml", ".erb", ".shtml", ".shtm", ".cmp"},
	},
	"Java": {
		LineComments:      []string{"//"},
		MultiLineComments: [][]string{{"/*", "*/"}},
		Extensions:        []string{".java", ".jav"},
	},
	"JavaScript": {
		LineComments:      []string{"//"},
		MultiLineComments: [][]string{{"/*", "*/"}},
		Extensions:        []string{".js", ".jsx", ".jsp", ".jspx", ".jspf", ".mjs"},
	},
	"Kotlin": {
		LineComments:      []string{"//"},
		MultiLineComments: [][]string{{"/*", "*/"}},
		Extensions:        []string{".kt", ".kts"},
	},
	"Flex": {
		LineComments:      []string{"//"},
		MultiLineComments: [][]string{{"/*", "*/"}},
		Extensions:        []string{".as"},
	},
	"PHP": {
		LineComments:      []string{"//", "#"},
		MultiLineComments: [][]string{{"/*", "*/"}},
		Extensions:        []string{".php", ".php3", ".php4", ".php5", ".phtml", ".inc"},
	},
	"Objective-C": {
		LineComments:      []string{"//"},
		MultiLineComments: [][]string{{"/*", "*/"}},
		Extensions:        []string{".m"},
	},
	"Oracle PL/SQL": {
		LineComments:      []string{"--"},
		MultiLineComments: [][]string{{"/*", "*/"}},
		Extensions:        []string{".pkb"},
	},
	"PL/I": {
		LineComments:      []string{"--"},
		MultiLineComments: [][]string{{"/*", "*/"}},
		Extensions:        []string{".pl1"},
	},
	"Python": {
		LineComments:      []string{"#"},
		MultiLineComments: [][]string{{"\"\"\"", "\"\"\""}},
		Extensions:        []string{".py", ".python", ".ipynb"},
	},

	"RPG": {
		LineComments:      []string{"#"},
		MultiLineComments: [][]string{},
		Extensions:        []string{".rpg"},
	},
	"Ruby": {
		LineComments:      []string{"#"},
		MultiLineComments: [][]string{{"=begin", "=end"}},
		Extensions:        []string{".rb"},
	},
	"Scala": {
		LineComments:      []string{"//"},
		MultiLineComments: [][]string{{"/*", "*/"}},
		Extensions:        []string{".scala"},
	},
	"Scss": {
		LineComments:      []string{"//"},
		MultiLineComments: [][]string{{"/*", "*/"}},
		Extensions:        []string{".scss"},
	},
	"SQL": {
		LineComments:      []string{"--"},
		MultiLineComments: [][]string{{"/*", "*/"}},
		Extensions:        []string{".sql"},
	},
	"Swift": {
		LineComments:      []string{"//"},
		MultiLineComments: [][]string{{"/*", "*/"}},
		Extensions:        []string{".swift"},
	},
	"TypeScript": {
		LineComments:      []string{"//"},
		MultiLineComments: [][]string{{"/*", "*/"}},
		Extensions:        []string{".ts", ".tsx"},
	},
	"T-SQL": {
		LineComments:      []string{"--"},
		MultiLineComments: [][]string{},
		Extensions:        []string{".tsql"},
	},
	"Vue": {
		LineComments:      []string{"<!--"},
		MultiLineComments: [][]string{{"<!--", "-->"}},
		Extensions:        []string{".vue"},
	},
	"Visual Basic .NET": {
		LineComments:      []string{"'"},
		MultiLineComments: [][]string{},
		Extensions:        []string{".vb"},
	},
	"XML": {
		LineComments:      []string{"<!--"},
		MultiLineComments: [][]string{{"<!--", "-->"}},
		Extensions:        []string{".xml", ".XML", ".xsd", ".xsl"},
	},
	"XHTML": {
		LineComments:      []string{"<!--"},
		MultiLineComments: [][]string{{"<!--", "-->"}},
		Extensions:        []string{".xhtml"},
	},
	"YAML": {
		LineComments:      []string{"#"},
		MultiLineComments: [][]string{},
		Extensions:        []string{".yaml", ".yml"},
	},
	"Terraform": {
		LineComments:      []string{},
		MultiLineComments: [][]string{},
		Extensions:        []string{".tf"},
	},
	"JCL": {
		LineComments:      []string{"//"},
		MultiLineComments: [][]string{{"/*", "*/"}},
		Extensions:        []string{".jcl", ".JCL"},
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
