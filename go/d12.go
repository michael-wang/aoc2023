package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func d12() {
	d12_Part1("../data/d12.txt")
	// d12_Part2("../data/d12.txt")
}

func d12_Part1(data string) (answer int) {
	m := &d12_Map{}
	m.Init(data)
	fmt.Println("[Day10 Part 1] answer: ", answer)
	return
}

func d12_Part2(data string) (answer int) {
	return
}

// func d12_Trace(m *d12_Map) (mm []*d12_Map) {
// }

type d12_Position struct {
	X, Y int
}

func (p d12_Position) Copy() d12_Position {
	return d12_Position{
		X: p.X,
		Y: p.Y,
	}
}

type d12_Map struct {
	Diagram []string
	S       d12_Position
	E       d12_Position
	// State
	C     d12_Position
	Path  []string
	Steps int
}

func (m *d12_Map) Init(data string) {
	// Open file
	f, err := os.Open(data)
	if err != nil {
		panic(fmt.Sprintf("Failed to open file: %s", data))
	}
	defer f.Close()

	input := bufio.NewScanner(f)
	diagram := []string{}
	path := []string{}
	for y := 0; input.Scan(); y++ {
		row := input.Text()
		if x := strings.Index(row, "S"); x != -1 {
			m.S.X = x
			m.S.Y = y
			m.C.X = x
			m.C.Y = y
		}
		if x := strings.Index(row, "E"); x != -1 {
			m.E.X = x
			m.E.Y = y
		}

		row = strings.Replace(row, "S", "a", 1)
		row = strings.Replace(row, "E", "z", 1)
		diagram = append(diagram, row)
		path = append(path, strings.Repeat(".", len(row)))
	}
	m.Diagram = diagram
	m.Path = path
}

func (m *d12_Map) Height(p d12_Position) (height int) {
	row := m.Diagram[p.Y]
	return int(row[p.X] - 'a')
}

func (m *d12_Map) PossibleMoves(p d12_Position) (pp []d12_Position) {
	// TODO: don't go back (by checking Path).
	// Left
	if p.X > 0 {
		q := d12_Position{X: p.X - 1, Y: p.Y}
		if intAbs(m.Height(p)-m.Height(q)) <= 1 {
			pp = append(pp, q)
		}
	}
	// Up
	if p.Y > 0 {
		q := d12_Position{X: p.X, Y: p.Y - 1}
		if intAbs(m.Height(p)-m.Height(q)) <= 1 {
			pp = append(pp, q)
		}
	}
	// Right
	if p.X < (len(m.Diagram[p.Y]) - 1) {
		q := d12_Position{X: p.X + 1, Y: p.Y}
		if intAbs(m.Height(p)-m.Height(q)) <= 1 {
			pp = append(pp, q)
		}
	}
	// Down
	if p.Y < (len(m.Diagram) - 1) {
		q := d12_Position{X: p.X, Y: p.Y + 1}
		if intAbs(m.Height(p)-m.Height(q)) <= 1 {
			pp = append(pp, q)
		}
	}
	return
}

func (m *d12_Map) DeepCopy() *d12_Map {
	return &d12_Map{
		Diagram: copySliceOfString(m.Diagram),
		S:       m.S.Copy(),
		E:       m.E.Copy(),
		C:       m.C.Copy(),
		Path:    copySliceOfString(m.Path),
		Steps:   m.Steps,
	}
}
