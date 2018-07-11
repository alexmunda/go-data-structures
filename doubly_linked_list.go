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

	list = append(&LinkedListNode{1, nil, nil}, list)
	list = append(&LinkedListNode{2, nil, nil}, list)
	list = append(&LinkedListNode{3, nil, nil}, list)
	list = append(&LinkedListNode{4, nil, nil}, list)
	list = prepend(&LinkedListNode{0, nil, nil}, list)

	fmt.Println("Forward")
	printForward(list.head)

	fmt.Println("\nBackward")
	printBackward(list.tail)
}
