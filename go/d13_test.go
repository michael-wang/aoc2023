package main

import (
	"testing"
)

func TestD13_Parse(t *testing.T) {
	row := "[1,1,3,1,1]"
	packet := &list{}
	packet.Parse(row)
	got := packet.String()
	if got != row {
		t.Errorf("Expected: %s\nGot:    %s", row, got)
	}

	row = "[[1],[2,3,4]]"
	packet.Parse(row)
	got = packet.String()
	if got != row {
		t.Errorf("Expected: %s\nGot:    %s", row, got)
	}

	row = "[[[]]]"
	packet.Parse(row)
	got = packet.String()
	if got != row {
		t.Errorf("Expected: %s\nGot:    %s", row, got)
	}

	row = "[1,[2,[3,[4,[5,6,7]]]],8,9]"
	packet.Parse(row)
	got = packet.String()
	if got != row {
		t.Errorf("Expected: %s\nGot:    %s", row, got)
	}

	row = "[[[],6,4,[1,8,7,4],1],[4,8,[[],4,[10,2,7,2],10,[]]],[[[10,7,0],10],[10]]]"
	packet.Parse(row)
	got = packet.String()
	if got != row {
		t.Errorf("Expected: %s\nGot:    %s", row, got)
	}
}

func TestD13_Load(t *testing.T) {
	expected := `[1,1,3,1,1]
[1,1,5,1,1]

[[1],[2,3,4]]
[[1],4]

[9]
[[8,7,6]]

[[4,4],4,4]
[[4,4],4,4,4]

[7,7,7,7]
[7,7,7]

[]
[3]

[[[]]]
[[]]

[1,[2,[3,[4,[5,6,7]]]],8,9]
[1,[2,[3,[4,[5,6,0]]]],8,9]
`

	pp := d13_Load("../data/d13_example.txt")
	got, sep := "", ""
	for _, pair := range pp {
		got += sep
		sep = "\n"

		got += pair.X.(*list).String()
		got += "\n"
		got += pair.Y.(*list).String()
		got += "\n"
	}

	if got != expected {
		for i, b := range []byte(expected) {
			if got[i] != b {
				t.Errorf("Expected:\n%s\n%x\nGot:\n%s\n%x", expected[:i], expected[i], got[:i], got[i])
				break
			}
		}
	}
}

func TestD13_CorrectOrder(t *testing.T) {
	left := &list{}
	left.Parse("[1,1,3,1,1]")
	right := &list{}
	right.Parse("[1,1,5,1,1]")
	expected := 1
	got := left.CorrectOrder(right, 0)
	if expected != got {
		t.Errorf("%s vs %s, expected: %d, got: %d", left.String(), right.String(), expected, got)
	}

	left.Parse("[[1],[2,3,4]]")
	right.Parse("[[1],4]")
	expected = 1
	got = left.CorrectOrder(right, 0)
	if expected != got {
		t.Errorf("%s vs %s, expected: %d, got: %d", left.String(), right.String(), expected, got)
	}

	left.Parse("[9]")
	right.Parse("[[8,7,6]]")
	expected = -1
	got = left.CorrectOrder(right, 0)
	if expected != got {
		t.Errorf("%s vs %s, expected: %d, got: %d", left.String(), right.String(), expected, got)
	}

	left.Parse("[[4,4],4,4]")
	right.Parse("[[4,4],4,4,4]")
	expected = 1
	got = left.CorrectOrder(right, 0)
	if expected != got {
		t.Errorf("%s vs %s, expected: %d, got: %d", left.String(), right.String(), expected, got)
	}

	left.Parse("[7,7,7,7]")
	right.Parse("[7,7,7]")
	expected = -1
	got = left.CorrectOrder(right, 0)
	if expected != got {
		t.Errorf("%s vs %s, expected: %d, got: %d", left.String(), right.String(), expected, got)
	}

	left.Parse("[]")
	right.Parse("[3]")
	expected = 1
	got = left.CorrectOrder(right, 0)
	if expected != got {
		t.Errorf("%s vs %s, expected: %d, got: %d", left.String(), right.String(), expected, got)
	}

	left.Parse("[[[]]]")
	right.Parse("[[]]")
	expected = -1
	got = left.CorrectOrder(right, 0)
	if expected != got {
		t.Errorf("%s vs %s, expected: %d, got: %d", left.String(), right.String(), expected, got)
	}

	left.Parse("[1,[2,[3,[4,[5,6,7]]]],8,9]")
	right.Parse("[1,[2,[3,[4,[5,6,0]]]],8,9]")
	expected = -1
	got = left.CorrectOrder(right, 0)
	if expected != got {
		t.Errorf("%s vs %s, expected: %d, got: %d", left.String(), right.String(), expected, got)
	}

	left.Parse("[7,7,7]")
	right.Parse("[7,7,7]")
	expected = 0
	got = left.CorrectOrder(right, 0)
	if expected != got {
		t.Errorf("%s vs %s, expected: %d, got: %d", left.String(), right.String(), expected, got)
	}
}

func TestD13_Part1(t *testing.T) {
	expected := 13
	got := d13_Part1("../data/d13_example.txt")
	if got != expected {
		t.Errorf("Day 13 Part 1, expected: %d, got: %d", expected, got)
	}
}
