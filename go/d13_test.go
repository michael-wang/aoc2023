package main

import (
	"sort"
	"testing"
)

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

func TestD13_Swap(t *testing.T) {
	row := "[1, 2, 3]"
	p1 := &d13_Packet{}
	p1.Parse(row, 0, len(row))

	row = "[[1], [2]]"
	p2 := &d13_Packet{}
	p2.Parse(row, 0, len(row))

	pp := d13_Packets{}
	pp = append(pp, *p1, *p2)
	pp.Swap(0, 1)
	if pp[0].Less(*p2) != 0 {
		t.Errorf("Expect pp[0]: %v but got: %v", *p2, pp[0])
	}
}

func TestD13_PacketsLess(t *testing.T) {
	pp := d13_Packets{}

	row := "[]"
	p1 := &d13_Packet{}
	p1.Parse(row, 0, len(row))
	pp = append(pp, *p1)

	row = "[[]]"
	p2 := &d13_Packet{}
	p2.Parse(row, 0, len(row))
	pp = append(pp, *p2)

	if !pp.Less(0, 1) {
		t.Errorf("Expect %v < %v but NOT", pp[0], pp[1])
	}

	pp.Clear()
	row = "[[[]]]"
	p1 = &d13_Packet{}
	p1.Parse(row, 0, len(row))
	pp = append(pp, *p1)

	row = "[1,1,3,1,1]"
	p2 = &d13_Packet{}
	p2.Parse(row, 0, len(row))
	pp = append(pp, *p2)

	if !pp.Less(0, 1) {
		t.Errorf("Expect %v < %v but NOT", pp[0], pp[1])
	}

	pp.Clear()
	row = "[1,[2,[3,[4,[5,6,0]]]],8,9]"
	p1 = &d13_Packet{}
	p1.Parse(row, 0, len(row))
	pp = append(pp, *p1)

	row = "[1,[2,[3,[4,[5,6,7]]]],8,9]"
	p2 = &d13_Packet{}
	p2.Parse(row, 0, len(row))
	pp = append(pp, *p2)

	if !pp.Less(0, 1) {
		t.Errorf("Expect %v < %v but NOT", pp[0], pp[1])
	}
}

func TestD13_Part2_Example(t *testing.T) {
	expected := `[]
[[]]
[[[]]]
[1,1,3,1,1]
[1,1,5,1,1]
[[1],[2,3,4]]
[1,[2,[3,[4,[5,6,0]]]],8,9]
[1,[2,[3,[4,[5,6,7]]]],8,9]
[[1],4]
[[2]]
[3]
[[4,4],4,4]
[[4,4],4,4,4]
[[6]]
[7,7,7]
[7,7,7,7]
[[8,7,6]]
[9]`

	packets := d13_Packets{}
	packets.Load("../data/d13_example.txt")
	d1 := &d13_Packet{}
	d1.Parse("[[2]]", 0, 5)
	packets = append(packets, *d1)
	d2 := &d13_Packet{}
	d2.Parse("[[6]]", 0, 5)
	packets = append(packets, *d2)

	sort.Sort(packets)
	str := packets.String()
	compareString(expected, str, t)

	got := packets.Find(*d1)
	exp := 10
	if got != exp {
		t.Errorf("Expect finding %s at %d but got %d", d1.String(), exp, got)
	}

	got = packets.Find(*d2)
	exp = 14
	if got != exp {
		t.Errorf("Expect finding %s at %d but got %d", d1.String(), exp, got)
	}
}

func compareString(s1, s2 string, t *testing.T) {
	for i, b := range []byte(s1) {
		if b != s2[i] {
			t.Errorf("index: %d\n%s\n%v\n%s\n%v", i, s1[:i+1], b, s2[:i+1], s2[i])
			return
		}
	}
}

func TestD13_Part2(t *testing.T) {
	expected := 140
	got := d13_Part2("../data/d13_example.txt")
	if got != expected {
		t.Errorf("Day 13 Part 2, expected: %d, got: %d", expected, got)
	}
}
