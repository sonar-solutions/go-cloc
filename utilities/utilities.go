package utilities

import (
	"flag"
	"go-cloc/logger"
	"go-cloc/scanner"
	"os"
	"time"
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
	LogLevel             string
	LocalScanFilePath    string
	IgnorePatterns       []string
	DumpCSVs             bool
	ResultsDirectoryPath string
	ScanId               string
}

func ParseArgsFromCLI() CLIArgs {
	// print out arguments
	printLanguagesArg := flag.Bool("print-languages", false, "Prints out the supported languages, file suffixes, and comment configurations. Does not run the tool.")

	// mandatory arguments
	localScanFilePathArg := flag.String("path", "", "Path to your local file or directory that you wish to scan")

	// optional arguments
	logLevelArg := flag.String("log-level", "INFO", "Log level - DEBUG, INFO, WARN, ERROR")
	scanIdArg := flag.String("scan-id", "", "Identifier for the scan. For reference in a csv file later")
	ignoreFilePathArg := flag.String("ignore-file", "", "Path to your ignore file. Defines directories and files to exclude when scanning. Please see the README.md for how to format your ignore configuration")
	dumpCSVsArg := flag.Bool("dump-csv", false, "When true, dumps results to a csv file, otherwise gives results in logs")
	resultsDirectoryPathArg := flag.String("results-directory-path", "", "Path to a new directory for storing the results. Default the tool will create one based on the start time")

	// parse the CLI arguments
	flag.Parse()

	// dereference all CLI args to make it easier to use
	printLanguages := *printLanguagesArg
	logLevel := *logLevelArg
	localScanFilePath := *localScanFilePathArg
	ignoreFilePath := *ignoreFilePathArg
	dumpCSVs := *dumpCSVsArg
	resultsDirectoryPath := *resultsDirectoryPathArg
	scanId := *scanIdArg

	// set log level
	logger.SetLogLevel(logger.ConvertStringToLogLevel(logLevel))
	logger.SetOutput(os.Stdout)

	logger.Info("Setting Log Level to " + logLevel)
	logger.Info("Parsing CLI arguments")

	// print out arguments
	logger.Debug("dump-csvs: ", dumpCSVs)

	// print out languages
	if printLanguages {
		scanner.PrintLanguages()
		os.Exit(0)
	}

	logger.Debug("Validating mandatory arguments")

	// validate mandatory arguments
	if localScanFilePath == "" {
		logger.Error("Requires : --path")
		os.Exit(-1)
	}
	if dumpCSVs && scanId == "" {
		logger.Error("Requires : --scan-id for --dump-csvs")
		os.Exit(-1)
	}

	// validate optional arguments

	// parse ignore patterns
	ignorePatterns := []string{}
	if ignoreFilePath != "" {
		logger.Debug("Parsing ignore-file ", ignoreFilePath)
		ignorePatterns = scanner.ReadIgnoreFile(ignoreFilePath)
		logger.Debug("Successfully read in the ignore-file ", ignoreFilePath)
		logger.Debug("Ignore Patterns: ", ignorePatterns)
	}

	if !dumpCSVs && resultsDirectoryPath != "" {
		logger.Error("Cannot simultaneously set --results-directory-path and --dump-csvs=false")
		logger.LogStackTraceAndExit(nil)
	}

	// set results directory if dumpCSVs is true
	if resultsDirectoryPath == "" && dumpCSVs {
		resultsDirectoryPath = time.Now().Format("20060102_150405") // Format: YYYYMMDD_HHMMSS
		logger.Debug("Results Directory Path: ", resultsDirectoryPath)
	}

	args := CLIArgs{
		LogLevel:             logLevel,
		LocalScanFilePath:    localScanFilePath,
		IgnorePatterns:       ignorePatterns,
		DumpCSVs:             dumpCSVs,
		ResultsDirectoryPath: resultsDirectoryPath,
		ScanId:               scanId,
	}

	return args
}
