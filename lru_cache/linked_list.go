package main

import "fmt"

//Node in the LinkedList
type Node struct {
	key   string
	value string
	prev  *Node
	next  *Node
}

//LinkedList struct
type LinkedList struct {
	root *Node
	end  *Node
}

//NewNode creates a NewNode
func NewNode(key string, value string, prev *Node, next *Node) *Node {
	return &Node{
		key:   key,
		value: value,
		prev:  prev,
		next:  next,
	}
}

func (l *LinkedList) print() {
	node := l.root
	fmt.Println(node)
	i := 0
	for {
		if node == nil {
			return
		}
		fmt.Printf("node number %v: %v\n", i, node.value)
		node = node.next
		i++
	}
}
