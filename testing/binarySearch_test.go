package testing

import (
	"testing"

	"github.com/Danny0728/GrokkingAlgorithms/algorithms"
)

func TestBinarySearch(t *testing.T) {
	inputArr := []int{1, 2, 5, 6, 7, 80, 90, 1000, 1100}
	numberToSearch := 80
	expected := 5

	result := algorithms.BinarySearch(inputArr, numberToSearch)
	if result != expected {
		t.Errorf("Binary Search returned %v, expected %v", result, expected)
	}
}
