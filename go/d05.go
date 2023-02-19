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
	d05_Part1("../data/d05_example.txt")
}

func d05_Part1(data string) (answer string) {
	count, height := d05_CountStacks(data)
	answer = d05_ParseStack(data, count, height)
	fmt.Println("[Day05 Par1] ", answer)
	return
}

func d05_CountStacks(data string) (count, height int) {
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
			count, err = strconv.Atoi(last)
			if err != nil {
				panic(fmt.Sprintf("Failed to convert %s to number", last))
			}
			return
		}
		height++
	}
	panic("Expect row: 1 2 3... but found none")
}

func d05_ParseStack(data string, numOfStacks, maxStackHeight int) (topCrates string) {
	f, err := os.Open(data)
	if err != nil {
		panic(fmt.Sprintf("Failed to open file: %s", data))
	}
	defer f.Close()

	// For simplicity, we parse stack from top to bottom, which creates a
	// problem: we can't build stack unless we reverse when done parsing.
	// Fix this by parsign into string, and then build stack with reverse
	// string order.
	tt := make([]*string, numOfStacks)
	for i := 0; i < numOfStacks; i++ {
		str := ""
		tt[i] = &str
	}

	input := bufio.NewScanner(f)
	for h := 0; h < maxStackHeight; h++ {
		if input.Scan() == false {
			panic("Premature termination of input data")
		}
		row := input.Text()
		d05_ParseStackRow(row, tt)
	}

	stacks := make([]*stack.Stack, numOfStacks)
	for i := 0; i < numOfStacks; i++ {
		t := tt[i]
		s := stack.New()
		fmt.Printf("[%d] ", i)
		for j := 0; j < len(*t); j++ {
			str := *t
			crate := string(str[j])
			fmt.Printf("%s ", crate)
			s.Push(string(str[j]))
		}
		fmt.Println()
		stacks[i] = s
	}

	d05_MoveStacks(input, stacks)

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

func d05_MoveStacks(input *bufio.Scanner, stacks []*stack.Stack) {
	// Skip 2 lines to moves.
	input.Scan()
	input.Scan()

	for input.Scan() {
		ss := strings.Split(input.Text(), " from ")
		s := ss[0][5:]
		count, err := strconv.Atoi(s)
		if err != nil {
			panic(fmt.Sprintf("Failed to parse 'count' from: %s, line: %s", s, input.Text()))
		}

		ss = strings.Split(ss[1], " to ")
		s = ss[0]
		from, err := strconv.Atoi(s)
		if err != nil {
			panic(fmt.Sprintf("Failed to parse 'from' from: %s, line: %s", s, input.Text()))
		}

		s = ss[1]
		to, err := strconv.Atoi(s)
		if err != nil {
			panic(fmt.Sprintf("Failed to parse 'to' from: %s, line: %s", s, input.Text()))
		}

		fmt.Printf("[%d, %d, %d]\n", count, from, to)

		for i := 0; i < count; i++ {
			crate := stacks[from-1].Pop()
			stacks[to-1].Push(crate)
		}
	}

	for _, s := range stacks {
		printStack(*s)
	}
}
