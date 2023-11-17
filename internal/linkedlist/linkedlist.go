package Linkedlist

func (l *Linkedlist) CircularAdd(val int) {
	newNode := &Node{val: val}

	if l.head == nil {
		l.head = newNode
		l.head.next = l.head // Pointing to itself
	} else {
		newNode.next = l.head // New node points to head
		l.tail.next = newNode // Previous tail points to new node
	}
	l.tail = newNode // Update the tail to be the new node
}

func (l *Linkedlist) CircularDelete(val int) {
	if l.head == nil {
		// The list is empty
		return
	}

	current := l.head
	var prev *Node = nil

	// If the list has only one node
	if current.next == l.head && current.val == val {
		l.head = nil
		return
	}

	// Searching for the node to delete
	for {
		if current.val == val {
			if current == l.head {
				// Node to delete is the head
				l.head = current.next
				l.tail.next = l.head
			} else {
				// Node is either a middle node or the tail
				prev.next = current.next
				if current == l.tail {
					// Node is the tail
					l.tail = prev
				}
			}
			return
		}

		prev = current
		current = current.next

		// If we have traversed the entire list and are back at the head
		if current == l.head {
			break
		}
	}
}
