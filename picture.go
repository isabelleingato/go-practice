package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	dslice := make([][]uint8, dy)
	for i, _ := range dslice {
		dslice[i] = make([]uint8, dx)
		for j, _ := range dslice[i] {
			dslice[i][j] = uint8(i*j);
		}
	}
	return dslice
}

func main() {
	pic.Show(Pic)
}
