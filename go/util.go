package main

import (
	"fmt"

	"github.com/golang-collections/collections/stack"
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

func newStacks(count int) (ss map[int]*stack.Stack) {
	ss = make(map[int]*stack.Stack, count)
	for i := 0; i < count; i++ {
		ss[i] = stack.New()
	}
	return
}

func printStack(s stack.Stack) {
	str := ""
	size := s.Len()
	for i := 0; i < size; i++ {
		str = fmt.Sprintf("%v %s", s.Pop(), str)
	}
	fmt.Println(str)
}

func stringStackReverse(s *stack.Stack) {
	t := []string{}
	for i := s.Len(); i >= 0; i-- {
		ele := s.Pop()
		fmt.Printf("%d %v\n", i, ele)
		t = append(t, ele.(string))
	}
	for i := len(t); i >= 0; i-- {
		s.Push(t[i])
	}
}

func copySliceOfString(src []string) []string {
	dst := make([]string, len(src))
	copy(dst, src)
	return dst
}

func intAbs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
