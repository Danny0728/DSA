package algorithms

// 4.1 write a recursive function to add in a list.
func AddNums(arr []int, index int) int {
	if index == 0 {
		return arr[index]
	}
	return arr[index] + AddNums(arr, index-1)
}

// 4.2 write a recursive function to return no.of items in a list.
func CountItems(arr []int, index int) int {
	if index < 0 {
		return 0
	}
	return 1 + CountItems(arr, index-1)
}

// 4.3  write a recursive function to get the max int in a slice
func GetMaxNo(arr []int, index, maxNumber int) int {
	if index < 0 {
		return maxNumber
	}
	// fmt.Println(maxNumber)
	if arr[index] > maxNumber {
		maxNumber = arr[index]
	}
	return GetMaxNo(arr, index-1, maxNumber)
}

/*4.4 write a binary search using recursive algorithm
 */
func BinarySearchRecursive(arr []int, high, low, noToSearch int) int {
	if high >= low {
		mid := low + (high-low)/2
		if arr[mid] == noToSearch {
			return mid
		}
		if arr[mid] < noToSearch {
			return BinarySearchRecursive(arr, high, mid+1, noToSearch)
		} else {
			return BinarySearchRecursive(arr, mid-1, low, noToSearch)
		}
	} else {
		return -1
	}
}
