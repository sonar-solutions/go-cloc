package main

import (
	"fmt"
	"go-cloc/logger"
	"go-cloc/report"
	"go-cloc/scanner"
	"go-cloc/utilities"
	"path/filepath"
)

func main() {
	// parse CLI arguments and store them in a struct
	args := utilities.ParseArgsFromCLI()

	// scan LOC for the directory
	logger.Info("Scanning ", args.LocalScanFilePath, "...")
	filePaths := scanner.WalkDirectory(args.LocalScanFilePath, args.IgnorePatterns)
	fileScanResultsArr := []scanner.FileScanResults{}
	for _, filePath := range filePaths {
		fileScanResultsArr = append(fileScanResultsArr, scanner.ScanFile(filePath))
	}

	logger.Debug("Calculating total LOC ...")

	// sort and calculate total LOC
	fileScanResultsArr = report.SortFileScanResults(fileScanResultsArr)
	repoTotalResult := report.CalculateTotalLineOfCode(fileScanResultsArr)

	// convert results into records for CSV or command line output
	records := report.ConvertFileResultsIntoRecords(fileScanResultsArr, repoTotalResult)

	// Dump results by file in a csv
	if args.CsvFilePath != "" {
		logger.Debug("Dumping results by file to ", args.CsvFilePath)
		report.WriteCsv(args.CsvFilePath, records)
		logger.Info("Done! Results can be found ", args.CsvFilePath)
	} else {
		// print results to the command line
		logger.Info("Results by file for ", args.LocalScanFilePath, ":")
		report.PrintCsv(records)
	}

	if args.HtmlReportsDirectoryPath != "" {
		logger.Info("Dumping HTML report to ", args.HtmlReportsDirectoryPath)
		fileNames, fileContents := report.GenerateHTMLReports(fileScanResultsArr)

		for index, _ := range fileNames {
			fileName := fileNames[index]
			fileContent := fileContents[index]
			report.WriteStringToFile(filepath.Join(args.HtmlReportsDirectoryPath, fileName), fileContent)
		}
		report.DumpSVGs(args.HtmlReportsDirectoryPath)
		logger.Info("Done! HTML report for ", args.LocalScanFilePath, " can be found in ", args.HtmlReportsDirectoryPath)
	}

	logger.Info("Total LOC for ", args.LocalScanFilePath, " is ", repoTotalResult.CodeLineCount)

	// Print the total LOC to standard output to make it easy for external tools to parse
	fmt.Println(repoTotalResult.CodeLineCount)
}
