package report

import (
	"go-cloc/logger"
	"go-cloc/scanner"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

type FileTreeComponent struct {
	parent                  *FileTreeComponent
	children                []*FileTreeComponent
	name                    string
	CodeLineCount           int
	LanguageToCodeLineCount map[string]int // map of language to code line count, empty by default
}

// Pair is a simple string-int pair
type Pair struct {
	Key   string
	Value int
}

func addChildIfNotExists(parent *FileTreeComponent, child *FileTreeComponent) *FileTreeComponent {
	// Check if child already exists under parent
	for _, existingChild := range parent.children {
		if existingChild.name == child.name {
			logger.Debug("Child " + child.name + " already exists under parent " + parent.name)
			return existingChild
		}
	}
	// Add child if it doesn't exist under parent
	parent.children = append(parent.children, child)
	child.parent = parent
	return child
}

// returns a full path in the tree from the root of the tree to the current component
func getFullPathInTree(component *FileTreeComponent, osFileSeparator string) string {
	if component.parent == nil {
		return component.name
	}
	return getFullPathInTree(component.parent, osFileSeparator) + osFileSeparator + component.name
}

// creates a unique file name based on the components location in the tree
func createUniqueFileNameFromComponentInTree(component *FileTreeComponent) string {
	if component == nil {
		return "index.html"
	}
	return getFullPathInTree(component, "-") + ".html"
}

// traverses the tree generating HTML reports for directories only
func generateHTMLReportsForTree(component *FileTreeComponent) ([]string, []string) {
	// return empty arrays if the component has no children, since this is a file
	if len(component.children) == 0 {
		return []string{}, []string{}
	}

	// given this is a directory, create a HTML report
	resultingFileNames := []string{}
	resultingFileContents := []string{}
	fileName := createUniqueFileNameFromComponentInTree(component)
	htmlContent := createHTMLPage(component)
	resultingFileNames = append(resultingFileNames, fileName)
	resultingFileContents = append(resultingFileContents, htmlContent)

	// recursively generate HTML reports for each child
	for _, child := range component.children {
		fileNames, fileContents := generateHTMLReportsForTree(child)
		resultingFileNames = append(resultingFileNames, fileNames...)
		resultingFileContents = append(resultingFileContents, fileContents...)
	}

	return resultingFileNames, resultingFileContents
}

// creates the HTML page for a given component, designed for directories
func createHTMLPage(component *FileTreeComponent) string {
	if len(component.children) == 0 {
		return ""
	}
	// Create HTML content here based on component
	htmlContent := "<!DOCTYPE html><html lang='en'><head><meta charset='UTF-8'><style>body{font-family:Arial,sans-serif}td{padding:8px;border-bottom:1px solid #ddd}th{background-color:#f2f2f2;padding:8px}a{color:#00f;text-decoration:none}a:hover{text-decoration:underline}.code-line-count{padding:8px;border-bottom:1px solid #ddd}.file,.folder{padding:10px;display:inline-block;width:20px;vertical-align:middle}</style><meta name='viewport' content='width=device-width,initial-scale=1'><title>File Tree Report</title></head><body><h1>File Tree Report</h1>"
	htmlContent += "<p><b>Current Path:</b><a href='" + createUniqueFileNameFromComponentInTree(component.parent) + "'> '" + getFullPathInTree(component, string(filepath.Separator)) + "' </a><span style='color:gray;'>&lAarr; Click to return</span></p>"
	htmlContent += "<table><thead><tr><th>File Name</th><th>Code Line Count</th></tr><tr><thead></thead></tr></thead><tbody>"
	for _, child := range component.children {
		htmlContent += "<tr><td>"
		// file
		if len(child.children) == 0 {
			htmlContent += "<img src='file-text.svg' alt='' class='file'> <span>" + child.name + "</span>"
		} else {
			htmlContent += "<img src='folder.svg' alt='' class='folder'> <a href='./" + createUniqueFileNameFromComponentInTree(child) + "'>" + child.name + "</a>"
		}
		htmlContent += "</td><td class='code-line-count'>" + strconv.Itoa(child.CodeLineCount) + "</td></tr>"

	}
	htmlContent += "</tbody>"
	htmlContent += "<tfoot><tr><th>Total Code Lines:</th><th class='code-line-count'>" + strconv.Itoa(component.CodeLineCount) + "</th></tfoot>"
	htmlContent += "</table>"
	// add language statistics
	htmlContent += "<h2>Language Statistics</h2>"
	htmlContent += "<table><thead><tr><th>Language</th><th>Code Line Count</th></tr><tr><thead></thead></tr></thead><tbody>"
	// iterate a map
	for _, pair := range sortKeysByValueInMap(component.LanguageToCodeLineCount) {
		htmlContent += "<tr><td>" + pair.Key + "</td><td>" + strconv.Itoa(pair.Value) + "</td></tr>"
	}
	htmlContent += "</tbody>"
	htmlContent += "<tfoot><tr><th>Total Code Lines:</th><th class='code-line-count'>" + strconv.Itoa(component.CodeLineCount) + "</th></tfoot>"
	htmlContent += "</table>"
	htmlContent += "</body></html>"
	return htmlContent
}

// combine two maps together into a single map and sum up their matching keys
func combineMapsAndSum(a map[string]int, b map[string]int) map[string]int {
	result := map[string]int{}
	// just copy all keys from the first map to the result map
	for key, value := range a {
		result[key] = value
	}
	// then add all keys from the second map to the result map
	for key, value := range b {
		if _, ok := result[key]; !ok {
			result[key] = value
		} else {
			result[key] += value
		}
	}
	return result

}

// sets the code line count for every component in the tree
func sumUpTotalLineOfCodeInTree(component *FileTreeComponent) (int, map[string]int) {
	if component == nil {
		return 0, map[string]int{}
	}
	// base case: if the component has no children, return its code line count
	if len(component.children) == 0 {
		return component.CodeLineCount, component.LanguageToCodeLineCount
	}

	sum := 0
	sumLanguageToCodeLineCount := map[string]int{}
	for _, child := range component.children {
		sumChildren, languageToCodeLineCount := sumUpTotalLineOfCodeInTree(child)
		sum += sumChildren
		sumLanguageToCodeLineCount = combineMapsAndSum(sumLanguageToCodeLineCount, languageToCodeLineCount)
		logger.Debug("languageToCodeLineCount: ", languageToCodeLineCount)
	}
	logger.Debug("Len of children: ", len(component.children))
	logger.Debug("sumLanguageToCodeLineCount: ", sumLanguageToCodeLineCount)
	component.CodeLineCount = sum
	component.LanguageToCodeLineCount = sumLanguageToCodeLineCount
	return sum, sumLanguageToCodeLineCount
}

// helper function to sort a list of FileTreeComponents by their CodeLineCount
func sortComponentsByCodeLineCount(components []*FileTreeComponent) []*FileTreeComponent {
	sort.Slice(components, func(i, j int) bool {
		return components[i].CodeLineCount > components[j].CodeLineCount
	})
	return components
}

// helper function to sum up total line of code in a FileTreeComponent and its children
func sortKeysByValueInMap(m map[string]int) []Pair {
	pairs := make([]Pair, 0, len(m))
	for k, v := range m {
		pairs = append(pairs, Pair{Key: k, Value: v})
	}

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].Value > pairs[j].Value
	})

	return pairs
}

// traverses the tree and every components' children by their CodeLineCount
func sortTreeByCodeLineCount(component *FileTreeComponent) {
	component.children = sortComponentsByCodeLineCount(component.children)
	for _, child := range component.children {
		sortTreeByCodeLineCount(child)
	}
}

// Creates HTML reports to visualize the LoC in the same file structure as was scanned. Helpful for identifying large directories.
func GenerateHTMLReports(fileScanResults []scanner.FileScanResults) ([]string, []string) {

	// create root node of the tree
	root := &FileTreeComponent{
		// TODO change me
		name:                    "index",
		parent:                  nil,
		children:                []*FileTreeComponent{},
		LanguageToCodeLineCount: map[string]int{},
	}

	// create tree structure based on the fileScanResults
	for _, result := range fileScanResults {
		previousComponent := root
		filePathComponents := ParseFileStructure(result.FilePath, string(filepath.Separator))
		for _, component := range filePathComponents {
			newChild := &FileTreeComponent{
				name:                    component,
				children:                []*FileTreeComponent{},
				LanguageToCodeLineCount: map[string]int{},
			}
			previousComponent = addChildIfNotExists(previousComponent, newChild)
		}
		previousComponent.CodeLineCount = result.CodeLineCount
		previousComponent.LanguageToCodeLineCount[result.LanguageName] = result.CodeLineCount
	}

	// calculate total LOC in the tree
	sumUpTotalLineOfCodeInTree(root)

	// sort the tree by CodeLineCount
	sortTreeByCodeLineCount(root)

	// generate HTML reports for each file in the tree
	return generateHTMLReportsForTree(root)
}

// simple function to write SVGs to files for the HTML reports to use
func DumpSVGs(outputFolderPath string) {
	WriteStringToFile(filepath.Join(outputFolderPath, "file-text.svg"),
		"<svg width='800px' height='800px' viewBox='0 0 24 24' fill='none' xmlns='http://www.w3.org/2000/svg'><path d='M15.3929 4.05365L14.8912 4.61112L15.3929 4.05365ZM19.3517 7.61654L18.85 8.17402L19.3517 7.61654ZM21.654 10.1541L20.9689 10.4592V10.4592L21.654 10.1541ZM3.17157 20.8284L3.7019 20.2981H3.7019L3.17157 20.8284ZM20.8284 20.8284L20.2981 20.2981L20.2981 20.2981L20.8284 20.8284ZM14 21.25H10V22.75H14V21.25ZM2.75 14V10H1.25V14H2.75ZM21.25 13.5629V14H22.75V13.5629H21.25ZM14.8912 4.61112L18.85 8.17402L19.8534 7.05907L15.8947 3.49618L14.8912 4.61112ZM22.75 13.5629C22.75 11.8745 22.7651 10.8055 22.3391 9.84897L20.9689 10.4592C21.2349 11.0565 21.25 11.742 21.25 13.5629H22.75ZM18.85 8.17402C20.2034 9.3921 20.7029 9.86199 20.9689 10.4592L22.3391 9.84897C21.9131 8.89241 21.1084 8.18853 19.8534 7.05907L18.85 8.17402ZM10.0298 2.75C11.6116 2.75 12.2085 2.76158 12.7405 2.96573L13.2779 1.5653C12.4261 1.23842 11.498 1.25 10.0298 1.25V2.75ZM15.8947 3.49618C14.8087 2.51878 14.1297 1.89214 13.2779 1.5653L12.7405 2.96573C13.2727 3.16993 13.7215 3.55836 14.8912 4.61112L15.8947 3.49618ZM10 21.25C8.09318 21.25 6.73851 21.2484 5.71085 21.1102C4.70476 20.975 4.12511 20.7213 3.7019 20.2981L2.64124 21.3588C3.38961 22.1071 4.33855 22.4392 5.51098 22.5969C6.66182 22.7516 8.13558 22.75 10 22.75V21.25ZM1.25 14C1.25 15.8644 1.24841 17.3382 1.40313 18.489C1.56076 19.6614 1.89288 20.6104 2.64124 21.3588L3.7019 20.2981C3.27869 19.8749 3.02502 19.2952 2.88976 18.2892C2.75159 17.2615 2.75 15.9068 2.75 14H1.25ZM14 22.75C15.8644 22.75 17.3382 22.7516 18.489 22.5969C19.6614 22.4392 20.6104 22.1071 21.3588 21.3588L20.2981 20.2981C19.8749 20.7213 19.2952 20.975 18.2892 21.1102C17.2615 21.2484 15.9068 21.25 14 21.25V22.75ZM21.25 14C21.25 15.9068 21.2484 17.2615 21.1102 18.2892C20.975 19.2952 20.7213 19.8749 20.2981 20.2981L21.3588 21.3588C22.1071 20.6104 22.4392 19.6614 22.5969 18.489C22.7516 17.3382 22.75 15.8644 22.75 14H21.25ZM2.75 10C2.75 8.09318 2.75159 6.73851 2.88976 5.71085C3.02502 4.70476 3.27869 4.12511 3.7019 3.7019L2.64124 2.64124C1.89288 3.38961 1.56076 4.33855 1.40313 5.51098C1.24841 6.66182 1.25 8.13558 1.25 10H2.75ZM10.0298 1.25C8.15538 1.25 6.67442 1.24842 5.51887 1.40307C4.34232 1.56054 3.39019 1.8923 2.64124 2.64124L3.7019 3.7019C4.12453 3.27928 4.70596 3.02525 5.71785 2.88982C6.75075 2.75158 8.11311 2.75 10.0298 2.75V1.25Z' fill='#1C274C'/><path opacity='0.5' d='M6 14.5H14' stroke='#1C274C' stroke-width='1.5' stroke-linecap='round'/><path opacity='0.5' d='M6 18H11.5' stroke='#1C274C' stroke-width='1.5' stroke-linecap='round'/><path opacity='0.5' d='M13 2.5V5C13 7.35702 13 8.53553 13.7322 9.26777C14.4645 10 15.643 10 18 10H22' stroke='#1C274C' stroke-width='1.5'/></svg>")
	WriteStringToFile(filepath.Join(outputFolderPath, "folder.svg"),
		"<?xml version='1.0' encoding='iso-8859-1'?> <!-- Uploaded to: SVG Repo, www.svgrepo.com, Generator: SVG Repo Mixer Tools --> <svg version='1.1' id='Layer_1' xmlns='http://www.w3.org/2000/svg' xmlns:xlink='http://www.w3.org/1999/xlink' viewBox='0 0 512 512' xml:space='preserve'><path id='SVGCleanerId_0' style='fill:#ffc36e' d='M183.295,123.586H55.05c-6.687,0-12.801-3.778-15.791-9.76l-12.776-25.55	l12.776-25.55c2.99-5.982,9.103-9.76,15.791-9.76h128.246c6.687,0,12.801,3.778,15.791,9.76l12.775,25.55l-12.776,25.55	C196.096,119.808,189.983,123.586,183.295,123.586z'/><g><path id='SVGCleanerId_0_1_' style='fill:#ffc36e' d='M183.295,123.586H55.05c-6.687,0-12.801-3.778-15.791-9.76l-12.776-25.55l12.776-25.55c2.99-5.982,9.103-9.76,15.791-9.76h128.246c6.687,0,12.801,3.778,15.791,9.76l12.775,25.55l-12.776,25.55C196.096,119.808,189.983,123.586,183.295,123.586z'/></g><path style='fill:#eff2fa' d='M485.517,70.621H26.483c-4.875,0-8.828,3.953-8.828,8.828v44.138h476.69V79.448	C494.345,74.573,490.392,70.621,485.517,70.621z'/><rect x='17.655' y='105.931' style='fill:#e1e6f2' width='476.69' height='17.655'/><path style='fill:#ffd782' d='M494.345,88.276H217.318c-3.343,0-6.4,1.889-7.895,4.879l-10.336,20.671	c-2.99,5.982-9.105,9.76-15.791,9.76H55.05c-6.687,0-12.801-3.778-15.791-9.76L28.922,93.155c-1.495-2.99-4.552-4.879-7.895-4.879	h-3.372C7.904,88.276,0,96.18,0,105.931v335.448c0,9.751,7.904,17.655,17.655,17.655h476.69c9.751,0,17.655-7.904,17.655-17.655	V105.931C512,96.18,504.096,88.276,494.345,88.276z'/><path style='fill:#ffc36e' d='M485.517,441.379H26.483c-4.875,0-8.828-3.953-8.828-8.828l0,0c0-4.875,3.953-8.828,8.828-8.828	h459.034c4.875,0,8.828,3.953,8.828,8.828l0,0C494.345,437.427,490.392,441.379,485.517,441.379z'/><path style='fill:#eff2fa' d='M326.621,220.69h132.414c4.875,0,8.828-3.953,8.828-8.828v-70.621c0-4.875-3.953-8.828-8.828-8.828	H326.621c-4.875,0-8.828,3.953-8.828,8.828v70.621C317.793,216.737,321.746,220.69,326.621,220.69z'/><path style='fill:#c7cfe2' d='M441.379,167.724h-97.103c-4.875,0-8.828-3.953-8.828-8.828l0,0c0-4.875,3.953-8.828,8.828-8.828	h97.103c4.875,0,8.828,3.953,8.828,8.828l0,0C450.207,163.772,446.254,167.724,441.379,167.724z'/><path style='fill:#d7deed' d='M441.379,203.034h-97.103c-4.875,0-8.828-3.953-8.828-8.828l0,0c0-4.875,3.953-8.828,8.828-8.828	h97.103c4.875,0,8.828,3.953,8.828,8.828l0,0C450.207,199.082,446.254,203.034,441.379,203.034z'/></svg>")
}

// ParseFileStructure takes a file path and returns an array of strings representing the directory structure. It splits the file path by the OS-specific separator (e.g., `/` on Unix/Linux, `\` on Windows). The function ensures that each directory component is non-empty before adding it to the result slice. This helps in handling cases where there might be consecutive separators or empty directories.
func ParseFileStructure(filePath string, osFileSeparator string) []string {
	logger.Debug("File separator for OS is: ", osFileSeparator)

	pathComponents := make([]string, 0)
	for _, dir := range strings.Split(filePath, osFileSeparator) {
		if dir != "" {
			pathComponents = append(pathComponents, dir)
		}
	}

	return pathComponents
}

// helper function that writes the content of a file to a specified path. It uses the `os` package to create and write to files. The function logs any errors encountered during the process.
func WriteStringToFile(filePath string, content string) error {
	logger.Debug("Writing contents to file: ", filePath)
	file, err := os.Create(filePath)
	if err != nil {
		logger.Error("Error creating file: ", err)
		return err
	}
	defer file.Close()
	_, err = file.WriteString(content)
	return err
}
