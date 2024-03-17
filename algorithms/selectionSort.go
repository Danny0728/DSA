package algorithms

// SelectionSort sorts a slice of integers in ascending order using the Selection Sort algorithm
/*
Average time complexity - O(n^2)
Worst Case - O(n^2)
BestCase - O(n^2)
*/
func SelectionSort(arr []int) {

	for i := 0; i < len(arr)-1; i++ {
		min_index := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[min_index] {
				min_index = j
			}
		}
		//this improves efficiency as it reduces unnecessary swap if that element is already in correct position.
		if min_index != i {
			arr[i], arr[min_index] = arr[min_index], arr[i]
		}
	}
}
