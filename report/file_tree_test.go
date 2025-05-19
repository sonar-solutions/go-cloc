package report

import (
	"go-cloc/scanner"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_file_tree_sum(t *testing.T) {
	mapA := map[string]int{"a": 10}
	mapB := map[string]int{"b": 20, "a": 5}
	result := combineMapsAndSum(mapA, mapB)

	// Assert
	assert.Equal(t, 15, result["a"])
	assert.Equal(t, 20, result["b"])
}

func Test_file_tree_createTreeFromScanResults(t *testing.T) {
	fileScanResults := []scanner.FileScanResults{
		{FilePath: "/home/file1.go", LanguageName: "go", CodeLineCount: 10},
		{FilePath: "/home/file2.java", LanguageName: "java", CodeLineCount: 20},
		{FilePath: "/test/file3.py", LanguageName: "python", CodeLineCount: 30},
	}
	root := createTreeFromScanResults(fileScanResults)

	// check root
	assert.NotNil(t, root)
	assert.Equal(t, 2, len(root.children))

	// check home
	home := root.children[0]

	assert.NotNil(t, home)

	assert.Equal(t, "home", home.name)
	assert.Equal(t, 2, len(home.children))

	file1 := home.children[0]
	assert.NotNil(t, file1)
	assert.Equal(t, 10, file1.CodeLineCount)
	assert.Equal(t, "file1.go", file1.name)

	file2 := home.children[1]
	assert.NotNil(t, file2)
	assert.Equal(t, 20, file2.CodeLineCount)
	assert.Equal(t, "file2.java", file2.name)

	// check test
	test := root.children[1]
	assert.NotNil(t, test)
	assert.Equal(t, "test", test.name)
	assert.Equal(t, 1, len(test.children))

	file3 := test.children[0]
	assert.NotNil(t, file3)
	assert.Equal(t, 30, file3.CodeLineCount)
	assert.Equal(t, "file3.py", file3.name)

}
