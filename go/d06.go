package main

import (
	"bufio"
	"fmt"
	"os"
)

//lint:ignore U1000 ignore
func d06() {
	d06_Part1("../data/d06.txt")
	d06_Part2("../data/d06.txt")
}

func d06_Part1(data string) (answer []int) {
	f, err := os.Open(data)
	if err != nil {
		panic(fmt.Sprintf("Failed to open file: %s", data))
	}
	defer f.Close()

	input := bufio.NewScanner(f)
	for input.Scan() {
		// Add 1 because quiz ask for index starting from 1.
		index := d06_FindMarker(input.Text(), 4) + 1
		answer = append(answer, index)
	}

	fmt.Printf("[Day 05 Part 1] answer: %d\n", answer)
	return
}

func d06_Part2(data string) (answer []int) {
	f, err := os.Open(data)
	if err != nil {
		panic(fmt.Sprintf("Failed to open file: %s", data))
	}
	defer f.Close()

	input := bufio.NewScanner(f)
	for input.Scan() {
		// Add 1 because quiz ask for index starting from 1.
		index := d06_FindMarker(input.Text(), 14) + 1
		answer = append(answer, index)
	}

	fmt.Printf("[Day 05 Part 1] answer: %d\n", answer)
	return
}

// n is marker length, n = 4 in part 1.
// Assume len(line) >= n
func d06_FindMarker(line string, n int) (index int) {
	rr := []rune(line)
	for i := (n - 1); i < len(rr); i++ {
		m := map[rune]int{}
		for j := i; j > (i - n); j-- {
			r := rr[j]
			m[r]++
			// d06_PrintMap(m)
			if m[r] > 1 {
				goto next
			}
		}
		return i
	next:
	}
	panic(fmt.Sprintf("Expect has marker but this line does not: %s", line))
}
