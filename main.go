package main

import (
	"fmt"
	"golang.org/x/tour/wc"
	"golang.org/x/tour/pic"
);

func main() {
	wc.Test(WordCount)
	fmt.Println(Sqrt(2))
	pic.Show(Pic)
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}