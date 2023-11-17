package main

import (
	linkedlist "Custom-Data-Structure-Linked-List-Implementation/internal/linkedlist"
	"fmt"
)

func main() {
	a := &linkedlist.Linkedlist{}
	a.Add(1)
	a.Add(2)
	a.Add(3)
	a.Add(4)
	a.Add(5)
	a.Add(7)
	a.Add(10)
	fmt.Println("Linked List: ")
	a.Add(6)
	a.Add(8)
	a.AddFront(98)
	a.Print()
	a.Update(4, 9)
	a.Print()
	a.AddFront(1)
	//a.Delete(5)
	a.Print()
}
