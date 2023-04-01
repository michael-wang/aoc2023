package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//lint:ignore U1000 ignore
func d09() {
	d09_Part1("../data/d09.txt")
	d09_Part2("../data/d09.txt")
}

type d09_Pos struct {
	x, y int
}

func (p *d09_Pos) toString() string {
	return fmt.Sprintf("%d,%d", p.x, p.y)
}

func d09_Part1(data string) (answer int) {
	answer = d09_Main(data, 2)
	fmt.Println("[Day09 Part 1] answer: ", answer)
	return
}

func d09_Part2(data string) (answer int) {
	answer = d09_Main(data, 10)
	fmt.Println("[Day09 Part 2] answer: ", answer)
	return
}

func d09_Main(data string, nKnots int) (answer int) {
	// Prepare data
	pp := make([]*d09_Pos, nKnots)
	for i := 0; i < nKnots; i++ {
		pp[i] = &d09_Pos{x: 0, y: 0}
	}

	visited := map[string]bool{
		pp[nKnots-1].toString(): true,
	}

	// Open file
	f, err := os.Open(data)
	if err != nil {
		panic(fmt.Sprintf("Failed to open file: %s", data))
	}
	defer f.Close()

	// Scan line by line
	input := bufio.NewScanner(f)
	for input.Scan() {
		d09_Move(pp, visited, input.Text())
	}

	answer = len(visited)
	return
}

func d09_Move(pp []*d09_Pos, visited map[string]bool, line string) {
	// fmt.Println(line)
	ss := strings.Split(line, " ")
	if len(ss) != 2 {
		panic(fmt.Sprintf("Invalid moving command: %s", line))
	}
	dir := ss[0]
	length, err := strconv.Atoi(ss[1])
	if err != nil {
		panic(fmt.Sprintf("Invalid length: %s", ss[1]))
	}

	n := len(pp)
	for i := 0; i < length; i++ {
		switch dir {
		case "L":
			pp[0].x--
			for j := 1; j < n; j++ {
				pp[j].Follow(pp[j-1])
			}
		case "R":
			pp[0].x++
			for j := 1; j < n; j++ {
				pp[j].Follow(pp[j-1])
			}
		case "U":
			pp[0].y++
			for j := 1; j < n; j++ {
				pp[j].Follow(pp[j-1])
			}
		case "D":
			pp[0].y--
			for j := 1; j < n; j++ {
				pp[j].Follow(pp[j-1])
			}
		default:
			panic(fmt.Sprintf("Unknown moving direction: %s", ss[0]))
		}
		visited[pp[len(pp)-1].toString()] = true
		/*
			fmt.Println("========", dir, "========")
			for i, p := range pp {
				fmt.Printf("[%d] %s\n", i, p.toString())
			}
		*/
	}
}

func (q *d09_Pos) Follow(p *d09_Pos) {
	dx := p.x - q.x
	dy := p.y - q.y
	// fmt.Println("dx: ", dx, ", dy: ", dy)
	if dx == 0 {
		// Moving along Y-axis
		if dy > 1 {
			q.y++
		} else if dy < -1 {
			q.y--
		}
	} else if dy == 0 {
		// Moving along X-Axis
		if dx > 1 {
			q.x++
		} else if dx < -1 {
			q.x--
		}
	} else {
		// Moving diagonally
		if dx > 1 {
			q.x++
			if dy > 0 {
				q.y++
			} else {
				q.y--
			}
		} else if dx < -1 {
			q.x--
			if dy > 0 {
				q.y++
			} else {
				q.y--
			}
		} else if dy > 1 {
			q.y++
			if dx > 0 {
				q.x++
			} else {
				q.x--
			}
		} else if dy < -1 {
			q.y--
			if dx > 0 {
				q.x++
			} else {
				q.x--
			}
		}
	}
}
