package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const d12_Debug = false
const d12_CanGoDown = true

func d12() {
	d12_Part1("../data/d12.txt")
	// d12_Part2("../data/d12.txt")
}

func d12_Part1(data string) (answer int, minStepsMap *d12_Map) {
	m := &d12_Map{}
	m.Init(data)
	mm := []*d12_Map{}
	nn := m.NextMoves()
	for i := 0; i < len(nn); i++ {
		n := nn[i]
		if d12_Debug {
			fmt.Printf("Next: %s\n", n.ToString())
		}
		mm = append(mm, d12_Travel(1, m.DeepCopy(), n)...)
	}

	// if d12_Debug {
	pause("Search done, next finding shortest path")
	// }

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

type d12_Maps []*d12_Map

func (mm d12_Maps) Dup(m *d12_Map) bool {
	for i := 0; i < len(mm); i++ {
		if m.Equals(mm[i]) {
			return true
		}
	}
	return false
}

var d12_Deadends = d12_Maps{}

func d12_Travel(deep int, m *d12_Map, next vec2) (mm []*d12_Map) {
	// if d12_Debug {
	fmt.Printf("travel[%d] curr: %s, next: %s, len(deadend): %d\n", deep, m.C.ToString(), next.ToString(), len(d12_Deadends))
	// }

	// Since we are trying deep first search, we have to save possible moves for later traversal.
	nn := map[*d12_Map]vec2{}

	for m.Move(next) {
		pp := m.NextMoves()
		if len(pp) == 0 {
			if d12_Deadends.Dup(m) {
				m.Print()
				panic("Duplicated deadend")
			}
			d12_Deadends = append(d12_Deadends, m)
			// Break instead of return, so we can explore other moves.
			break
		}

		if d12_Debug {
			m.Print()
			fmt.Printf("travel[%d] next moves: %v\n", deep, pp)
			// pause("Press ENTER to continue......")
		}

		// Pick first candidate as next move
		next = pp[0]
		// Add other candidate for later traversal, if any.
		for i := 1; i < len(pp); i++ {
			nn[m.DeepCopy()] = pp[i]
		}
	}

	if m.C.Equals(m.E) {
		mm = append(mm, m)
	}

	// Check other moves
	for n, pos := range nn {
		mm = append(mm, d12_Travel(deep+1, n, pos)...)
	}
	return
}

type d12_Map struct {
	Diagram sliceS `json:"diagram"`
	S       vec2   `json:"s"`
	E       vec2   `json:"e"`
	// State
	C     vec2 `json:"c"`
	Steps int  `json:"steps"`
}

func (m *d12_Map) Equals(other *d12_Map) bool {
	if m.Steps != other.Steps {
		return false
	}
	if !m.Diagram.Equals(other.Diagram) {
		return false
	}
	if !m.C.Equals(other.C) {
		return false
	}
	if !m.E.Equals(other.E) {
		return false
	}
	if !m.S.Equals(other.S) {
		return false
	}
	return true
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
	diagram := sliceS{}
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

func (m *d12_Map) Traveled(p vec2) bool {
	v := m.Diagram[p.Y][p.X]
	return v < 'a' || 'z' < v
}

func (m *d12_Map) Height(p vec2) (height int) {
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

func (m *d12_Map) NextMoves() (pp []vec2) {
	if m.C == m.E {
		// Reach end
		return
	}

	cH := m.Height(m.C)
	// Map from height diff to position so we can sort by diff later.
	// Climbing up seems a better choice, also faster for debug.
	moves := map[int][]vec2{
		-1: []vec2{}, // for diff < 0
		0:  []vec2{},
		1:  []vec2{},
	}

	down := vec2{X: m.C.X, Y: m.C.Y + 1}
	diff := m.Height(down) - cH
	if diff <= 1 && !m.Traveled(down) {
		if diff == 1 {
			moves[1] = append(moves[1], down)
		} else if diff == 0 {
			moves[0] = append(moves[0], down)
		} else if d12_CanGoDown {
			moves[-1] = append(moves[-1], down)
		}
	}

	right := vec2{X: m.C.X + 1, Y: m.C.Y}
	diff = m.Height(right) - cH
	if diff <= 1 && !m.Traveled(right) {
		if diff == 1 {
			moves[1] = append(moves[1], right)
		} else if diff == 0 {
			moves[0] = append(moves[0], right)
		} else if d12_CanGoDown {
			moves[-1] = append(moves[-1], right)
		}
	}

	left := vec2{X: m.C.X - 1, Y: m.C.Y}
	diff = m.Height(left) - cH
	if diff <= 1 && !m.Traveled(left) {
		if diff == 1 {
			moves[1] = append(moves[1], left)
		} else if diff == 0 {
			moves[0] = append(moves[0], left)
		} else if d12_CanGoDown {
			moves[-1] = append(moves[-1], left)
		}
	}

	up := vec2{X: m.C.X, Y: m.C.Y - 1}
	diff = m.Height(up) - cH
	if diff <= 1 && !m.Traveled(up) {
		if diff == 1 {
			moves[1] = append(moves[1], up)
		} else if diff == 0 {
			moves[0] = append(moves[0], up)
		} else if d12_CanGoDown {
			moves[-1] = append(moves[-1], up)
		}
	}

	pp = append(pp, moves[1]...)
	pp = append(pp, moves[0]...)
	pp = append(pp, moves[-1]...)
	return
}

func (m *d12_Map) Move(to vec2) (more bool) {
	if d12_Debug {
		fmt.Printf("move: %s\n", to.ToString())
	}

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
		Diagram: m.Diagram.DeepCopy(),
		S:       m.S.Copy(),
		E:       m.E.Copy(),
		C:       m.C.Copy(),
		Steps:   m.Steps,
	}
}

func (m *d12_Map) Set(x, y int, char byte) {
	m.Diagram[y] = stringsReplaceByte(m.Diagram[y], x, char)
}
