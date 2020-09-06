package main

import(
	"golang.org/x/tour/tree"
	"fmt"
);

type Result struct {
    val int
    close bool
}

func walk(t *tree.Tree, ch chan Result) int {
	if (t == nil) {
		return 0
	}
	// preorder traversal
	leftCount := walk(t.Left, ch)
	ch <- Result{t.Value, false}
	rightCount := walk(t.Right, ch)
	return 1 + leftCount + rightCount
}

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan Result) {
	count := walk(t, ch)
	fmt.Println(count) // debugging
	ch <- Result{0, true}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	c1 := make(chan Result)
	c2 := make(chan Result)
	go Walk(t1, c1)
	go Walk(t2, c2)
	for {
		v1 := <- c1
		v2 := <- c2
		fmt.Println(v1.val, v2.val) // debugging
		if (v1.close != v2.close || v1.val != v2.val) {
			return false
		}
		if (v1.close && v2.close) {
			return true
		}
	}
}