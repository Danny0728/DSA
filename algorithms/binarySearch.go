package algorithms

func BinarySearch(arr interface{}, item interface{}) int {
	switch arr := arr.(type) {
	case []int:
		return BinarySearchForInt(arr, item.(int))
	case []string:
		return BinarySearchforString(arr, item.(string))
	default:
		return -1 // Unsupported type
	}
}

// array of int and one item tocheck
func BinarySearchForInt(arr []int, numberToCheck int) int {
	low := 0
	high := len(arr) - 1

	for low <= high { //not narrowed down to single element
		mid := (low + high) / 2
		guess := arr[mid]
		if guess == numberToCheck {
			return mid
		}
		if guess < numberToCheck {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1 // return -1 if item doesnt exists
}
func BinarySearchforString(arr []string, numberToCheck string) int {
	low := 0
	high := len(arr) - 1

	for low <= high { //not narrowed down to single element
		mid := (low + high) / 2
		guess := arr[mid]
		if guess == numberToCheck {
			return mid
		}
		if guess < numberToCheck {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1 // return -1 if item doesnt exists
}
