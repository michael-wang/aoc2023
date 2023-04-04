package main

import (
	"fmt"
	"os"
)

func slicePop(s [][]int, i int) [][]int {
	return append(s[:i], s[i+1:]...)
}

func intSlicePop(s []int, i int) (val int, s2 []int) {
	val = s[i]
	s2 = append(s[:i], s[i+1:]...)
	return
}

func intSlicePopMax(s []int) (max int, s2 []int) {
	if len(s) == 0 {
		panic("Cannot find max of zero length slcie")
	}
	max_i := 0
	max = s[max_i]
	for i := 1; i < len(s); i++ {
		if s[i] > max {
			max_i = i
			max = s[i]
		}
	}
	_, s2 = intSlicePop(s, max_i)
	return
}

func shortestString(ss []string) (index int) {
	if len(ss) == 0 {
		panic("Expect at least one element, but got empty slice")
	}

	min := len(ss[0])
	for i := 1; i < len(ss); i++ {
		curr := len(ss[i])
		if curr < min {
			min = curr
			index = i
		}
	}
	return
}

func pause(msg string) {
	fmt.Println(msg)
	os.Stdin.Read(make([]byte, 1))
}

type mtx22 [][]int

func (m mtx22) Copy(value int) (out mtx22) {
	out = make(mtx22, len(m))
	for y := range m {
		out[y] = make([]int, len(m[y]))
		for x := range m[y] {
			out[y][x] = value
		}
	}
	return
}

func (m mtx22) Print() {
	for y := range m {
		fmt.Println(m[y])
	}
}

func (m mtx22) DeepCopy() (out mtx22) {
	out = make(mtx22, len(m))
	for y := range m {
		out[y] = make([]int, len(m[y]))
		copy(out[y], m[y])
	}
	return out
}

type vec2 struct {
	X, Y interface{}
}
