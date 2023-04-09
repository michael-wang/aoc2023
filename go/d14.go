package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const d14_Debug = false

func d14() {
	d14_Part1("../data/d14.txt")
	d14_Part2("../data/d14.txt")
}

func d14_Part1(data string) (answer int) {
	cave := d14_Cave{}
	cave.Build(data, false)

	for cave.dropSand() {
		if d14_Debug {
			fmt.Println(cave.String())
			pause("Press ENTER to continue...")
		}
	}
	if d14_Debug {
		fmt.Println(cave.String())
		pause("Press ENTER to continue...")
	}

	answer = cave.count(d14_Sand)
	fmt.Println("[Day14 Part 1] answer: ", answer)
	return
}

func d14_Part2(data string) (answer int) {
	cave := d14_Cave{}
	cave.Build(data, true)

	for cave.dropSand() {
		if d14_Debug {
			fmt.Println(cave.String())
			pause("Press ENTER to continue...")
		}
	}
	if d14_Debug {
		fmt.Println(cave.String())
		pause("Press ENTER to continue...")
	}

	answer = cave.count(d14_Sand)
	fmt.Println("[Day14 Part 2] answer: ", answer)
	return
}

func d14_Load(data string) (pp paths) {
	f, err := os.Open(data)
	if err != nil {
		panic(fmt.Sprintf("Failed to open file: %s", data))
	}
	defer f.Close()

	input := bufio.NewScanner(f)
	for input.Scan() {
		ss := strings.Split(input.Text(), " -> ")
		vv := []vec2{}
		for _, s := range ss {
			v := vec2{}
			v.Parse(s)
			vv = append(vv, v)
		}
		p := path{}
		p.Build(vv)
		pp = append(pp, p)
	}
	return
}

type d14_Cave struct {
	m  mtx
	pp paths

	min    vec2
	max    vec2
	source vec2
}

const (
	d14_Air = iota
	d14_Rock
	d14_Sand
	d14_Source
)

func (cave *d14_Cave) Build(data string, addFloor bool) {
	cave.pp = d14_Load(data)
	cave.min = cave.pp.min()
	cave.max = cave.pp.max()
	if addFloor {
		cave.min.X -= (cave.max.Y - cave.min.Y + 1)
		cave.max.X += (cave.max.Y - cave.min.Y + 1)
		cave.max.Y += 2
	}

	m := makeMatrix(cave.max.X+2, cave.max.Y+2)
	for _, p := range cave.pp {
		for _, v := range p {
			m[v.X][v.Y] = d14_Rock
		}
	}

	cave.source = vec2{500, 0}
	m[cave.source.X][cave.source.Y] = d14_Source

	if addFloor {
		for x := 0; x < len(m); x++ {
			m[x][cave.max.Y] = d14_Rock
		}
	}
	cave.m = m
	if d14_Debug {
		fmt.Println(cave.String())
	}
}

func (cave *d14_Cave) dropSand() (rest bool) {
	sand := vec2{cave.source.X, cave.source.Y}
	cave.m[sand.X][sand.Y] = d14_Sand
	for sand.Y < cave.max.Y {
		if cave.update(&sand) {
			if sand.X == cave.source.X && sand.Y == cave.source.Y {
				break
			}
			return true
		}
	}
	return false
}

func (cave *d14_Cave) update(sand *vec2) (rest bool) {
	if sand.X < cave.min.X || sand.X > cave.max.X {
		return true
	} else if sand.Y > cave.max.Y {
		return true
	}

	m := cave.m
	if m[sand.X][sand.Y+1] == d14_Air {
		m[sand.X][sand.Y] = d14_Air
		sand.Y++
		m[sand.X][sand.Y] = d14_Sand
		return false
	} else if m[sand.X-1][sand.Y+1] == d14_Air {
		m[sand.X][sand.Y] = d14_Air
		sand.X--
		sand.Y++
		m[sand.X][sand.Y] = d14_Sand
		return false
	} else if m[sand.X+1][sand.Y+1] == d14_Air {
		m[sand.X][sand.Y] = d14_Air
		sand.X++
		sand.Y++
		m[sand.X][sand.Y] = d14_Sand
		return false
	}
	return true
}

func (cave d14_Cave) count(obj int) (total int) {
	m := cave.m
	for x := 0; x < len(m); x++ {
		for y := 0; y < len(m[x]); y++ {
			if m[x][y] == obj {
				total++
			}
		}
	}
	return total
}

func (cave d14_Cave) String() string {
	return cave.StringXRange(cave.min.X, len(cave.m))
}
func (cave d14_Cave) StringXRange(x0, x1 int) string {
	min := cave.min
	min.X = x0
	min.lowerBound(0, 0)
	max := cave.max
	max.X = x1
	max.Y += 2
	max.upperBound(len(cave.m), len(cave.m[0]))
	s := ""
	m := cave.m
	for y := 0; y < max.Y; y++ {
		for x := min.X; x < max.X; x++ {
			if x == cave.source.X && y == cave.source.Y {
				if m[x][y] == d14_Sand {
					s += "⊕"
				} else {
					s += "+"
				}
				continue
			}
			switch m[x][y] {
			case d14_Air:
				s += "."
			case d14_Rock:
				s += "#"
			case d14_Sand:
				s += "o"
			default:
				s += "?"
			}
		}
		s += "\n"
	}
	return s
}
