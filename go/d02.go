package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	rock uint8 = iota
	paper
	scissor
)

type d2_Round struct {
	op uint8
	my uint8
}

func (g d2_Round) toString() string {
	m := map[uint8]string{
		rock:    "✊",
		paper:   "✋",
		scissor: "✌️ ",
	}
	return fmt.Sprintf("%s    %s", m[g.op], m[g.my])
}

func d02_Part1() {
	games := d02_Parse("../data/day02.txt")
	score := d02_CountScore(games)
	fmt.Println("Games score: ", score)
}

func d02_Parse(name string) (games []d2_Round) {
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

func d02_ParseMove(line string) d2_Round {
	if len(line) != 3 {
		panic(fmt.Sprintf("Expect 3 char line but got: %s", line))
	}
	g := d2_Round{}
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

func d02_CountScore(gg []d2_Round) (score int) {
	m := map[uint8]map[uint8]int{
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

	fmt.Println("對手  我  分數")
	for i := 0; i < len(gg); i++ {
		game := gg[i]
		s := m[game.op][game.my]
		if i < 7 {
			fmt.Printf("%s    %d\n", game.toString(), s)
		}
		score += s
	}
	return
}
