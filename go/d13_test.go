package main

import (
	"testing"
)

func TestD13_find(t *testing.T) {
	row := "[1,23,456]"
	list := nestedList{}
	end := len(row) - 1
	notNumber := func(b byte) bool {
		return b < '0' || '9' < b
	}
	got := list.find(row, notNumber, 2, end)
	exp := 2
	if got != exp {
		t.Errorf("Find %v in '%s' where start: %d, got: %d but expect: %d", ',', row, 2, got, exp)
	}

	got = list.find(row, notNumber, 4, end)
	exp = 5
	if got != exp {
		t.Errorf("Find %v in '%s' where start: %d, got: %d but expect: %d", ',', row, 4, got, exp)
	}

	got = list.find(row, notNumber, 7, end)
	exp = 9
	if got != exp {
		t.Errorf("Find %v in '%s' where start: %d, got: %d but expect: %d", ',', row, 7, got, exp)
	}
}

func TestD13_FindPair(t *testing.T) {
	row := "[9]"
	list := nestedList{}
	start, end := 0, len(row)
	got := list.findPair(row, '[', ']', start, end)
	exp := 2
	check := func(got, exp, start, end int) {
		if got != exp {
			t.Errorf("Failed to find pair for '%s', start: %d, end: %d, got: %d, expected: %d", row, start, end, got, exp)
		}
	}
	check(got, exp, start, end)

	row = "[11,[22,[33,44]]]"
	start, end = 0, len(row)
	got = list.findPair(row, '[', ']', start, end)
	exp = 16
	check(got, exp, start, end)

	start, end = 4, len(row)
	got = list.findPair(row, '[', ']', start, end)
	exp = 15
	check(got, exp, start, end)

	start, end = 8, len(row)
	got = list.findPair(row, '[', ']', start, end)
	exp = 14
	check(got, exp, start, end)
}

func TestD13_NestedListParse(t *testing.T) {
	row := "[]"
	list := nestedList{}
	start, end := 0, len(row)
	expLen := 0
	expNext := len(row)
	gotNext := list.Parse(row, '[', ']', start, end)
	gotLen := len(list)
	check := func(row string, list nestedList, expLen, gotLen, expNext, gotNext, start, end int) {
		if gotLen != expLen {
			t.Errorf("Failed to parse: '%s', expected: %d, got: %d, start: %d, end: %d", row, expLen, gotLen, start, end)
		}
		if gotNext != expNext {
			t.Errorf("Wrong next value for parsing: '%s', expected: %d, got: %d, start: %d, end: %d", row, expNext, gotNext, start, end)
		}
	}
	check(row, list, expLen, gotLen, expNext, gotNext, start, end)

	row = "[1,22,333]"
	list.Clear()
	start, end = 0, len(row)
	expLen = 3
	expNext = len(row)
	gotNext = list.Parse(row, '[', ']', start, end)
	gotLen = len(list)
	check(row, list, expLen, gotLen, expNext, gotNext, start, end)

	row = "[[1], [2,	3 ,4]]"
	list.Clear()
	start, end = 0, len(row)
	expLen = 2
	expNext = len(row)
	gotNext = list.Parse(row, '[', ']', start, end)
	gotLen = len(list)
	check(row, list, expLen, gotLen, expNext, gotNext, start, end)
}

func TestD13_CorrectOrder(t *testing.T) {
	left := &nestedList{}
	row := "[1,1,3,1,1]"
	left.Parse(row, '[', ']', 0, len(row))
	right := &nestedList{}
	row = "[1,1,5,1,1]"
	right.Parse(row, '[', ']', 0, len(row))
	expected := 1
	got := left.D13_CorrectOrder(*right)
	if expected != got {
		t.Errorf("%v vs %v, expected: %d, got: %d", left, right, expected, got)
	}

	row = "[[1],[2,3,4]]"
	left.Clear()
	left.Parse(row, '[', ']', 0, len(row))

	row = "[[1],4]"
	right.Clear()
	right.Parse(row, '[', ']', 0, len(row))
	expected = 1
	got = left.D13_CorrectOrder(*right)
	if expected != got {
		t.Errorf("%v vs %v, expected: %d, got: %d", left, right, expected, got)
	}

	row = "[9]"
	left.Clear()
	left.Parse(row, '[', ']', 0, len(row))

	row = "[[8,7,6]]"
	right.Clear()
	right.Parse(row, '[', ']', 0, len(row))
	expected = -1
	got = left.D13_CorrectOrder(*right)
	if expected != got {
		t.Errorf("%v vs %v, expected: %d, got: %d", left, right, expected, got)
	}

	row = "[[4,4],4,4]"
	left.Clear()
	left.Parse(row, '[', ']', 0, len(row))

	row = "[[4,4],4,4,4]"
	right.Clear()
	right.Parse(row, '[', ']', 0, len(row))
	expected = 1
	got = left.D13_CorrectOrder(*right)
	if expected != got {
		t.Errorf("%v vs %v, expected: %d, got: %d", left, right, expected, got)
	}

	row = "[7,7,7,7]"
	left.Clear()
	left.Parse(row, '[', ']', 0, len(row))

	row = "[7,7,7]"
	right.Clear()
	right.Parse(row, '[', ']', 0, len(row))
	expected = -1
	got = left.D13_CorrectOrder(*right)
	if expected != got {
		t.Errorf("%v vs %v, expected: %d, got: %d", left, right, expected, got)
	}

	row = "[]"
	left.Clear()
	left.Parse(row, '[', ']', 0, len(row))

	row = "[3]"
	right.Clear()
	right.Parse(row, '[', ']', 0, len(row))
	expected = 1
	got = left.D13_CorrectOrder(*right)
	if expected != got {
		t.Errorf("%v vs %v, expected: %d, got: %d", left, right, expected, got)
	}

	row = "[[[]]]"
	left.Clear()
	left.Parse(row, '[', ']', 0, len(row))

	row = "[[]]"
	right.Clear()
	right.Parse(row, '[', ']', 0, len(row))
	expected = -1
	got = left.D13_CorrectOrder(*right)
	if expected != got {
		t.Errorf("%v vs %v, expected: %d, got: %d", left, right, expected, got)
	}

	row = "[1,[2,[3,[4,[5,6,7]]]],8,9]"
	left.Clear()
	left.Parse(row, '[', ']', 0, len(row))

	row = "[1,[2,[3,[4,[5,6,0]]]],8,9]"
	right.Clear()
	right.Parse(row, '[', ']', 0, len(row))
	expected = -1
	got = left.D13_CorrectOrder(*right)
	if expected != got {
		t.Errorf("%v vs %v, expected: %d, got: %d", left, right, expected, got)
	}

	row = "[7,7,7]"
	left.Clear()
	left.Parse(row, '[', ']', 0, len(row))

	row = "[7,7,7]"
	right.Clear()
	right.Parse(row, '[', ']', 0, len(row))
	expected = 0
	got = left.D13_CorrectOrder(*right)
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
