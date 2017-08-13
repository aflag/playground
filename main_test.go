package main

import (
	"reflect"
	"testing"
)

func TestToString(t *testing.T) {
	tree := BinaryTree{
		Value: 5,
		Left: &BinaryTree{
			Value: 3,
			Left: &BinaryTree{
				Value: 2,
				Left:  &BinaryTree{Value: 1},
			},
			Right: &BinaryTree{Value: 4},
		},
		Right: &BinaryTree{Value: 6},
	}
	if tree.String() != "1 2 3 4 5 6" {
		t.Errorf("expected: 1 2 3 4 5 6, got: %s", tree.String())
	}
}

func TestDeserialize(t *testing.T) {
	data := []int{
		GoLeft,
		GoLeft,
		GoLeft,
		Read,
		1,
		Return,
		Read,
		2,
		Return,
		Read,
		3,
		GoRight,
		Read,
		4,
		Return,
		Return,
		Read,
		5,
		GoRight,
		Read,
		6,
		Return,
		Return,
	}

	output := Deserialize(data).String()

	if output != "1 2 3 4 5 6" {
		t.Errorf("expected: 1 2 3 4 5 6, got: %s", output)
	}
}

func TestSerialize(t *testing.T) {
	data := []int{
		GoLeft,
		GoLeft,
		GoLeft,
		Read,
		1,
		Return,
		Read,
		2,
		Return,
		Read,
		3,
		GoRight,
		Read,
		4,
		Return,
		Return,
		Read,
		5,
		GoRight,
		Read,
		6,
		Return,
		Return,
	}
	tree := BinaryTree{
		Value: 5,
		Left: &BinaryTree{
			Value: 3,
			Left: &BinaryTree{
				Value: 2,
				Left:  &BinaryTree{Value: 1},
			},
			Right: &BinaryTree{Value: 4},
		},
		Right: &BinaryTree{Value: 6},
	}

	serialized := tree.Serialize()

	if !reflect.DeepEqual(serialized, data) {
		t.Errorf("expected: %v, got: %v", data, serialized)
	}
}

func TestSerializeDeserialize(t *testing.T) {
	tree := BinaryTree{
		Value: 5,
		Left: &BinaryTree{
			Value: 3,
			Left: &BinaryTree{
				Value: 2,
				Left:  &BinaryTree{Value: 1},
			},
			Right: &BinaryTree{Value: 4},
		},
		Right: &BinaryTree{Value: 6},
	}

	deserialized := Deserialize(tree.Serialize())

	if tree.String() != deserialized.String() {
		t.Errorf("expected: %s, got: %s", tree.String(), deserialized.String())
	}
}
