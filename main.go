package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	algo "github.com/Danny0728/GrokkingAlgorithms/algorithms"
	"github.com/Danny0728/GrokkingAlgorithms/dataStructure"
)

func main() {
	var (
		startTime time.Time
		a         = []int{13, 2, 3, 10, 5, 6, 7, 18, 22, 0, 8, 99, 1}
	)
	selectionSortFlag := flag.Bool("selectionSort", false, "Runs algorithms/SelectionSort.go ")
	divideandConquerFlag := flag.Bool("d&C", false, "Runs algorithms/divideandconquer.go ")
	binarySearchFlag := flag.Bool("binarySearch", false, "Runs algorithms/binarySearch.go ")
	quickSortFlag := flag.Bool("quickSort", false, "Runs algorithms/SelectionSort.go ")
	binaryTreeFlag := flag.String("binaryTree", "", "Specify the binary tree traversal: preorder,inorder,postorder,levelorder,all")
	bstFlag := flag.Bool("bst", false, "Runs dataStructure/bst.go ")

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
		fmt.Println("SqrtRoot of 25:", algo.BinarySearchFindSqrtRoot(25))
	}

	if *quickSortFlag {
		startTime = time.Now()
		fmt.Println("QuickSort:", algo.QuickSort(a))
		log.Println("Total execution Time of QuickSort:", time.Since(startTime))

		startTime = time.Now()
		algo.InPlaceQuickSort(a, len(a)-1, 0)
		log.Println("InPlaceQuickSort:", a)
		log.Println("Total execution Time of InPlaceQuickSort:", time.Since(startTime))
	}

	if *binaryTreeFlag != "" {
		nodes := []int{1, 2, 4, -1, -1, 5, -1, -1, 3, -1, 6, -1, -1}
		tree := dataStructure.NewBinaryTree(nodes)
		root := tree.BuildTree()

		//check if the subRoot is part of tree
		subRootNodes := []int{2, 4, -1, -1, 5, -1, -1}
		subTree := dataStructure.NewBinaryTree(subRootNodes)
		subRoot := subTree.BuildTree()

		fmt.Println("Is subRoot: ", root.IsSubTree(subRoot))
		fmt.Println("No of Nodes: ", root.CountNodes())
		fmt.Println("Sum of Nodes: ", root.Sum())
		fmt.Println("Height of the tree: ", root.Height())
		fmt.Println("Diameter of the Tree(Linear Approach - O(N^2)): ", root.DiameterOfTree())
		fmt.Println("Diameter of the Tree(Optimized Approach - O(N)): ", root.DiameterOfTreeOptimized().Diameter)
		fmt.Println("Sum of the Nodes At Kth level(depth): ", root.SumOfNodesAtKLevel(2))
		if root != nil {
			switch *binaryTreeFlag {
			case "preorder":
				fmt.Print("Pre-order traversal of the binary tree: ")
				root.PreOrderTraversal()
				fmt.Println()
			case "inorder":
				fmt.Print("In-order traversal of the binary tree: ")
				root.InOrderTraversal()
				fmt.Println()
			case "postorder":
				fmt.Print("Post-order traversal of the binary tree: ")
				root.PostOrderTraversal()
				fmt.Println()
			case "levelorder":
				fmt.Print("Level-order traversal of the binary tree: \n")
				root.LevelOrder()
				fmt.Println()
			case "all":
				fmt.Print("Pre-order traversal of the binary tree: ")
				root.PreOrderTraversal()
				fmt.Println()
				fmt.Print("In-order traversal of the binary tree: ")
				root.InOrderTraversal()
				fmt.Println()
				fmt.Print("Post-order traversal of the binary tree: ")
				root.PostOrderTraversal()
				fmt.Println()
			default:
				fmt.Println("Invalid binary tree traversal type.")
			}
		} else {
			fmt.Println("The tree is empty.")
		}
	}
	if *bstFlag {

		nodes := []int{10, 1, 3, 4, 2, 7, 6, 8, 11, 12, 13}
		bst := dataStructure.NewBinarySearchTree(nodes)
		root := bst.BuildBST()
		fmt.Print("In-order traversal of the binary tree: ")
		root.InOrderTraversalBST()
		fmt.Println()

		fmt.Println("Input key exists?: ", bst.Search(root, 1))

		root.DeleteNode(13)
		fmt.Print("In-order traversal of the binary tree(after deleting node): ")
		root.InOrderTraversalBST()
		fmt.Println()

		fmt.Print("PrintInRange of bst :")
		root.PrintInRange(2, 11)
		fmt.Println()

		fmt.Print("PrintPaths of bst : \n")
		root.PrintRootToLeafPaths([]int{})
		fmt.Println()
	}
}
