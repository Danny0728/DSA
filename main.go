package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	algo "github.com/Danny0728/GrokkingAlgorithms/algorithms"
)

func main() {
	var (
		startTime time.Time
		a         = []int{13, 2, 3, 10, 5, 6, 7, 18, 22, 0, 8, 99, 1}
	)
	selectionSortFlag := flag.Bool("selectionSort", false, "Run algorithms/SelectionSort.go ")
	divideandConquerFlag := flag.Bool("d&C", false, "Run algorithms/divideandconquer.go ")
	binarySearchFlag := flag.Bool("binarySearch", false, "Run algorithms/binarySearch.go ")
	quickSortFlag := flag.Bool("quickSort", false, "Run algorithms/SelectionSort.go ")

	flag.Parse()

	if *selectionSortFlag {
		startTime = time.Now()
		// arr := []int{64, 25, 12, 22, 11}
		fmt.Println("Unsorted array:", a)
		algo.SelectionSort(a)
		fmt.Println("Sorted array:", a)
		log.Println("Total execution Time of SelectionSort:", time.Since(startTime))
	}

	if *divideandConquerFlag {
		// using recursive function add all the elements in the array and return the total
		startTime = time.Now()
		fmt.Println("AddNums using recursive function:", algo.AddNums(a, len(a)-1))
		fmt.Println("GetMax No using recursive function:", algo.GetMaxNo(a, len(a)-1, a[len(a)-1]))
		fmt.Println("CountItems using recursive function:", algo.CountItems(a, len(a)-1))

		fmt.Println("BinarySearch using recursive function:", algo.BinarySearchRecursive(a, len(a)-1, 0, 10))
		log.Println("Total execution Time of Binary Search:", time.Since(startTime))
	}

	if *binarySearchFlag {
		sortedArray := []int{1, 2, 5, 6, 7, 80, 90, 1000, 1100}
		numberToSearch := 1100
		output := algo.BinarySearch(sortedArray, numberToSearch)
		fmt.Println(output)

		sortedAr := []string{"apple", "appy", "appyz", "orange", "pineapple", "strawberry"}
		stringToSearch := "orange"
		result := algo.BinarySearch(sortedAr, stringToSearch)
		fmt.Println(result)
	}

	if *quickSortFlag {
		startTime = time.Now()
		fmt.Println("QuickSort:", algo.QuickSort(a))
		log.Println("Total execution Time of QuickSort:", time.Since(startTime))

		startTime = time.Now()
		algo.InPlaceQuickSort(a, 0, len(a)-1)
		log.Println("Total execution Time of InPlaceQuickSort:", time.Since(startTime))
	}
}
