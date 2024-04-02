package testing

import (
	"reflect"
	"testing"

	"github.com/Danny0728/GrokkingAlgorithms/algorithms"
)

func TestQuickSort(t *testing.T) {
	input := []int{13, 2, 3, 10, 5, 6, 7, 18, 22, 0, 8, 99, 1}
	expected := []int{0, 1, 2, 3, 5, 6, 7, 8, 10, 13, 18, 22, 99}

	result := algorithms.QuickSort(input)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("QuickSort returned %v, \n Expected %v", result, expected)
	}
}
