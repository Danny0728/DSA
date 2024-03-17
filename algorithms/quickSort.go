package algorithms

import (
	"math/rand"
)

// quickSort uses divide&conquer and recursion for implementation
/*
Average time complexity - O(n logn)
Worst Case - O(n^2)
BestCase - O(n logn)
*/
func QuickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	pivotIndex := rand.Intn(len(arr))
	pivot := arr[pivotIndex]
	arr[0], arr[pivotIndex] = arr[pivotIndex], arr[0]
	var (
		leftArr, rightArr []int
	)

	for _, num := range arr[1:] {
		if num <= pivot {
			leftArr = append(leftArr, num)
		} else {
			rightArr = append(rightArr, num)
		}
	}

	sortedLeft := QuickSort(leftArr)
	sortedRight := QuickSort(rightArr)

	return append(append(sortedLeft, pivot), sortedRight...)
}

// I wanted to do inplace quickSort as selectionSort is also inplace inherently
func InPlaceQuickSort(arr []int, high, low int) {
	if low < high {
		pivotIndex := Partition(arr, high, low)
		InPlaceQuickSort(arr, low, pivotIndex-1)
		InPlaceQuickSort(arr, pivotIndex+1, high)
	}
}

func Partition(arr []int, high, low int) int {
	pivot := arr[high]
	i := low - 1 //i is just a conceptual counter, to let us understand that if an element is less that the pivot
	//they have a place to be.

	for j := low; j < high; j++ {
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[high] = arr[high], arr[i+1] //so that element to its left are lesser than pivot while to its right greater.
	return i + 1
}
