package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

//lint:ignore U1000 ignore
func d01() {
	d01_Part1("../data/day01.txt")
	d01_Part2("../data/day01.txt")
}

func d01_Part1(data string) (answer int) {
	elves := d01_Parse(data)
	answer, _ = d01_MaxElf(elves)
	fmt.Printf("[Day01 Part1] Max elf calories: %d\n", answer)
	return
}

func d01_Part2(data string) (answer int) {
	elves := d01_Parse(data)
	for i := 0; i < 3; i++ {
		max, i := d01_MaxElf(elves)
		answer += max
		elves = slicePop(elves, i)
	}
	fmt.Println("[Day01 Part2] Sum of top 3 elves: ", answer)
	return
}

func d01_Parse(name string) [][]int {
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

			if !input.Scan() {
				elves = append(elves, elf)
				return elves
			}
		}
		elves = append(elves, elf)
	}
	return elves
}

func d01_MaxElf(elves [][]int) (max, index int) {
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
