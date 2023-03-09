package main

import (
	"bufio"
	"encoding/json"
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
	mm := []*d12_Map{}
	pp := m.NextMoves()
	for _, p := range pp {
		mm = append(mm, d12_Travel(m.DeepCopy(), p)...)
	}

	// Find min steps
	min := mm[0].Steps
	for _, m := range mm {
		fmt.Println(m.Path.ToString())
		if m.Steps < min {
			min = m.Steps
		}
	}
	answer = min
	fmt.Println("[Day10 Part 1] answer: ", answer)
	return
}

func d12_Part2(data string) (answer int) {
	return
}

func d12_Travel(m *d12_Map, next d12_Position) (mm []*d12_Map) {
	// fmt.Printf("travel next: %s\n", next.ToString())

	for m.Move(next) {
		pp := m.NextMoves()
		if len(pp) == 0 {
			return
		}
		next = pp[0]
		for i := 1; i < len(pp); i++ {
			n := m.DeepCopy()
			mm = append(mm, d12_Travel(n, pp[i])...)
		}
	}
	mm = append(mm, m)
	return
}

type d12_Position struct {
	X int `json="x"`
	Y int `json="y"`
}

func (p d12_Position) Copy() d12_Position {
	return d12_Position{
		X: p.X,
		Y: p.Y,
	}
}

func (p d12_Position) Equals(q d12_Position) bool {
	return p.X == q.X && p.Y == q.Y
}

func (p d12_Position) ToString() string {
	return fmt.Sprintf("(X: %d, Y: %d)", p.X, p.Y)
}

type d12_Map struct {
	Diagram []string     `json="diagram"`
	S       d12_Position `json="s"`
	E       d12_Position `json="e"`
	// State
	C     d12_Position `json="c"`
	Path  *Path        `json="path"`
	Steps int          `json="steps"`
}

func (m *d12_Map) ToString() string {
	bb, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(bb)
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
	path := &Path{}
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
		path.InitRow(len(row))
	}
	m.Diagram = diagram
	m.Path = path
}

func (m *d12_Map) Height(p d12_Position) (height int) {
	row := m.Diagram[p.Y]
	return int(row[p.X] - 'a')
}

func (m *d12_Map) NextMoves() (pp []d12_Position) {
	if m.C == m.E {
		// Reach end
		return
	}

	p := m.C
	// TODO: don't go back (by checking Path).
	// Left
	if p.X > 0 {
		q := d12_Position{X: p.X - 1, Y: p.Y}
		if !m.Path.Traveled(q.X, q.Y) {
			diff := m.Height(q) - m.Height(p)
			if diff == 0 || diff == 1 {
				pp = append(pp, q)
			}
		}
	}
	// Up
	if p.Y > 0 {
		q := d12_Position{X: p.X, Y: p.Y - 1}
		if !m.Path.Traveled(q.X, q.Y) {
			diff := m.Height(q) - m.Height(p)
			if diff == 0 || diff == 1 {
				pp = append(pp, q)
			}
		}
	}
	// Right
	if p.X < (len(m.Diagram[p.Y]) - 1) {
		q := d12_Position{X: p.X + 1, Y: p.Y}
		if !m.Path.Traveled(q.X, q.Y) {
			diff := m.Height(q) - m.Height(p)
			if diff == 0 || diff == 1 {
				pp = append(pp, q)
			}
		}
	}
	// Down
	if p.Y < (len(m.Diagram) - 1) {
		q := d12_Position{X: p.X, Y: p.Y + 1}
		if !m.Path.Traveled(q.X, q.Y) {
			diff := m.Height(q) - m.Height(p)
			if diff == 0 || diff == 1 {
				pp = append(pp, q)
			}
		}
	}
	return
}

func (m *d12_Map) Move(to d12_Position) (more bool) {
	// Update C
	old := m.C
	m.C = to

	// Update Path
	if old.Y != to.Y {
		if to.Y > old.Y {
			m.Path.Set(old.X, old.Y, "v")
		} else {
			m.Path.Set(old.X, old.Y, "^")
		}
	} else {
		if to.X > old.X {
			m.Path.Set(old.X, old.Y, ">")
		} else {
			m.Path.Set(old.X, old.Y, "<")
		}
	}

	m.Steps++

	more = !to.Equals(m.E)
	return
}

func (m *d12_Map) DeepCopy() *d12_Map {
	return &d12_Map{
		Diagram: copySliceOfString(m.Diagram),
		S:       m.S.Copy(),
		E:       m.E.Copy(),
		C:       m.C.Copy(),
		Path:    m.Path.Copy(),
		Steps:   m.Steps,
	}
}

type Path struct {
	Diagram []string `json="diagram"`
}

func (p *Path) InitRow(size int) {
	p.Diagram = append(p.Diagram, strings.Repeat(".", size))
}

func (p *Path) Set(x, y int, char string) {
	row := p.Diagram[y]
	row = row[:x] + char + row[x+1:]
	p.Diagram[y] = row
}

func (p *Path) Copy() *Path {
	return &Path{
		Diagram: copySliceOfString(p.Diagram),
	}
}

func (p *Path) Traveled(x, y int) bool {
	return p.Diagram[y][x] != '.'
}

func (p *Path) ToString() string {
	s := fmt.Sprintf("Path{\n")
	for _, row := range p.Diagram {
		s = fmt.Sprintf("%s%s\n", s, row)
	}
	s = fmt.Sprintf("%s\n}\n", s)
	return s
}
