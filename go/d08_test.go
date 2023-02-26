package main

import (
	"fmt"
	"testing"
)

func TestD08(t *testing.T) {
	expected := 16 + 5
	got := d08_Part1("../data/d08_example.txt")
	if got != expected {
		t.Error(fmt.Sprintf("Part 1 failed, expect answer: %d, but got: %d", expected, got))
	}

	expX, expY := 2, 3
	gotX, gotY := d08_Part2("../data/d08_example.txt")
	if gotX != expX {
		t.Error(fmt.Sprintf("Part 2 failed, expect X: %d, but got X: %d", expX, gotX))
	}
	if gotY != expY {
		t.Error(fmt.Sprintf("Part 2 failed, expect Y: %d, but got Y: %d", expY, gotY))
	}
}
