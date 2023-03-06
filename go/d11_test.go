package main

import (
	"testing"
)

func TestD11_Part1(t *testing.T) {
	expected := 10605
	got := d11_Part1("../data/d11_example.txt")
	if got != expected {
		t.Errorf("Day 11 Part 1 failed, expect answer: %d, but got: %d", expected, got)
	}

}
