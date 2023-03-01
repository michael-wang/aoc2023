package main

import (
	"testing"
)

func TestD10(t *testing.T) {
	expected := 13140
	got := d10_Part1("../data/d10_example.txt")
	if got != expected {
		t.Errorf("Day 10 Part 1 failed, expect answer: %d, but got: %d", expected, got)
	}

	/*
		expected = 36
		got = d10_Part2("../data/d10.txt")
		if got != expected {
			t.Errorf("Day 10 Part 2 failed, expect: %d, but got: %d", expected, got)
		}
	*/
}
