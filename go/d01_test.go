package main

import (
	"fmt"
	"testing"
)

func TestD01(t *testing.T) {
	expected := 24000
	got := d01_Part1("../data/d01_example.txt")
	if got != expected {
		t.Error(fmt.Sprintf("Part 1 failed: expect %d, but got: %d", expected, got))
	}

	expected = 45000
	got = d01_Part2("../data/d01_example.txt")
	if got != expected {
		t.Error(fmt.Sprintf("Part 2 failed: expect %d, but got: %d", expected, got))
	}
}
