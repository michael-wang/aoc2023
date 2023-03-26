package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const d12_Debug = false

func d12() {
	d12_Part1("../data/d12.txt")
	// d12_Part2("../data/d12.txt")
}

func d12_Part1(data string) (answer int, minStepsMap *d12_Map) {
	m := &d12_Map{}
	m.Init(data)
	mm := []*d12_Map{}
	pp := m.NextMoves()
	for i := 0; i < len(pp); i++ {
		p := pp[i]
		mm = append(mm, d12_Travel(1, m.DeepCopy(), p)...)
	}

	if d12_Debug {
		pause("Search done, next finding shortest path")
	}

	// Find min steps
	answer = mm[0].Steps
	minStepsMap = mm[0]
	for _, m := range mm {
		if m.Steps <= answer {
			fmt.Printf("Shorter path found: %d\n", m.Steps)
			m.Print()
			answer = m.Steps
			minStepsMap = m
		}
	}
	fmt.Println("[Day10 Part 1] answer: ", answer)
	return
}

func d12_Part2(data string) (answer int) {
	return
}

func d12_Travel(deep int, m *d12_Map, next d12_Position) (mm []*d12_Map) {
	// fmt.Printf("travel next: %s\n", next.ToString())

	for m.Move(next) {
		pp := m.NextMoves()
		// if m.Steps > 110 {
		fmt.Printf("[%d] next moves: %v\n", deep, pp)
		m.Print()
		if d12_Debug {
			pause("Press any key to continue......")
		}
		// }
		if len(pp) == 0 {
			fmt.Println("DEAD END!!!!")
			return
		}
		next = pp[0]
		for i := 1; i < len(pp); i++ {
			mm = append(mm, d12_Travel(deep+1, m.DeepCopy(), pp[i])...)
		}
	}
	mm = append(mm, m)
	return
}

type d12_Position struct {
	X int `json:"x"`
	Y int `json:"y"`
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
	return fmt.Sprintf("(%d, %d)", p.X, p.Y)
}

type d12_Map struct {
	Diagram []string     `json:"diagram"`
	S       d12_Position `json:"s"`
	E       d12_Position `json:"e"`
	// State
	C     d12_Position `json:"c"`
	Steps int          `json:"steps"`
}

func (m *d12_Map) Print() {
	fmt.Printf("Map C: %s, Steps: %d\n", m.C.ToString(), m.Steps)
	// Combine diagram and path to print one map.
	// The input of d12 is large, two maps (diagram and path) are hard to debug.
	for y := 0; y < len(m.Diagram); y++ {
		row := m.Diagram[y]
		if y == m.C.Y {
			char := row[m.C.X]
			row = stringsReplaceByte(row, m.C.X, char-('a'-'A'))
		}
		fmt.Println(row)
	}
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
	}
	m.Diagram = diagram
}

func (m *d12_Map) Traveled(p d12_Position) bool {
	v := m.Diagram[p.Y][p.X]
	return v < 'a' || 'z' < v
}

func (m *d12_Map) Height(p d12_Position) (height int) {
	// Some value which can never reached.
	const invalid = 'z' - 'a' + 1
	if p.Y < 0 || p.Y >= len(m.Diagram) {
		return invalid
	}

	size := len(m.Diagram[p.Y])
	if p.X < 0 || p.X >= size {
		return invalid
	}

	return int(m.Diagram[p.Y][p.X] - 'a')
}

func (m *d12_Map) NextMoves() (pp []d12_Position) {
	if m.C == m.E {
		// Reach end
		return
	}

	cH := m.Height(m.C)

	down := d12_Position{X: m.C.X, Y: m.C.Y + 1}
	diff := m.Height(down) - cH
	if 0 <= diff && diff <= 1 && !m.Traveled(down) {
		if diff == 1 {
			pp = append([]d12_Position{down}, pp...)
		} else {
			pp = append(pp, down)
		}
	}

	right := d12_Position{X: m.C.X + 1, Y: m.C.Y}
	diff = m.Height(right) - cH
	if 0 <= diff && diff <= 1 && !m.Traveled(right) {
		if diff == 1 {
			pp = append([]d12_Position{right}, pp...)
		} else {
			pp = append(pp, right)
		}
	}

	left := d12_Position{X: m.C.X - 1, Y: m.C.Y}
	diff = m.Height(left) - cH
	if 0 <= diff && diff <= 1 && !m.Traveled(left) {
		if diff == 1 {
			pp = append([]d12_Position{left}, pp...)
		} else {
			pp = append(pp, left)
		}
	}

	up := d12_Position{X: m.C.X, Y: m.C.Y - 1}
	diff = m.Height(up) - cH
	if 0 <= diff && diff <= 1 && !m.Traveled(up) {
		if diff == 1 {
			pp = append([]d12_Position{up}, pp...)
		} else {
			pp = append(pp, up)
		}
	}

	return
}

func (m *d12_Map) Move(to d12_Position) (more bool) {
	// Update C
	old := m.C
	m.C = to

	// Update Path
	if to.Y > old.Y {
		m.Set(old.X, old.Y, 'v')
	} else if to.Y < old.Y {
		m.Set(old.X, old.Y, '^')
	} else if to.X > old.X {
		m.Set(old.X, old.Y, '>')
	} else {
		m.Set(old.X, old.Y, '<')
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
		Steps:   m.Steps,
	}
}

func (m *d12_Map) Set(x, y int, char byte) {
	m.Diagram[y] = stringsReplaceByte(m.Diagram[y], x, char)
}
