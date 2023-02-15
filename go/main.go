package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	d01()
	d01_part2()
}

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

func pop(s [][]int, i int) [][]int {
	return append(s[:i], s[i+1:]...)
}
