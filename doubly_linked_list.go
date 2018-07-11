package main

import "fmt"

// LinkedListNode does stuff
type LinkedListNode struct {
	value    int
	next     *LinkedListNode
	previous *LinkedListNode
}

// LinkedList does other stuff
type LinkedList struct {
	head *LinkedListNode
	tail *LinkedListNode
}

func append(node *LinkedListNode, list *LinkedList) *LinkedList {
	if list == nil {
		list = &LinkedList{node, node}
		list.head.next = node
		list.tail.previous = node
		list.head.previous = nil
	} else {
		list.tail.next = node
		node.previous = list.tail
		list.tail = node
	}

	return list
}

func prepend(node *LinkedListNode, list *LinkedList) *LinkedList {
	if list == nil {
		list = &LinkedList{node, node}
		list.head.next = node
		list.tail.previous = node
		list.head.previous = nil
	} else {
		list.head.previous = node
		node.next = list.head
		list.head = node
	}

	return list
}

func printForward(node *LinkedListNode) *LinkedListNode {
	if node == nil {
		return nil
	}

	fmt.Print(node.value)

	if node.next != nil {
		fmt.Print(" -> ")
	}
	return printForward(node.next)
}

func printBackward(node *LinkedListNode) *LinkedListNode {
	if node == nil {
		return nil
	}

	fmt.Print(node.value)

	if node.previous != nil {
		fmt.Print(" -> ")
	}
	return printBackward(node.previous)
}

func main() {
	var list *LinkedList
	node1 := LinkedListNode{1, nil, nil}
	node2 := LinkedListNode{2, nil, nil}
	node3 := LinkedListNode{3, nil, nil}
	node4 := LinkedListNode{4, nil, nil}
	node5 := LinkedListNode{0, nil, nil}

	list = append(&node1, list)
	list = append(&node2, list)
	list = append(&node3, list)
	list = append(&node4, list)
	list = prepend(&node5, list)

	fmt.Println("Forward")
	printForward(list.head)

	fmt.Println("\nBackward")
	printBackward(list.tail)
}
