package utilities

import (
	"flag"
	"go-cloc/logger"
	"go-cloc/scanner"
	"os"
	"path/filepath"
	"strings"
)

// Modes
const (
	LOCAL       string = "Local"
	GITHUB      string = "GitHub"
	AZUREDEVOPS string = "AzureDevOps"
	GITLAB      string = "GitLab"
	BITBUCKET   string = "Bitbucket"
)

type CLIArgs struct {
	LogLevel                        string
	LocalScanFilePath               string
	IgnorePatterns                  []string
	CsvFilePath                     string
	HtmlReportsDirectoryPath        string
	OverrideLanguagesConfigFilePath string
}

func CleanLocalFilePath(targetPath string) string {
	logger.Debug("CleanLocalFilePath targetPath before: '", targetPath, "'")
	targetPath = filepath.Clean(targetPath)
	// On windows this may be needed if spaces are in the file path
	targetPath = strings.TrimSuffix(targetPath, "\"")
	logger.Debug("CleanLocalFilePath targetPath after: '", targetPath, "'")
	return targetPath
}

func ParseArgsFromCLI() CLIArgs {
	// print out arguments
	printLanguagesArg := flag.Bool("print-languages", false, "Prints out the supported languages, file suffixes, and comment configurations. Does not run the tool.")

	// optional arguments
	logLevelArg := flag.String("log-level", "INFO", "Log level - DEBUG, INFO, WARN, ERROR")
	ignoreFilePathArg := flag.String("ignore-file-path", "", "Path to your ignore file. Defines directories and files to exclude when scanning. Please see the README.md for how to format your ignore configuration")
	csvFilePathArg := flag.String("csv", "", "Path to dump results to a csv file, otherwise results are printed to standard out")
	htmlReportsDirectoryPathArg := flag.String("html", "", "Path to dump HTML reports into a specified directory, otherwise HTML reports are not generated. Note this directory must already exist.")
	overrideLanguageConfigFilePathArg := flag.String("override-languages", "", "Path to languages configuration to override the default configuration.")

	// parse the CLI arguments
	flag.Parse()

	// dereference all CLI args to make it easier to use
	printLanguages := *printLanguagesArg

	// print out languages
	if printLanguages {
		scanner.PrintLanguages()
		os.Exit(0)
	}

	// Collect the remaining arguments
	cliArgs := flag.Args()

	// Ensure at least one argument
	if len(cliArgs) < 1 {
		logger.Error("Requires a path to the file or directory to scan as the first command line argument, ex: 'go-cloc file1.js'")
		os.Exit(-1)
	}

	// Parse any remaining flags after the first non-flag argument
	flag.CommandLine.Parse(cliArgs[1:])

	// dereference all CLI args to make it easier to use
	logLevel := *logLevelArg
	ignoreFilePath := *ignoreFilePathArg
	csvFilePath := *csvFilePathArg
	htmlReportsDirectoryPath := *htmlReportsDirectoryPathArg
	overrideLanguageConfigFilePath := *overrideLanguageConfigFilePathArg

	// Check if the directory exists
	if htmlReportsDirectoryPath != "" {
		// only create the folder if the folder does not exist
		_, err := os.Stat(htmlReportsDirectoryPath)
		if os.IsNotExist(err) {
			logger.Error("Folder does not exist. Please create it first. Path: ", htmlReportsDirectoryPath)
			os.Exit(-1)
		}
	}

	// set log level
	logger.SetLogLevel(logger.ConvertStringToLogLevel(logLevel))
	logger.SetOutput(os.Stdout)

	logger.Info("Setting Log Level to " + logLevel)
	logger.Info("Parsing CLI arguments")

	// print out arguments
	logger.Debug("csv-file-path: ", csvFilePath)
	logger.Debug("html-reports-directory-path: ", htmlReportsDirectoryPath)
	logger.Debug("ignore-file-path: ", ignoreFilePath)
	logger.Debug("override-language-config-file-path: ", overrideLanguageConfigFilePath)

	// Set file path to scan
	localScanFilePath := CleanLocalFilePath(cliArgs[0])

	// validate optional arguments

	// parse ignore patterns
	ignorePatterns := []string{}
	if ignoreFilePath != "" {
		logger.Debug("Parsing ignore-file ", ignoreFilePath)
		ignorePatterns = scanner.ReadIgnoreFile(ignoreFilePath)
		logger.Debug("Successfully read in the ignore-file ", ignoreFilePath)
		logger.Debug("Ignore Patterns: ", ignorePatterns)
	}

	// override languages config
	if overrideLanguageConfigFilePath != "" {
		logger.Debug("Overriding default languages with ", overrideLanguageConfigFilePath)
		scanner.LoadLanguages(overrideLanguageConfigFilePath)
	}

	args := CLIArgs{
		LogLevel:                        logLevel,
		LocalScanFilePath:               localScanFilePath,
		IgnorePatterns:                  ignorePatterns,
		CsvFilePath:                     csvFilePath,
		HtmlReportsDirectoryPath:        htmlReportsDirectoryPath,
		OverrideLanguagesConfigFilePath: overrideLanguageConfigFilePath,
	}

	return args
}
