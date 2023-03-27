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

func stringsReplaceByte(s string, i int, c byte) string {
	out := []byte(s)
	out[i] = c
	return string(out)
}

type vec2 struct {
	X int
	Y int
}

func (v vec2) Copy() vec2 {
	return vec2{
		X: v.X,
		Y: v.Y,
	}
}

func (v vec2) Equals(other vec2) bool {
	return v.X == other.X && v.Y == other.Y
}

func (v vec2) ToString() string {
	return fmt.Sprintf("(%d, %d)", v.X, v.Y)
}

type sliceS []string

func (s sliceS) Equals(other sliceS) bool {
	if len(s) != len(other) {
		return false
	}
	for i := 0; i < len(s); i++ {
		if s[i] != other[i] {
			return false
		}
	}
	return true
}

func (s sliceS) DeepCopy() sliceS {
	t := make([]string, len(s))
	copy(t, s)
	return t
}
