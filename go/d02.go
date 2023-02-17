// TODO After done, I realize this quiz is all about table look up. All I need
// is a 2 level table: from [A/B/C] to [X/Y/Z] to score.
// The 2nd part just need an update to the table, and corresponding change to
// score.
package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	d02_Rock uint8 = iota
	d02_Paper
	d02_Scissor
)

type d02_Round struct {
	op     uint8
	my     uint8
	result uint8
}

func (r d02_Round) toString() string {
	m := map[uint8]string{
		d02_Rock:    "✊",
		d02_Paper:   "✋",
		d02_Scissor: "✌️ ",
	}

	n := map[uint8]string{
		d02_Unknown: "？",
		d02_Lose:    "👎",
		d02_Win:     "👍",
		d02_Draw:    "🤝",
	}
	return fmt.Sprintf("%s    %s    %s", m[r.op], m[r.my], n[r.result])
}

func d02_Part1() {
	rounds := d02_Parse("../data/day02.txt")
	score := d02_CountScore(rounds)
	fmt.Println("[Part 1] Games score: ", score)
}

func d02_Parse(name string, withResult ...bool) (rounds []d02_Round) {
	f, err := os.Open(name)
	if err != nil {
		panic(fmt.Sprintf("Failed to open file: %s", name))
	}
	defer f.Close()

	input := bufio.NewScanner(f)
	for input.Scan() {
		line := input.Text()
		if len(line) != 3 {
			panic(fmt.Sprintf("Expect 3 char line but got: %s", line))
		}

		r := d02_Round{
			op: d02_AsMove(line[0]),
		}
		if len(withResult) > 0 && withResult[0] {
			r.result = d02_AsResult(line[2])
		} else {
			r.my = d02_AsMove(line[2])
		}

		rounds = append(rounds, r)
	}
	return
}

func d02_AsMove(c byte) uint8 {
	switch c {
	case 'A':
		fallthrough
	case 'X':
		return d02_Rock
	case 'B':
		fallthrough
	case 'Y':
		return d02_Paper
	case 'C':
		fallthrough
	case 'Z':
		return d02_Scissor
	default:
		panic(fmt.Sprintf("Invalid opponent move: %x (valid: 'A', 'B', or 'C')", c))
	}
}

func d02_CountScore(rr []d02_Round) (score int) {
	m := map[uint8]map[uint8]int{
		d02_Rock: {
			d02_Rock:    1 + 3,
			d02_Paper:   2 + 6,
			d02_Scissor: 3 + 0,
		},
		d02_Paper: {
			d02_Rock:    1 + 0,
			d02_Paper:   2 + 3,
			d02_Scissor: 3 + 6,
		},
		d02_Scissor: {
			d02_Rock:    1 + 6,
			d02_Paper:   2 + 0,
			d02_Scissor: 3 + 3,
		},
	}

	fmt.Println("對手  我   結果   分數")
	for i, r := range rr {
		s := m[r.op][r.my]
		if i < 7 {
			fmt.Printf("%s     %d\n", r.toString(), s)
		}
		score += s
	}
	return
}

// for part 2, extending d02_Round by adding result field with these values.
const (
	d02_Unknown uint8 = iota
	d02_Lose
	d02_Draw
	d02_Win
)

func d02_AsResult(c byte) uint8 {
	switch c {
	case 'X':
		return d02_Lose
	case 'Y':
		return d02_Draw
	case 'Z':
		return d02_Win
	default:
		panic(fmt.Sprintf("Invalid result: %x (valid: 'X', 'Y', or 'Z')", c))
	}
}

func d02_Part2() {
	rounds := d02_Parse("../data/day02.txt", true)
	d02_CalculateMyMove(rounds)
	score := d02_CountScore(rounds)
	fmt.Println("[Part 2] Games score: ", score)
}

func d02_CalculateMyMove(rr []d02_Round) {
	// map from op to result to my
	m := map[uint8]map[uint8]uint8{
		d02_Rock: {
			d02_Lose: d02_Scissor,
			d02_Win:  d02_Paper,
			d02_Draw: d02_Rock,
		},
		d02_Paper: {
			d02_Lose: d02_Rock,
			d02_Win:  d02_Scissor,
			d02_Draw: d02_Paper,
		},
		d02_Scissor: {
			d02_Lose: d02_Paper,
			d02_Win:  d02_Rock,
			d02_Draw: d02_Scissor,
		},
	}

	for i, r := range rr {
		rr[i].my = m[r.op][r.result]
	}
}
