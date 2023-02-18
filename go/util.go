package main

import (
	"fmt"

	"github.com/golang-collections/collections/stack"
)

func slicePop(s [][]int, i int) [][]int {
	return append(s[:i], s[i+1:]...)
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
