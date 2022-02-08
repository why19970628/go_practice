package main

import "fmt"

type Tree struct {
	value string
	Left  *Tree
	Right *Tree
}

func solution(tree *Tree) {
	if tree != nil {
		fmt.Println(tree.value)
		solution(tree.Left)
		solution(tree.Right)
	}
}

func main() {
	t := &Tree{}
	t.value = "value"
	t.Left = &Tree{value: "left1"}
	t.Left.Left = &Tree{value: "left2"}

	t.Right = &Tree{value: "right1"}
	t.Right.Right = &Tree{value: "right2"}

	solution(t)
}
