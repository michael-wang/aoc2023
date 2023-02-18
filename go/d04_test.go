package main

import (
	"fmt"
	"testing"
)

func TestD04(t *testing.T) {
	exp1 := 2
	exp2 := 4
	got1, got2 := d04_Part1and2("../data/d04_example.txt")
	if got1 != exp1 {
		t.Error(fmt.Sprintf("Part 1 failed: expect %d, but got: %d", exp1, got1))
	}
	if got2 != exp2 {
		t.Error(fmt.Sprintf("Part 2 failed: expect %d, but got: %d", exp2, got2))
	}
}
