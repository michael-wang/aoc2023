package main

import (
	"fmt"
	"testing"
)

func TestD07(t *testing.T) {
	expected := 94853 + 584
	got := d07_Part1("../data/d07_example.txt")
	if got != expected {
		t.Error(fmt.Sprintf("Part 1 failed, expect answer: %d, but got: %d", expected, got))
	}

	expected = 24933642
	got = d07_Part2("../data/d07_example.txt")
	if got != expected {
		t.Error(fmt.Sprintf("Part 2 failed, expect answer: %d, but got: %d", expected, got))
	}
}
