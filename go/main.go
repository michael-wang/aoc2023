package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// d01()
	// d01_part2()
	d02_Part1()
}

// --------------------------------------------------------
// Day 01
// --------------------------------------------------------
func d01() {
	elves := d01_parse("../data/day01.txt")
	max, _ := max_elf(elves)
	fmt.Printf("[Day01] Max elf calories: %d\n", max)
}

func d01_part2() {
	elves := d01_parse("../data/day01.txt")
	sum := 0
	for i := 0; i < 3; i++ {
		max, i := max_elf(elves)
		sum += max
		elves = pop(elves, i)
	}
	fmt.Println("[Day01 Part2] Sum of top 3 elves: ", sum)
}

func d01_parse(name string) [][]int {
	f, err := os.Open(name)
	if err != nil {
		panic(fmt.Sprintf("Failed to open file: %s", name))
	}
	defer f.Close()

	elves := make([][]int, 0)
	input := bufio.NewScanner(f)
	for input.Scan() {
		// parse elf
		elf := make([]int, 0)
		for len(input.Text()) > 0 {
			calories, err := strconv.Atoi(input.Text())
			if err != nil {
				panic(fmt.Sprintf("Failed to convert string to integer: %s", input.Text()))
			}
			elf = append(elf, calories)

			if input.Scan() == false {
				elves = append(elves, elf)
				return elves
			}
		}
		elves = append(elves, elf)
	}
	return elves
}

func max_elf(elves [][]int) (max, index int) {
	for i, elf := range elves {
		curr := 0
		for _, calories := range elf {
			curr += calories
		}
		if curr > max {
			max = curr
			index = i
		}
	}
	return
}

// --------------------------------------------------------
// Day 02
// --------------------------------------------------------
type move int

const (
	rock move = iota
	paper
	scissor
)

type game struct {
	op move
	my move
}

func (g game) toString() string {
	m := map[move]string{
		rock:    "✊",
		paper:   "✋",
		scissor: "✌️ ",
	}
	return fmt.Sprintf("%s - %s", m[g.op], m[g.my])
}

func d02_Part1() {
	games := d02_Parse("../data/day02.txt")
	score := d02_CountScore(games)
	fmt.Println("Games score: ", score)
}

func d02_Parse(name string) (games []game) {
	f, err := os.Open(name)
	if err != nil {
		panic(fmt.Sprintf("Failed to open file: %s", name))
	}
	defer f.Close()

	input := bufio.NewScanner(f)
	for input.Scan() {
		g := d02_ParseMove(input.Text())
		games = append(games, g)
	}
	return
}

func d02_ParseMove(line string) game {
	if len(line) != 3 {
		panic(fmt.Sprintf("Expect 3 char line but got: %s", line))
	}
	g := game{}
	switch line[0] {
	case 'A':
		g.op = rock
	case 'B':
		g.op = paper
	case 'C':
		g.op = scissor
	default:
		panic(fmt.Sprintf("Invalid opponent move: %x (valid: 'A', 'B', or 'C')", line[0]))
	}

	switch line[2] {
	case 'X':
		g.my = rock
	case 'Y':
		g.my = paper
	case 'Z':
		g.my = scissor
	default:
		panic(fmt.Sprintf("Invalid my move: %x (valid: 'X', 'Y', or 'Z')", line[2]))
	}
	return g
}

func d02_CountScore(gg []game) (score int) {
	m := map[move]map[move]int{
		rock: {
			rock:    1 + 3,
			paper:   2 + 6,
			scissor: 3 + 0,
		},
		paper: {
			rock:    1 + 0,
			paper:   2 + 3,
			scissor: 3 + 6,
		},
		scissor: {
			rock:    1 + 6,
			paper:   2 + 0,
			scissor: 3 + 3,
		},
	}

	for i := 0; i < len(gg); i++ {
		game := gg[i]
		s := m[game.op][game.my]
		if i < 7 {
			fmt.Printf("%s, score: %d\n", game.toString(), s)
		}
		score += s
	}
	return
}

// Utility functions

func pop(s [][]int, i int) [][]int {
	return append(s[:i], s[i+1:]...)
}
