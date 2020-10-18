package main

import "fmt"

type node struct {
	data int
	next *node
}

type linkedlist struct {
	head   *node
	length int
}

func (l *linkedlist) prepend(n *node) {
	second := l.head
	l.head = n
	l.head.next = second
	l.length++
}

//printing values not the address
func (l linkedlist) printListData() {
	f := l.head
	for l.length != 0 {
		fmt.Println("%d ", f.data)
		f = f.next
		l.length--
	}
	fmt.Printf("\n")
}

// deleting a node from the list
func (l *linkedlist) deleteValue(value int) {
	if l.length == 0 {
		return
	}

	if l.head.data == value {
		l.head = l.head.next
		l.length--
		return
	}

	delete := l.head
	for delete.next.data != value {
		if delete.next.next == nil {
			return
		}

		delete = delete.next
	}
	delete.next = delete.next.next
	l.length--
}

func main() {
	myList := linkedlist{}
	node1 := &node{data: 100}
	node2 := &node{data: 200}
	node3 := &node{data: 300}
	node4 := &node{data: 400}
	node5 := &node{data: 500}
	node6 := &node{data: 600}
	myList.prepend(node1)
	myList.prepend(node2)
	myList.prepend(node3)
	myList.prepend(node4)
	myList.prepend(node5)
	myList.prepend(node6)
	myList.printListData()
	myList.deleteValue(400)
	myList.deleteValue(600)
	myList.printListData()
	emptyList := linkedlist{}
	emptyList.deleteValue(10)
}
