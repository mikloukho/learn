// Односвязный список. Твой код, переведён из package main в пакет,
// чтобы покрывался тестами и не конфликтовал с другими main в репо.
package datastructures

import (
	"fmt"
)

type Node struct {
	data int
	next *Node
}

type List struct {
	head *Node
}

func (l *List) add(value int) {
	newNode := &Node{data: value}

	if l.head == nil {
		l.head = newNode
		return
	}

	curr := l.head
	for curr.next != nil {
		curr = curr.next
	}

	curr.next = newNode
}

func (l *List) remove(value int) {
	if l.head == nil {
		return
	}

	if l.head.data == value {
		l.head = l.head.next
		return
	}

	curr := l.head
	for curr.next != nil && curr.next.data != value {
		curr = curr.next
	}

	if curr.next != nil {
		curr.next = curr.next.next
	}
}

func printList(l *List) {
	curr := l.head
	for curr != nil {
		fmt.Printf("%d ", curr.data)
		curr = curr.next
	}
	fmt.Println()
}
