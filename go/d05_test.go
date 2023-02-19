package main

import (
	"fmt"
	"testing"
)

func TestD05(t *testing.T) {
	expected := "CMZ"
	got := d05_Part1("../data/d05_example.txt")
	if got != expected {
		t.Error(fmt.Sprintf("Part 1 failed: expect %s, but got: %s", expected, got))
	}

	/*
		expected = ""
		got = d03_Part2("../data/d05_example.txt")
		if got != expected {
			t.Error(fmt.Sprintf("Part 2 failed: expect %s, but got: %s", expected, got))
		}
	*/
}
