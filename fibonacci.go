package main

import (
	"fmt"
	"strings"
);

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	index := 0
	last := 0
	secondToLast := 1
	fn := func() int {
		var ret int
		if (index == 0) {
			ret = 0
		} else if (index == 1) {
			ret = 1
		} else {
			ret = last + secondToLast
		}
		secondToLast = last
		last = ret
		index++
		return ret
	}
	return fn
}

func Fib(num int) string {
	var seq strings.Builder
	f := fibonacci()
	for i := 0; i < num; i++ {
		seq.WriteString(fmt.Sprint(f()) + " ")
	}
	return seq.String()
}
