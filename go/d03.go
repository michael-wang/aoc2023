package main

import (
	"bufio"
	"fmt"
	"os"
)

func d03() {
	d03_Part1("../data/day03.txt")
	d03_Part2("../data/day03.txt")
}

func d03_Part1(data string) (answer int) {
	f, err := os.Open(data)
	if err != nil {
		panic(fmt.Sprintf("Failed to open file: %s", data))
	}
	defer f.Close()

	input := bufio.NewScanner(f)
	for input.Scan() {
		answer += d03_SumPriorities(input.Text())
	}
	fmt.Println("[Day03 Part1] Priority sum: ", answer)
	return
}

func d03_SumPriorities(r string) (sum int) {
	found := map[rune]bool{}
	half := len(r) / 2
	for _, char := range r[:half] {
		if found[char] {
			continue
		}
		for _, other := range r[half:] {
			if char == other {
				// char is duplicated
				sum += d03_Priority(char)
				found[char] = true
				break
			}
		}
	}
	return
}

// Rule of priority:
// 'A' through 'Z': 27 throough 52
// 'a' through 'z': 1 through 26
func d03_Priority(char rune) int {
	if char <= 'Z' {
		return int(char - '&')
	} else {
		return int(char - '`')
	}
}

func d03_Part2(data string) (answer int) {
	f, err := os.Open(data)
	if err != nil {
		panic(fmt.Sprintf("Failed to open file: %s", data))
	}
	defer f.Close()

	input := bufio.NewScanner(f)
	var group [3]string
	i := 0
	for input.Scan() {
		group[i] = input.Text()
		i++
		if i < 3 {
			// Keep filling group
			continue
		}

		char := d03_FindCommon(group)
		answer += d03_Priority(char)
		// Reset group index
		i = 0
	}
	fmt.Printf("[Day03 Part2] Sum of common items: %d\n", answer)
	return
}

func d03_FindCommon(group [3]string) (comm rune) {
	// Find shortest string as target.
	// Group are divided into 3 strings identity by indexes i, j, and k.
	i, j, k := shortestString(group[:]), 0, 0
	switch i {
	case 0:
		j, k = 1, 2
	case 1:
		j, k = 0, 2
	case 2:
		j, k = 0, 1
	}

	// m, n, o are items corresponding to indexes i, j, and k.
	for _, m := range group[i] {
		for _, n := range group[j] {
			if m == n {
				// m Found in j group, now proceed to k group:
				for _, o := range group[k] {
					if m == o {
						return m
					}
				}
			}
		}
	}
	panic(fmt.Sprintf("No common item found in group: %s", group))
}
