package main

import (
	"testing"
)

func TestD14_Load(t *testing.T) {
	got := d14_Load("../data/d14_example.txt")
	exp := []path{
		{
			vec2{498, 4},
			vec2{498, 5},
			vec2{498, 6},
			vec2{497, 6},
			vec2{496, 6},
		},
		{
			vec2{503, 4},
			vec2{502, 4},
			vec2{502, 5},
			vec2{502, 6},
			vec2{502, 7},
			vec2{502, 8},
			vec2{502, 9},
			vec2{501, 9},
			vec2{500, 9},
			vec2{499, 9},
			vec2{498, 9},
			vec2{497, 9},
			vec2{496, 9},
			vec2{495, 9},
			vec2{494, 9},
		},
	}
	for i, p := range exp {
		if !checkPath(p, got[i], t) {
			return
		}
	}
}

func checkPath(exp, got path, t *testing.T) bool {
	if len(exp) != len(got) {
		t.Errorf("len(p): %d != len(q): %d", len(exp), len(got))
		return false
	}
	for i, v := range exp {
		if v != got[i] {
			t.Errorf("exp[%d]: %v != got[%d]: %v", i, v, i, got[i])
			return false
		}
	}
	return true
}

func TestD14_Part1(t *testing.T) {
	exp := 24
	got := d14_Part1("../data/d14_example.txt")
	if got != exp {
		t.Errorf("Part 1 example expected: %d, got: %d", exp, got)
	}

	exp = 755
	got = d14_Part1("../data/d14.txt")
	if got != exp {
		t.Errorf("Part 1 expected: %d, got: %d", exp, got)
	}
}

func TestD14_Part2(t *testing.T) {
	exp := 93
	got := d14_Part2("../data/d14_example.txt")
	if got != exp {
		t.Errorf("Part 2 example expected: %d, got: %d", exp, got)
	}

	exp = 29805
	got = d14_Part2("../data/d14.txt")
	if got != exp {
		t.Errorf("Part 2 expected: %d, got: %d", exp, got)
	}
}
