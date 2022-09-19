package main

import (
	"fmt"
	"strings"
)

type node struct {
	Value []byte
	Left  *node
	Right *node
}

type bstWrapper struct {
	Root *node
	Len  int
}

func (n node) String() string {
	return fmt.Sprintf("%v", n.Value)
}

func NewBstWrapper(firstValue []byte) *bstWrapper {
	newBst := bstWrapper{
		Root: &node{Value: firstValue, Left: nil, Right: nil},
		Len:  1,
	}

	return &newBst
}

func (bst bstWrapper) String() string {
	sb := strings.Builder{}
	bst.inOrderTraversal(&sb)
	return sb.String()
}

func (bst bstWrapper) inOrderTraversal(sb *strings.Builder) {
	bst.inOrderTraversalByNode(sb, bst.Root)
}

func (bst bstWrapper) inOrderTraversalByNode(sb *strings.Builder, root *node) {
	if root == nil {
		return
	}

	bst.inOrderTraversalByNode(sb, root.Left)
	sb.WriteString(fmt.Sprintf("%s ", root))
	bst.inOrderTraversalByNode(sb, root.Right)
}

// Add() addes a value to our BST
// *bstWrapper pointer because we are making changes to it
func (bst *bstWrapper) Add(value []byte) {
	bst.Root = bst.addByNode(bst.Root, value)
	bst.Len++
}

func (bst *bstWrapper) addByNode(root *node, value []byte) *node {
	if root == nil {
		return &node{Value: value}
	}

	if len(value) < len(root.Value) {
		root.Left = bst.addByNode(root.Left, value)
	} else {
		root.Right = bst.addByNode(root.Right, value)
	}

	return root
}

// Remove() removes a value in our BST
func (bst *bstWrapper) Remove(value []byte) {
	bst.removeByNode(bst.Root, value)
}

func (bst *bstWrapper) removeByNode(root *node, value []byte) *node {
	if root == nil {
		return root
	}

	if len(value) > len(root.Value) {
		root.Right = bst.removeByNode(root.Right, value)
	} else if len(value) < len(root.Value) {
		root.Left = bst.removeByNode(root.Left, value)
	} else {
		if root.Left == nil {
			return root.Right
		} else {
			tmpPointer := root.Left

			for tmpPointer.Right != nil {
				tmpPointer = tmpPointer.Right
			}

			root.Value = tmpPointer.Value

			root.Left = bst.removeByNode(root.Left, value)
		}
	}

	return root
}

// Search() searches for a value in our BST
func (bst bstWrapper) Search(value []byte) *node {
	return bst.searchByNode(bst.Root, value)
}

func (bst bstWrapper) searchByNode(root *node, value []byte) *node {
	if root == nil {
		return nil
	}

	if len(value) == len(root.Value) {
		return root
	} else if len(value) < len(root.Value) {
		return bst.searchByNode(root.Left, value)
	} else {
		return bst.searchByNode(root.Right, value)
	}
}
