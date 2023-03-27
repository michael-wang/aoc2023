package main

import (
	"bufio"
	"fmt"
	"os"
	"testing"
)

func TestD12_MapParse(t *testing.T) {
	expDiagram := []string{
		"aabqponm", // Replace 'S' by 'a'
		"abcryxxl",
		"accszzxk", // Replace 'E' by 'z'
		"acctuvwj",
		"abdefghi",
	}
	expS := vec2{X: 0, Y: 0}
	expE := vec2{X: 5, Y: 2}
	expC := vec2{X: 0, Y: 0}

	m := &d12_Map{}
	m.Init("../data/d12_example.txt")
	if m.S != expS {
		t.Errorf("Expected S: %v, got S: %v\n", expS, m.S)
	}
	if m.E != expE {
		t.Errorf("Expected E: %v, got E: %v\n", expE, m.E)
	}
	if m.C != expC {
		t.Errorf("Expected C: %v, got C: %v\n", expC, m.C)
	}
	for i := 0; i < len(expDiagram); i++ {
		expRow := expDiagram[i]
		row := m.Diagram[i]
		if expRow != row {
			t.Errorf("Diagram row: %d, expected: %s, got: %s\n", i, expRow, row)
		}
	}
	// got := d11_Part1("../data/d11_example.txt")
	// if got != expected {
	// 	t.Errorf("Day 11 Part 1 failed, expect answer: %d, but got: %d", expected, got)
	// }
}

func TestD12_ParseAnswer(t *testing.T) {
	expected := []string{
		"v..v<<<<",
		">v.vv<<^",
		".>vv>E^^",
		"..v>>>^^",
		"..>>>>>^",
	}

	got := d12_ParseAnswer("../data/d12_example_answer.txt")

	for i := 0; i < len(expected); i++ {
		expRow := expected[i]
		row := got[i]
		if expRow != row {
			t.Errorf("Answer row: %d, expected: %s, got: %s\n", i, expRow, row)
		}
	}
}

func d12_ParseAnswer(data string) (diagram []string) {
	// Open file
	f, err := os.Open(data)
	if err != nil {
		panic(fmt.Sprintf("Failed to open file: %s", data))
	}
	defer f.Close()

	input := bufio.NewScanner(f)
	for y := 0; input.Scan(); y++ {
		diagram = append(diagram, input.Text())
	}
	return
}

func TestD12_PossibleMoves(t *testing.T) {
	m := &d12_Map{}
	m.Init("../data/d12_example.txt")

	expected := []vec2{
		{X: 0, Y: 1},
		{X: 1, Y: 0},
	}
	got := m.NextMoves()
	if len(got) != len(expected) {
		t.Errorf("[Base case] len(expected): %d != len(got): %d", len(expected), len(got))
		return
	}
	for i := 0; i < len(got); i++ {
		if !expected[i].Equals(got[i]) {
			t.Errorf("[Base case] Expected[%d]: %v != got: %v\n", i, expected[i], got[i])
		}
	}
}

func TestD12_PossibleMoves_Height(t *testing.T) {
	m := &d12_Map{}
	m.Init("../data/d12_example.txt")

	// test for height diff > 1
	m.C = vec2{X: 3, Y: 2}
	expected := []vec2{
		{X: 3, Y: 3},
		{X: 2, Y: 2},
		{X: 3, Y: 1},
	}
	got := m.NextMoves()
	if len(got) != len(expected) {
		t.Errorf("[height] len(expected): %d != len(got): %d", len(expected), len(got))
		return
	}
	for i := 0; i < len(got); i++ {
		if !expected[i].Equals(got[i]) {
			t.Errorf("[height] Expected[%d]: %v != got: %v\n", i, expected[i], got[i])
		}
	}
}

func TestD12_PossibleMoves_DontGoBack(t *testing.T) {
	m := &d12_Map{}
	m.Init("../data/d12_example.txt")

	m.Set(0, 1, '>')
	expected := []vec2{
		{X: 1, Y: 0},
	}
	got := m.NextMoves()
	if len(got) != len(expected) {
		t.Errorf("[back] len(expected): %d != len(got): %d", len(expected), len(got))
		return
	}
	for i := 0; i < len(got); i++ {
		if !expected[i].Equals(got[i]) {
			t.Errorf("[back] Expected[%d]: %v != got: %v\n", i, expected[i], got[i])
		}
	}
}

func TestD12_PossibleMoves_ReachEnd(t *testing.T) {
	m := &d12_Map{}
	m.Init("../data/d12_example.txt")
	m.C = m.E

	expected := []vec2{}
	got := m.NextMoves()
	if len(got) != len(expected) {
		t.Errorf("[end] len(expected): %d != len(got): %d", len(expected), len(got))
	}
}

func TestD12_Part1(t *testing.T) {
	expected := 31
	got, _ := d12_Part1("../data/d12_example.txt")
	if got != expected {
		t.Errorf("Day 12 Part 1 failed, expect answer: %d, but got: %d", expected, got)
	}
}

func TestD12_Deadends(t *testing.T) {
	m := &d12_Map{}
	m.Init("../data/d12.txt")

	mm := d12_Maps{}
	mm = append(mm, m)

	if !mm.Dup(m) {
		t.Error("Expect duplicated but not")
	}

	m2 := m.DeepCopy()
	next := m2.NextMoves()[0]
	m2.Move(next)
	if mm.Dup(m2) {
		t.Error("Expect NOT duplicated but IS")
	}
	mm = append(mm, m2)

	m3 := m2.DeepCopy()
	if !mm.Dup(m3) {
		t.Error("Expect duplicated but not")
	}
}
