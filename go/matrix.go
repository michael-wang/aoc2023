package main

import "fmt"

type mtx [][]int

// Return m by n matrix.
// Notice it's X major.
func makeMatrix(m, n int) (out mtx) {
	for x := 0; x < m; x++ {
		out = append(out, make([]int, n))
	}
	return
}

func (m mtx) Copy(value int) (out mtx) {
	out = make(mtx, len(m))
	for y := range m {
		out[y] = make([]int, len(m[y]))
		for x := range m[y] {
			out[y][x] = value
		}
	}
	return
}

func (m mtx) Print() {
	for y := range m {
		fmt.Println(m[y])
	}
}

func (m mtx) DeepCopy() (out mtx) {
	out = make(mtx, len(m))
	for y := range m {
		out[y] = make([]int, len(m[y]))
		copy(out[y], m[y])
	}
	return out
}
