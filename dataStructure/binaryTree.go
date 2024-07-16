package dataStructure

import (
	"container/list"
	"fmt"
	"math"
)

// Node represents a node in a binary tree.
type Node struct {
	data  int
	left  *Node
	right *Node
}

type TreeInfo struct {
	height, Diameter int
}

// BinaryTree represents a binary tree.
type BinaryTree struct {
	idx   int
	nodes []int
}

// NewBinaryTree initializes a new BinaryTree.
func NewBinaryTree(nodes []int) *BinaryTree {
	return &BinaryTree{idx: -1, nodes: nodes}
}

/*
BuildTree builds a binary tree from the given nodes.
*/
func (bt *BinaryTree) BuildTree() *Node {
	bt.idx++
	if bt.idx >= len(bt.nodes) || bt.nodes[bt.idx] == -1 {
		return nil
	}
	newNode := &Node{data: bt.nodes[bt.idx]}
	newNode.left = bt.BuildTree()
	newNode.right = bt.BuildTree()
	return newNode
}

/*
Pre order traversal of a binary tree.
Time Complexity - O(n)
Space Complexity - O(n)
*/
func (n *Node) PreOrderTraversal() { // Root, Left, Right
	if n != nil {
		fmt.Printf("%d ", n.data)
		n.left.PreOrderTraversal()
		n.right.PreOrderTraversal()
	} else {
		// -1 represents a nil node.
		// fmt.Printf("-1 ")
		return
	}
}

/*
In order traversal of a binary tree.
Time Complexity - O(n)
Space Complexity - O(n)
*/
func (n *Node) InOrderTraversal() { // Left, Root, Right
	if n != nil {
		n.left.InOrderTraversal()
		fmt.Printf("%d ", n.data)
		n.right.InOrderTraversal()
	} else {
		// -1 represents a nil node.
		// fmt.Printf("-1 ")
		return
	}
}

/*
Post order traversal of a binary tree.
Time Complexity - O(n)
Space Complexity - O(n)
*/
func (n *Node) PostOrderTraversal() { // Left, Right, Root
	if n != nil {
		n.left.PostOrderTraversal()
		n.right.PostOrderTraversal()
		fmt.Printf("%d ", n.data)
	} else {
		// -1 represents a nil node.
		// fmt.Printf("-1 ")
		return
	}
}

/*
Level order traversal of a binary tree.
Time Complexity - O(n)
Space Complexity - O(n)
*/
func (n *Node) LevelOrder() {
	if n == nil {
		return
	}
	queue := list.New()
	queue.PushBack(n)
	queue.PushBack(nil)
	for queue.Len() > 0 {
		currElement := queue.Remove(queue.Front())
		curr, ok := currElement.(*Node)
		if !ok {
			fmt.Println() //level marker to indicate the end of a level
			// queue empty
			if queue.Len() == 0 { // queue is empty
				break
			} else {
				queue.PushBack(nil) // if queue is not empty add another level marker
			}
		} else {
			fmt.Print(curr.data, " ")
			if curr.left != nil {
				queue.PushBack(curr.left)
			}
			if curr.right != nil {
				queue.PushBack(curr.right)
			}
		}
	}
}

/*
Count the number of nodes in the binary tree.
Time Complexity - O(n)
Space Complexity - O(n)
*/
func (n *Node) CountNodes() int {
	if n == nil {
		return 0
	}
	return 1 + n.left.CountNodes() + n.right.CountNodes()
}

/*
Sum of all the nodes in the binary tree.
Time Complexity - O(n)
Space Complexity - O(n)
*/
func (n *Node) Sum() int {
	if n == nil {
		return 0
	}
	return n.data + n.left.Sum() + n.right.Sum()
}

/*
Height of a binary tree is the number of edges in the longest path from the root node to the leaf node.
Time Complexity - O(n)
Space Complexity - O(n)
*/
func (n *Node) Height() int {
	if n == nil {
		return 0
	}
	leftHeight := n.left.Height()
	rightHeight := n.right.Height()
	if leftHeight > rightHeight {
		return 1 + leftHeight
	}
	return 1 + rightHeight
}

/*
Diameter of a binary tree is the number of nodes in the longest path between any two leaf nodes.
Time Complexity - O(n^2)
Space Complexity - O(n)
*/
func (n *Node) DiameterOfTree() float64 {
	if n == nil {
		return 0
	}
	leftHeight := n.left.Height()
	rightHeight := n.right.Height()

	leftDiameter := n.left.DiameterOfTree()
	rightDiameter := n.right.DiameterOfTree()

	return math.Max(float64(leftHeight+rightHeight+1), math.Max(float64(leftDiameter), float64(rightDiameter)))
}

/*
Diameter of a binary tree is the number of nodes in the longest path between any two leaf nodes.
Time Complexity - O(n)
Space Complexity - O(n)
*/
func (n *Node) DiameterOfTreeOptimized() *TreeInfo {
	if n == nil {
		return &TreeInfo{0, 0}
	}

	leftTreeInfo := n.left.DiameterOfTreeOptimized()
	rightTreeInfo := n.right.DiameterOfTreeOptimized()

	myHeight := 1 + int(math.Max(float64(leftTreeInfo.height), float64(rightTreeInfo.height)))

	diam1 := leftTreeInfo.height                            // left tree diameter
	diam2 := rightTreeInfo.height                           // right tree diameter
	diam3 := leftTreeInfo.height + rightTreeInfo.height + 1 // diameter through root

	myDiameter := int(math.Max(float64(diam3), math.Max(float64(diam1), float64(diam2))))
	return &TreeInfo{myHeight, myDiameter}
}

func (n *Node) IsIdentical(subRoot *Node) bool {
	//check if the root is null (leaf or root data is nil) and subRoot is nil
	if (n == nil) && (subRoot == nil) {
		//returns true cause nil indicates that the leaf of the node is found
		return true
	}
	//check whose leaf has occurred early(cause nil indicates leaf)
	if (n == nil) || (subRoot == nil) {
		return false
	}
	if n.data != subRoot.data { //checking if the root nodes are similar
		return false
	}
	return n.left.IsIdentical(subRoot.left) && n.right.IsIdentical(subRoot.right)
}
func (n *Node) IsSubTree(subRoot *Node) bool {
	//check for nil cause nil exists in every binary tree hence return true
	if subRoot == nil {
		return true
	}
	//check if root is nil so return false as nil cannot have any subtree
	if n == nil {
		return false
	}

	if n.IsIdentical(subRoot) {
		return true
	}

	// check if the inner left or right subtree of the main binary tree matches with subRoot
	return n.left.IsSubTree(subRoot) || n.right.IsSubTree(subRoot)
}

func (n *Node) SumOfNodesAtKLevel(level int) int {

	var (
		depth = 1
		sum   int
	)
	if n == nil {
		return 0
	}
	if level == 0 {
		return 0
	}

	queue := list.New()
	queue.PushBack(n)
	queue.PushBack(nil)
	for queue.Len() > 0 {
		currElement := queue.Remove(queue.Front())
		curr, ok := currElement.(*Node)
		if !ok {
			if queue.Len() == 0 { // queue is empty
				break
			} else {
				depth++
				//if the depth goes above level just break to avoid unnecessary iterations
				if depth > level {
					break
				}
				queue.PushBack(nil) // if queue is not empty add another level marker
			}
		} else {
			//inside else to avoid nil nodes
			if depth == level {
				sum += curr.data
			}
			if curr.left != nil {
				queue.PushBack(curr.left)
			}
			if curr.right != nil {
				queue.PushBack(curr.right)
			}
		}
	}
	return sum
}
