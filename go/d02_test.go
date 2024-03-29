package main

import (
	"testing"
)

func TestD02(t *testing.T) {
	expected := 15
	got := d02_Part1("../data/d02_example.txt")
	if got != expected {
		t.Errorf("Part 1 failed: expect %d, but got: %d", expected, got)
	}

	expected = 12
	got = d02_Part2("../data/d02_example.txt")
	if got != expected {
		t.Errorf("Part 2 failed: expect %d, but got: %d", expected, got)
	}
}
