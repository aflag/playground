package main

import (
	"fmt"
)

const (
	GoLeft = iota
	GoRight
	Return
	Read
)

type BinaryTree struct {
	Value int
	Left  *BinaryTree
	Right *BinaryTree
}

func (b *BinaryTree) String() string {
	result := ""
	if b.Left != nil {
		result = b.Left.String() + " "
	}
	result += fmt.Sprintf("%d", b.Value)
	if b.Right != nil {
		result += " " + b.Right.String()
	}
	return result
}

func deserialize(ptr int, data []int) (int, *BinaryTree) {
	tree := new(BinaryTree)

	for ptr < len(data) {
		switch cmd := data[ptr]; cmd {
		case GoLeft:
			ptr, tree.Left = deserialize(ptr+1, data)
		case GoRight:
			ptr, tree.Right = deserialize(ptr+1, data)
		case Read:
			ptr++
			tree.Value = data[ptr]
			ptr++
		case Return:
			ptr++
			return ptr, tree
		}
	}

	return ptr, tree
}

func Deserialize(data []int) *BinaryTree {
	_, t := deserialize(0, data)
	return t
}

func (b *BinaryTree) serialize(data *[]int) {
	if b.Left != nil {
		*data = append(*data, GoLeft)
		b.Left.serialize(data)
	}
	*data = append(*data, Read)
	*data = append(*data, b.Value)
	if b.Right != nil {
		*data = append(*data, GoRight)
		b.Right.serialize(data)
	}
	*data = append(*data, Return)
}

// size complexity O(5*N), where N = number of nodes
// time complexity O(N) (amortized)
func (b *BinaryTree) Serialize() []int {
	data := new([]int)
	b.serialize(data)
	return *data
}
