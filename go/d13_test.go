package main

import (
	"testing"
)

func TestD13_Swap(t *testing.T) {
	row := "[1, 2, 3]"
	packet := &d13_Packet{}
	packet.Parse(row, 0, len(row))
	packet.Swap(0, 2)

	exp := d13_Packet{
		NestedList: NestedList{3, 2, 1},
	}
	if packet.Less(exp) != 0 {
		t.Errorf("Expect %v, got: %v", exp, packet)
	}

	row = "[[1], [2]]"
	packet.Clear()
	packet.Parse(row, 0, len(row))
	packet.Swap(0, 1)
	exp.Clear()
	expRow := "[[2], [1]]"
	exp.Parse(expRow, 0, len(expRow))
	if packet.Less(exp) != 0 {
		t.Errorf("Expect %v, got: %v", exp, packet)
	}
}
func TestD13_Less(t *testing.T) {
	left := &d13_Packet{}
	row := "[1,1,3,1,1]"
	left.Parse(row, 0, len(row))
	right := &d13_Packet{}
	row = "[1,1,5,1,1]"
	right.Parse(row, 0, len(row))
	expected := 1
	got := left.Less(*right)
	if expected != got {
		t.Errorf("%v vs %v, expected: %d, got: %d", left, right, expected, got)
	}

	row = "[[1],[2,3,4]]"
	left.Clear()
	left.Parse(row, 0, len(row))

	row = "[[1],4]"
	right.Clear()
	right.Parse(row, 0, len(row))
	expected = 1
	got = left.Less(*right)
	if expected != got {
		t.Errorf("%v vs %v, expected: %d, got: %d", left, right, expected, got)
	}

	row = "[9]"
	left.Clear()
	left.Parse(row, 0, len(row))

	row = "[[8,7,6]]"
	right.Clear()
	right.Parse(row, 0, len(row))
	expected = -1
	got = left.Less(*right)
	if expected != got {
		t.Errorf("%v vs %v, expected: %d, got: %d", left, right, expected, got)
	}

	row = "[[4,4],4,4]"
	left.Clear()
	left.Parse(row, 0, len(row))

	row = "[[4,4],4,4,4]"
	right.Clear()
	right.Parse(row, 0, len(row))
	expected = 1
	got = left.Less(*right)
	if expected != got {
		t.Errorf("%v vs %v, expected: %d, got: %d", left, right, expected, got)
	}

	row = "[7,7,7,7]"
	left.Clear()
	left.Parse(row, 0, len(row))

	row = "[7,7,7]"
	right.Clear()
	right.Parse(row, 0, len(row))
	expected = -1
	got = left.Less(*right)
	if expected != got {
		t.Errorf("%v vs %v, expected: %d, got: %d", left, right, expected, got)
	}

	row = "[]"
	left.Clear()
	left.Parse(row, 0, len(row))

	row = "[3]"
	right.Clear()
	right.Parse(row, 0, len(row))
	expected = 1
	got = left.Less(*right)
	if expected != got {
		t.Errorf("%v vs %v, expected: %d, got: %d", left, right, expected, got)
	}

	row = "[[[]]]"
	left.Clear()
	left.Parse(row, 0, len(row))

	row = "[[]]"
	right.Clear()
	right.Parse(row, 0, len(row))
	expected = -1
	got = left.Less(*right)
	if expected != got {
		t.Errorf("%v vs %v, expected: %d, got: %d", left, right, expected, got)
	}

	row = "[1,[2,[3,[4,[5,6,7]]]],8,9]"
	left.Clear()
	left.Parse(row, 0, len(row))

	row = "[1,[2,[3,[4,[5,6,0]]]],8,9]"
	right.Clear()
	right.Parse(row, 0, len(row))
	expected = -1
	got = left.Less(*right)
	if expected != got {
		t.Errorf("%v vs %v, expected: %d, got: %d", left, right, expected, got)
	}

	row = "[7,7,7]"
	left.Clear()
	left.Parse(row, 0, len(row))

	row = "[7,7,7]"
	right.Clear()
	right.Parse(row, 0, len(row))
	expected = 0
	got = left.Less(*right)
	if expected != got {
		t.Errorf("%v vs %v, expected: %d, got: %d", left, right, expected, got)
	}
}

func TestD13_Part1(t *testing.T) {
	expected := 13
	got := d13_Part1("../data/d13_example.txt")
	if got != expected {
		t.Errorf("Day 13 Part 1, expected: %d, got: %d", expected, got)
	}
}
