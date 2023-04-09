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
	// d14_Part2("../data/d14.txt")
}

func d14_Part1(data string) (answer int) {
	cave := d14_Cave{}
	cave.Build(data)

	for cave.dropSand() {
		if d14_Debug {
			fmt.Println(cave.String())
		}
	}

	answer = cave.count(d14_Sand)
	fmt.Println("[Day14 Part 1] answer: ", answer)
	return
}

func d14_Part2(data string) (answer int) {
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

func (cave *d14_Cave) Build(data string) {
	cave.pp = d14_Load(data)
	cave.min = cave.pp.min()
	cave.max = cave.pp.max()
	cave.max.X++
	cave.max.Y++

	m := makeMatrix(cave.max.X+1, cave.max.Y+1)
	for _, p := range cave.pp {
		for _, v := range p {
			m[v.X][v.Y] = d14_Rock
		}
	}

	cave.source = vec2{500, 0}
	m[cave.source.X][cave.source.Y] = d14_Source
	cave.m = m
	if d14_Debug {
		fmt.Println(cave.String())
	}
}

func (cave *d14_Cave) dropSand() (rest bool) {
	sand := vec2{cave.source.X, cave.source.Y + 1}
	cave.m[sand.X][sand.Y] = d14_Sand
	for sand.Y < cave.max.Y {
		if cave.update(&sand) {
			return true
		}
	}
	return false
}

func (cave *d14_Cave) update(sand *vec2) (rest bool) {
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
	for x := 0; x < cave.max.X; x++ {
		for y := 0; y < cave.max.Y; y++ {
			if m[x][y] == obj {
				total++
			}
		}
	}
	return total
}

func (cave d14_Cave) String() string {
	min := cave.min
	min.X--
	min.Y--
	min.lowerBound(0, 0)
	max := cave.max
	max.X++
	max.Y++
	max.upperBound(len(cave.m), len(cave.m[0]))
	s := ""
	for y := 0; y < max.Y; y++ {
		for x := min.X; x < max.X; x++ {
			switch cave.m[x][y] {
			case d14_Air:
				s += "."
			case d14_Rock:
				s += "#"
			case d14_Sand:
				s += "o"
			case d14_Source:
				s += "+"
			default:
				s += "?"
			}
		}
		s += "\n"
	}
	return s
}
