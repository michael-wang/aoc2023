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
	ss := newStacks(count)
	d05_ParseStack(data, ss, height)
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

func d05_ParseStack(data string, ss map[int]*stack.Stack, height int) {
	f, err := os.Open(data)
	if err != nil {
		panic(fmt.Sprintf("Failed to open file: %s", data))
	}
	defer f.Close()

	input := bufio.NewScanner(f)
	for h := 0; h < height; h++ {
		if input.Scan() == false {
			panic("Premature termination of input data")
		}
		row := input.Text()
		d05_ParseStackRow(row, ss)
	}

	for i := 0; i < len(ss); i++ {
		//stringStackReverse(s)
		printStack(*ss[i])
	}
	return
}

func d05_ParseStackRow(row string, ss map[int]*stack.Stack) {
	for i := 1; i < len(row); i += 4 {
		crate := string(row[i])
		if crate != " " {
			is := (i - 1) / 4
			s := ss[is]
			s.Push(crate)
			fmt.Printf("i: %d, crate: %s, #is[%d]: %d\n", i, crate, is, s.Len())
			//printStack(*s)
		}
	}
}
