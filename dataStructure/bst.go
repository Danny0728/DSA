package dataStructure

import "fmt"

type BinarySearchTree struct {
	nodes []int
}

// NewBinarySearchTree initializes a BST.
func NewBinarySearchTree(nodes []int) *BinarySearchTree {
	return &BinarySearchTree{nodes: nodes}
}

func (bst *BinarySearchTree) BuildBST() *Node {
	if len(bst.nodes) == 0 {
		return nil
	}
	root := &Node{data: bst.nodes[0]}
	for _, value := range bst.nodes[1:] {
		bst.Insert(root, value)
	}
	return root
}

/* Insert a node in BST
Time Complexity - O(h) where h is the height of the tree
Space Complexity - O(h) where h is the height of the tree
*/
func (bst *BinarySearchTree) Insert(root *Node, data int) *Node {
	if root == nil {
		return &Node{data: data}
	}
	if data < root.data {
		root.left = bst.Insert(root.left, data)
	} else if data > root.data {
		root.right = bst.Insert(root.right, data)
	}
	return root
}

// inorder traversal for BST
func (n *Node) InOrderTraversalBST() {
	if n != nil {
		n.left.InOrderTraversalBST()
		fmt.Printf("%d ", n.data)
		n.right.InOrderTraversalBST()
	}
}

/*
Search a node in BST
Time Complexity - O(h) where h is the height of the tree
Space Complexity - O(h) where h is the height of the tree
*/
func (bst *BinarySearchTree) Search(root *Node, data int) bool {
	if root == nil {
		return false
	}
	if root.data == data {
		return true
	} else if data > root.data {
		return bst.Search(root.right, data)
	} else if data < root.data {
		return bst.Search(root.left, data)
	}
	return false
}

func (root *Node) DeleteNode(val int) *Node {
	if val > root.data {
		root.right = root.right.DeleteNode(val)
	} else if val < root.data {
		root.left = root.left.DeleteNode(val)
	} else { //root.data = val
		//Case 1 --> deleting Node has No Child
		if root.left == nil && root.right == nil {
			return nil
		}
		//Case 2 --> deleting Node has Atleast one child
		if root.left == nil {
			return root.right
		} else if root.right == nil {
			return root.left
		}
		//Case 3 --> deleting Node has Two Children
		inOrderSuccessor := root.right.InOrderSuccessor()
		root.data = inOrderSuccessor.data
		root.right = root.right.DeleteNode(inOrderSuccessor.data)
	}
	return root
}
func (root *Node) InOrderSuccessor() *Node {
	for root.left != nil {
		root = root.left
	}
	return root
}

func (root *Node) PrintInRange(x, y int) *Node {
	if root == nil {
		return nil
	}
	if root.data >= x && root.data <= y {
		root.left.PrintInRange(x, y)
		fmt.Print(root.data, " ")
		root.right.PrintInRange(x, y)
	} else if root.data >= y {
		root.left.PrintInRange(x, y)
	} else {
		root.right.PrintInRange(x, y)
	}
	return nil
}
func (root *Node) PrintRootToLeafPaths(path []int) {
	if root == nil {
		return
	}
	path = append(path, root.data)

	if root.left == nil && root.right == nil { //leaf node condition
		PrintPath(path)
	} else { //non-leaf node condition
		if root.left != nil {
			root.left.PrintRootToLeafPaths(path)
		}
		if root.right != nil {
			root.right.PrintRootToLeafPaths(path)
		}
	}
}

func PrintPath(path []int) {
	for i, value := range path {
		if i == len(path)-1 {
			fmt.Print(value)
		} else {
			fmt.Print(value, "->")
		}
	}
	fmt.Println()
}
