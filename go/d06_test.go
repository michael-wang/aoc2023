package main

import (
	"testing"
)

func TestD06(t *testing.T) {
	expected := []int{7, 5, 6, 10, 11}
	got := d06_Part1("../data/d06_example.txt")
	for i, exp := range expected {
		if got[i] != exp {
			t.Errorf("Part 1 failed at line: %d, expect answer: %d, but got: %d", i, exp, got[i])
		}
	}

	expected = []int{19, 23, 23, 29, 26}
	got = d06_Part2("../data/d06_example_p2.txt")
	for i, exp := range expected {
		if got[i] != exp {
			t.Errorf("Part 1 failed at line: %d, expect answer: %d, but got: %d", i, exp, got[i])
		}
	}
}
