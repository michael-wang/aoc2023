package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	d2_Rock uint8 = iota
	d2_Paper
	d2_Scissor
)

type d2_Round struct {
	op uint8
	my uint8
}

func (g d2_Round) toString() string {
	m := map[uint8]string{
		d2_Rock:    "✊",
		d2_Paper:   "✋",
		d2_Scissor: "✌️ ",
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
		g.op = d2_Rock
	case 'B':
		g.op = d2_Paper
	case 'C':
		g.op = d2_Scissor
	default:
		panic(fmt.Sprintf("Invalid opponent move: %x (valid: 'A', 'B', or 'C')", line[0]))
	}

	switch line[2] {
	case 'X':
		g.my = d2_Rock
	case 'Y':
		g.my = d2_Paper
	case 'Z':
		g.my = d2_Scissor
	default:
		panic(fmt.Sprintf("Invalid my move: %x (valid: 'X', 'Y', or 'Z')", line[2]))
	}
	return g
}

func d02_CountScore(gg []d2_Round) (score int) {
	m := map[uint8]map[uint8]int{
		d2_Rock: {
			d2_Rock:    1 + 3,
			d2_Paper:   2 + 6,
			d2_Scissor: 3 + 0,
		},
		d2_Paper: {
			d2_Rock:    1 + 0,
			d2_Paper:   2 + 3,
			d2_Scissor: 3 + 6,
		},
		d2_Scissor: {
			d2_Rock:    1 + 6,
			d2_Paper:   2 + 0,
			d2_Scissor: 3 + 3,
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
