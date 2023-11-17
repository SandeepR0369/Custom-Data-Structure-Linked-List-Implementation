package Linkedlist

import "fmt"

type Node struct {
	val  int
	next *Node
}

type Linkedlist struct {
	head   *Node
	tail   *Node
	length int
}

func (l *Linkedlist) Add(val int) {
	newNode := new(Node)
	newNode.val = val

	if l.head == nil {
		l.head = newNode
	} else {
		i := l.head
		for i.next != nil {
			i = i.next
		}
		i.next = newNode
	}
}

func (l *Linkedlist) AddFront(val int) {
	newNode := new(Node)
	newNode.val = val
	newNode.next = l.head
	l.head = newNode
}

func (l *Linkedlist) Update(val, x int) {
	current := l.head
	for current != nil {
		if current.val == val {
			current.val = x
			break
		}
		current = current.next
	}

}

/*
func (l *Linkedlist) Delete(val int) {
	current := l.head

	if current == nil {
		return
	}

	if current.val == val {
		current = current.next
		return
	}

	for current.next != nil && current.val != val {
		current = current.next
	}

	if current.next != nil {
		current.next = current.next.next
	}
}*/

func (l *Linkedlist) Delete(val int) {
	current := l.head

	// If the list is empty, return immediately
	if current == nil {
		return
	}

	// If the head is the node to be deleted
	if current.val == val {
		l.head = current.next
		return
	}

	// Find the node before the node to be deleted
	for current.next != nil && current.next.val != val {
		current = current.next
	}

	// If the next node is the node to delete
	if current.next != nil {
		current.next = current.next.next
	}
}

func (l *Linkedlist) Print() {
	current := l.head
	for current != nil {
		fmt.Print(current.val, " ")
		current = current.next
	}
	fmt.Println()
}
