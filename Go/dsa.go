package main

import "fmt"

var count = 0

// ---Binary Search Tree Data Structure---
type BST struct{
	key int
	left *BST
	right *BST
}

// Insert function: Arranging values in the tree based on the key
// If value < key, move to the left else move to the right
func (n *BST) Insert (k int){
	if n.key < k{
		if n.right == nil{
			n.right = &BST{key: k}
		} else {
			n.right.Insert(k)
		} 
	} else if n.key > k {
		if n.left == nil{
			n.left = &BST{key: k}
		} else {
			n.left.Insert(k)
		} 		
	}
}

// Follows similar principle to Insertion,
// smaller than key? check left side,
// larger than key? check right side and repeat.
// If there's nothing and answer still hasn't been
// found, it's not there at all.
func (n *BST) Search (k int) bool{
	count++

	if n == nil{
		return false
	}

	if n.key < k{
		// value is bigger so move right
		return n.right.Search(k)
	} else if n.key > k{
		// value is smaller so move left
		return n.left.Search(k)
	}

	return true
}



func main(){

	bst := &BST{key:100}

	bst.Insert(20)
	bst.Insert(19)
	bst.Insert(30)
	bst.Insert(250)
	bst.Insert(120)
	bst.Insert(75)

	fmt.Println(bst.Search(120))
	fmt.Println("It takes",count, "steps")
}