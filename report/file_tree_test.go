package report

import (
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
