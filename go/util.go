package main

func slicePop(s [][]int, i int) [][]int {
	return append(s[:i], s[i+1:]...)
}
