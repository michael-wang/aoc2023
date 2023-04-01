package main

import (
	"testing"
)

func TestD05(t *testing.T) {
	expected := "CMZ"
	got := d05_Part1("../data/d05_example.txt")
	if got != expected {
		t.Errorf("Part 1 failed: expect %s, but got: %s", expected, got)
	}

	expected = "MCD"
	got = d05_Part2("../data/d05_example.txt")
	if got != expected {
		t.Errorf("Part 2 failed: expect %s, but got: %s", expected, got)
	}
}
