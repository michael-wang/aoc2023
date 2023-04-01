package main

import (
	"fmt"
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

func TestD12_Matrix2D_Copy(t *testing.T) {
	_, _, source := d12_Load("../data/d12_example.txt")
	value := 1
	clone := source.Copy(value)

	fmt.Println("source:")
	source.Print()
	fmt.Println("clone:")
	clone.Print()

	if len(source) != len(clone) {
		t.Errorf("len(source):%d != len(clone): %d", len(source), len(clone))
	}
	for y := range source {
		if len(source[y]) != len(clone[y]) {
			t.Errorf("y: %d, len(source row): %d != len(clone row): %d", y, len(source[y]), len(clone[y]))
		}
		for x := range clone[y] {
			if clone[y][x] != value {
				t.Errorf("y: %d, x: %d, clone[y][x]: %d != value: %d", y, x, clone[y][y], value)
			}
		}
	}
}
func TestD12_Matrix2D_DeepCopy(t *testing.T) {
	_, _, source := d12_Load("../data/d12_example.txt")
	clone := source.DeepCopy()

	expected := source[0][0]
	got := clone[0][0]

	source[0][0] = -1
	fmt.Println("source:")
	source.Print()
	fmt.Println("clone:")
	clone.Print()

	if got != expected {
		t.Errorf("Expect clone[0][0] not correlated with source[0][0] but it is")
	}
}
