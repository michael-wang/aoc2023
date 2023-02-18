package main

import (
	"fmt"
	"testing"
)

func TestD03(t *testing.T) {
	expected := 157
	got := d03_Part1("../data/d03_example.txt")
	if got != expected {
		t.Error(fmt.Sprintf("Part 1 failed: expect %d, but got: %d", expected, got))
	}

	expected = 70
	got = d03_Part2("../data/d03_example.txt")
	if got != expected {
		t.Error(fmt.Sprintf("Part 2 failed: expect %d, but got: %d", expected, got))
	}
}
