package dataStructure

import (
	"container/list"
	"fmt"
)

// Node represents a node in a binary tree.
type Node struct {
	data  int
	left  *Node
	right *Node
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

// BuildTree constructs the binary tree from the given nodes.
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
