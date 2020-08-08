package main

import (
	"code.google.com/p/go-tour/tree"
	"fmt"
)


/* Statement: Given two sorted binary trees, we should deduce whether
they are equal or not. Two binary trees are said to be equal if they have 
the same multiset of values. 

	Methods: tree.New(k) returns a randomly-structured sorted binary tree with values k, 2k, ... 10k*/



/* if a binary tree is sorted then its inorder travelsal considers nodes
in a sorted (ascending) order. */
func WalkHelper(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}

	WalkHelper(t.Left, ch)
	ch <- t.Value
	WalkHelper(t.Right, ch)
}


func Walk(t *tree.Tree, ch chan int) {
	defer close(ch)
	WalkHelper(t, ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	channel1, channel2 := make(chan int), make(chan int)
	go Walk(t1, channel1)
	go Walk(t2, channel2)
	
	for {
		result1, status1 := <- channel1
		result2, status2 := <- channel2
		if result1 != result2 || status1 != status2 {
			return false
		}

		if !status1 {
			break
		}
	}

	return true
}

func main() {
	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(1), tree.New(2)))
}
