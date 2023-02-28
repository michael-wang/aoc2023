package main

import (
	"testing"
)

func TestD09(t *testing.T) {
	expected := 13
	got := d09_Part1("../data/d09_example.txt")
	if got != expected {
		t.Errorf("Part 1 failed, expect answer: %d, but got: %d", expected, got)
	}

	expected = 36
	got = d09_Part2("../data/d09_example2.txt")
	if got != expected {
		t.Errorf("Part 2 failed, expect: %d, but got: %d", expected, got)
	}
}
