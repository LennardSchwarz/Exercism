package binarysearchtree

type BinarySearchTree struct {
	left  *BinarySearchTree
	data  int
	right *BinarySearchTree
}

// NewBst creates and returns a new BinarySearchTree.
func NewBst(i int) *BinarySearchTree {
	return &BinarySearchTree{data: i}
}

// Insert inserts an int into the BinarySearchTree.
// Inserts happen based on the rules of a binary search tree
func (bst *BinarySearchTree) Insert(i int) {
	if i <= bst.data {
		if bst.left == nil {
			bst.left = NewBst(i)
		} else {
			bst.left.Insert(i)
		}
	} else {
		if bst.right == nil {
			bst.right = NewBst(i)
		} else {
			bst.right.Insert(i)
		}
	}
}

// SortedData returns the ordered contents of BinarySearchTree as an []int.
// The values are in increasing order starting with the lowest int value.
// A BinarySearchTree that has the numbers [1,3,7,5] added will return the
// []int [1,3,5,7].
func (bst *BinarySearchTree) SortedData() []int {
	// Check if bst is empty
	if bst == nil {
		return nil
	}

	var sortedData []int

	// Traverse the left subtree and collect its data
	if bst.left != nil {
		sortedData = append(sortedData, bst.left.SortedData()...)
	}

	// Append the current node's data
	sortedData = append(sortedData, bst.data)

	// Traverse the right subtree and collect its data
	if bst.right != nil {
		sortedData = append(sortedData, bst.right.SortedData()...)
	}

	return sortedData
}
