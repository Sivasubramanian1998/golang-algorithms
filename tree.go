package main

import "fmt"

var count int

//Node represents the components of a binary tree
type node struct {
	key   int
	Left  *node
	Right *node
}

//Insert will add the node in binary tree
func (n *node) insert(k int) {
	if n.key < k {
		//move right
		if n.Right == nil {
			n.Right = &node{key: k}
		} else {
			n.Right.insert(k)
		}
	} else if n.key > k {
		//move left
		if n.Left == nil {
			n.Left = &node{key: k}
		} else {
			n.Left.insert(k)
		}
	}
}

//Search will take in a key value
//and return true if there is a node with that value
func (n *node) search(k int) bool {

	count++

	if n == nil {
		return false
	}
	if n.key < k {
		//move right
		return n.Right.search(k)
	} else if n.key > k {
		//move left
		return n.Left.search(k)
	}

	return true
}

func main() {
	tree := &node{key: 500}
	tree.insert(600)
	tree.insert(700)
	fmt.Println(tree)
	tree.insert(123)
	tree.insert(345)
	tree.insert(639)
	tree.insert(243)
	tree.insert(546)
	tree.insert(735)
	tree.insert(408)
	tree.insert(467)

	fmt.Println(tree.search(735))

	fmt.Println(count)
}
