package main

import (
	"testing"
)

func TestD12_Part1(t *testing.T) {
	expected := 31
	got := d12_Part1("../data/d12_example.txt")
	if got != expected {
		t.Errorf("Day 12 Part 1 failed, expect answer: %d, but got: %d", expected, got)
	}
}

func TestD12_MapParse(t *testing.T) {
	heightMap := []string{
		"aabqponm", // Replace 'S' by 'a'
		"abcryxxl",
		"accszzxk", // Replace 'E' by 'z'
		"acctuvwj",
		"abdefghi",
	}
	expS := point{X: 0, Y: 0, Height: 0, Dist: 0}
	expE := point{X: 5, Y: 2, Height: 'z' - 'a', Dist: 0}

	start, end, m := d12_Load("../data/d12_example.txt")

	if !expS.Equals(start) {
		t.Errorf("Expected start: %s, got: %s\n", expS.ToString(), start.ToString())
	}
	if !expE.Equals(end) {
		t.Errorf("Expected end: %s, got: %s\n", expE.ToString(), end.ToString())
	}
	for y := 0; y < len(heightMap); y++ {
		row := heightMap[y]
		for x := 0; x < len(row); x++ {
			h := int(row[x] - 'a')
			if h != m[y][x] {
				t.Errorf("heights map x: %d, y: %d, expected height: %d, but got: %d", x, y, h, m[y][x])
				return
			}
		}
	}
}
