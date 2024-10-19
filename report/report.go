package report

import (
	"encoding/csv"
	"go-cloc/logger"
	"go-cloc/scanner"
	"os"
	"sort"
	"strconv"
)

type RepoTotal struct {
	RepositoryId  string
	CodeLineCount int
}

// SortFileScanResults sorts the file scan results by CodeLineCount in descending order
func SortFileScanResults(fileScanResultsArr []scanner.FileScanResults) []scanner.FileScanResults {
	// Sort by CodeLineCount desc
	sort.Slice(fileScanResultsArr, func(a, b int) bool {
		return fileScanResultsArr[a].CodeLineCount > fileScanResultsArr[b].CodeLineCount
	})
	return fileScanResultsArr
}

// SortRepoTotalResults sorts the repo total results by CodeLineCount in descending order
func SortRepoTotalResults(repoTotalArr []RepoTotal) []RepoTotal {
	// Sort by CodeLineCount desc
	sort.Slice(repoTotalArr, func(a, b int) bool {
		return repoTotalArr[a].CodeLineCount > repoTotalArr[b].CodeLineCount
	})
	return repoTotalArr
}

// CalculateTotalLineOfCode calculates the total number of lines of code for all files scanned
func CalculateTotalLineOfCode(fileScanResultsArr []scanner.FileScanResults) scanner.FileScanResults {
	totalResults := scanner.FileScanResults{}

	totalResults.FilePath = "total"
	for _, results := range fileScanResultsArr {
		totalResults.BlankLineCount += results.BlankLineCount
		totalResults.CommentsLineCount += results.CommentsLineCount
		totalResults.CodeLineCount += results.CodeLineCount
		totalResults.TotalLines += results.TotalLines
	}
	return totalResults
}

// OutputCSV writes the results of the scan to a CSV file
// Returns the total number of lines of code for all files scanned
func ConvertFileResultsIntoRecords(fileScanResultsArr []scanner.FileScanResults, totalResults scanner.FileScanResults) [][]string {
	// Create CSV information
	records := [][]string{
		{"filePath", "blank", "comment", "code"},
	}

	for _, results := range fileScanResultsArr {
		row := []string{results.FilePath, strconv.Itoa(results.BlankLineCount), strconv.Itoa(results.CommentsLineCount), strconv.Itoa(results.CodeLineCount)}
		records = append(records, row)
	}
	// Append Total Row
	totalRow := []string{"total", strconv.Itoa(totalResults.BlankLineCount), strconv.Itoa(totalResults.CommentsLineCount), strconv.Itoa(totalResults.CodeLineCount)}
	records = append(records, totalRow)
	return records
}

// WriteCsv writes the records to a CSV file
func WriteCsv(outputFilePath string, records [][]string) error {
	// Write to csv
	f, err := os.Create(outputFilePath)
	if err != nil {
		logger.Error("Error creating csv file: ", err)
	}
	defer f.Close()

	w := csv.NewWriter(f)
	defer w.Flush()

	for _, row := range records {
		err = w.Write(row)
		if err != nil {
			logger.Error("Error writing to csv: ", err)
			return err
		}
	}
	return nil
}

// PrintCsv prints the records to the console, useful for debugging
func PrintCsv(records [][]string) {
	for _, row := range records {
		outputString := ""
		for col_index, col := range row {
			outputString += col
			if col_index < len(row)-1 {
				outputString += ","
			}
		}
		logger.Debug(outputString)
	}
}

func ConvertRepoTotalsIntoRecords(repoTotals []RepoTotal) [][]string {
	// Create CSV information
	records := [][]string{
		{"repository", "lineOfCodeCount"},
	}
	sum := 0
	for _, repoResult := range repoTotals {
		row := []string{repoResult.RepositoryId, strconv.Itoa(repoResult.CodeLineCount)}
		records = append(records, row)
		// keep running total
		sum += repoResult.CodeLineCount
	}
	// Create total row
	totalRow := []string{"total", strconv.Itoa(sum)}
	records = append(records, totalRow)
	return records
}
