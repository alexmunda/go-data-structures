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

func remove(val int, list *LinkedList) *LinkedList {
	cur := list.head

	for {
		if cur == nil {
			return list
		}

		if cur.value == val {
			if cur.previous == nil {
				list.head = cur.next
				list.head.previous = nil
				return list
			}

			if cur.next == nil {
				list.tail = cur.previous
				list.tail.next = nil
				return list
			}

			cur.previous.next = cur.next
			cur.next.previous = cur.previous
			return list
		}

		cur = cur.next
	}
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
	var list = append(&LinkedListNode{1, nil, nil}, nil)
	append(&LinkedListNode{2, nil, nil}, list)
	append(&LinkedListNode{3, nil, nil}, list)
	append(&LinkedListNode{4, nil, nil}, list)
	prepend(&LinkedListNode{0, nil, nil}, list)

	fmt.Println("Forward")
	printForward(list.head)

	fmt.Println("\nBackward")
	printBackward(list.tail)

	fmt.Println("\nRemove 1")
	remove(1, list)
	printForward(list.head)
}
