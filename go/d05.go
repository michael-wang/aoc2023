package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/golang-collections/collections/stack"
)

func d05() {
	d05_Part1("../data/d05.txt")
	d05_Part2("../data/d05.txt")
}

func d05_Part1(data string) (answer string) {
	width, height := d05_CountStacks(data)
	answer = d05_ParseStack(data, width, height, false)
	fmt.Println("[Day05 Par1] ", answer)
	return
}

func d05_Part2(data string) (answer string) {
	width, height := d05_CountStacks(data)
	answer = d05_ParseStack(data, width, height, true)
	fmt.Println("[Day05 Par2] ", answer)
	return
}

func d05_CountStacks(data string) (width, height int) {
	f, err := os.Open(data)
	if err != nil {
		panic(fmt.Sprintf("Failed to open file: %s", data))
	}
	defer f.Close()

	input := bufio.NewScanner(f)
	for input.Scan() {
		s := strings.TrimSpace(input.Text())
		if strings.HasPrefix(s, "1") {
			last := string(s[len(s)-1])
			width, err = strconv.Atoi(last)
			if err != nil {
				panic(fmt.Sprintf("Failed to convert %s to number", last))
			}
			return
		}
		height++
	}
	panic("Expect row: 1 2 3... but found none")
}

// width: number of crate stacks
// height: number of crates from heighest stack, i.e. lines of text to parse
// before moving instructions.
func d05_ParseStack(data string, width, height int, moveCratesWithSameOrder bool) (topCrates string) {
	f, err := os.Open(data)
	if err != nil {
		panic(fmt.Sprintf("Failed to open file: %s", data))
	}
	defer f.Close()

	// For simplicity, we parse stack from top to bottom, which creates a
	// problem: we can't build stack unless we reverse when done parsing.
	// Fix this by parsign into string, and then build stack with reverse
	// string order.
	ss := make([]*string, width)
	for i := 0; i < width; i++ {
		s := ""
		ss[i] = &s
	}

	input := bufio.NewScanner(f)
	for h := 0; h < height; h++ {
		if input.Scan() == false {
			panic("Premature termination of input data")
		}
		row := input.Text()
		d05_ParseStackRow(row, ss)
	}

	stacks := make([]*stack.Stack, width)
	for i := 0; i < width; i++ {
		s := ss[i]
		stack := stack.New()
		// fmt.Printf("[%d] ", i)
		for j := 0; j < len(*s); j++ {
			str := *s
			// crate := string(str[j])
			// fmt.Printf("%s ", crate)
			stack.Push(string(str[j]))
		}
		// fmt.Println()
		stacks[i] = stack
	}

	// Skip 2 lines to move instructions.
	input.Scan()
	input.Scan()

	d05_ParseMoves(input, stacks, moveCratesWithSameOrder)

	for _, s := range stacks {
		if s.Len() > 0 {
			topCrates += s.Pop().(string)
		}
	}
	return
}

func d05_ParseStackRow(row string, ss []*string) {
	for i := 1; i < len(row); i += 4 {
		crate := string(row[i])
		if crate != " " {
			is := (i - 1) / 4
			s := ss[is]
			*s = crate + *s
		}
	}
}

func d05_ParseMoves(input *bufio.Scanner, stacks []*stack.Stack, sameOrder bool) {
	var count, from, to int
	var err error
	for input.Scan() {
		ss := strings.Split(input.Text(), " from ")
		s := ss[0][5:]
		count, err = strconv.Atoi(s)
		if err != nil {
			panic(fmt.Sprintf("Failed to parse 'count' from: %s, line: %s", s, input.Text()))
		}

		ss = strings.Split(ss[1], " to ")
		s = ss[0]
		from, err = strconv.Atoi(s)
		if err != nil {
			panic(fmt.Sprintf("Failed to parse 'from' from: %s, line: %s", s, input.Text()))
		}

		s = ss[1]
		to, err = strconv.Atoi(s)
		if err != nil {
			panic(fmt.Sprintf("Failed to parse 'to' from: %s, line: %s", s, input.Text()))
		}

		// fmt.Printf("[%d, %d, %d]\n", count, from, to)

		/*
			for i, s := range stacks {
				fmt.Printf("[%d] ", i+1)
				printStack(*s)
			}
		*/
		d05_MoveCrates(count, from, to, stacks, sameOrder)
	}

}

func d05_MoveCrates(count, from, to int, stacks []*stack.Stack, sameOrder bool) {
	if sameOrder {
		crates := ""
		for i := 0; i < count; i++ {
			c := stacks[from-1].Pop().(string)
			crates = c + crates
		}
		for i := 0; i < count; i++ {
			c := string(crates[i])
			stacks[to-1].Push(c)
		}
	} else {
		for i := 0; i < count; i++ {
			crate := stacks[from-1].Pop()
			stacks[to-1].Push(crate)
		}
	}
}
