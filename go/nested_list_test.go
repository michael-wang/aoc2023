package main

import (
	"testing"
)

func Test_Find(t *testing.T) {
	row := "[1,23,456]"
	list := NestedList{}
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

func Test_FindPair(t *testing.T) {
	row := "[9]"
	list := NestedList{}
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

func Test_Parse(t *testing.T) {
	row := "[]"
	list := NestedList{}
	start, end := 0, len(row)
	expLen := 0
	expNext := len(row)
	gotNext := list.Parse(row, '[', ']', start, end)
	gotLen := len(list)
	check := func(row string, list NestedList, expLen, gotLen, expNext, gotNext, start, end int) {
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
