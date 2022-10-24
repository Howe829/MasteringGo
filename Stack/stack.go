package main

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

var size = 0
var stack = new(Node)

func Push(v int) bool {
	if stack == nil {
		stack = &Node{v, nil}
		size = 1
		return true
	}
	temp := &Node{v, nil}
	temp.Next = stack
	stack = temp
	size++
	return true
}

func Pop(t *Node) (int, bool) {
	if size == 0 {
		return 0, false
	}
	if size == 1 {
		size = 0
		stack = nil
		return t.Value, true
	}

	stack = stack.Next
	size--
	return t.Value, true
}
func traverse(t *Node) {
	if size == 0 {
		fmt.Println("Empty Stack!")
		return
	}
	for t.Next != nil {
		fmt.Print(t.Value)
		fmt.Print("->")
		t = t.Next
	}
	fmt.Println()
}
func main() {
	Push(10)
	Push(5)
	Push(1)

	traverse(stack)
	value, _ := Pop(stack)
	fmt.Println(value)

}
