package main

import (
	"bufio"
	"fmt"
	"os"
)

func d13() {
	d13_Part1("../data/d13.txt")
	// d13_Part2("../data/d13.txt")
}

func d13_Part1(data string) (answer int) {
	f, err := os.Open(data)
	if err != nil {
		panic(fmt.Sprintf("Failed to open file: %s", data))
	}
	defer f.Close()

	input := bufio.NewScanner(f)
	index := 1
	for input.Scan() {
		left := nestedList{}
		row := input.Text()
		left.Parse(row, '[', ']', 0, len(row))

		right := nestedList{}
		input.Scan()
		row = input.Text()
		right.Parse(row, '[', ']', 0, len(row))

		if left.D13_CorrectOrder(right) == 1 {
			answer += index
		}

		// Eat blank line
		input.Scan()
		index++
	}

	fmt.Println("[Day13 Part 1] answer: ", answer)
	return
}

func d13_Part2(data string) (answer int) {
	fmt.Println("[Day13 Part 2] answer: ", answer)
	return
}

// Return 1 if in correct order,
//
//	-1 if not,
//	0 if undetermined
func (left nestedList) D13_CorrectOrder(right nestedList) int {
	for i := 0; i < len(left); i++ {
		if len(right) <= i {
			return -1
		}

		lInt, rInt := left.Int(i), right.Int(i)
		lList, rList := left.List(i), right.List(i)

		if lInt != nil && rInt != nil {
			if *lInt < *rInt {
				return 1
			} else if *lInt > *rInt {
				return -1
			}
		} else if lList != nil && rList != nil {
			correct := (*lList).D13_CorrectOrder(*rList)
			if correct != 0 {
				return correct
			}
		} else {
			correct := 0
			if lInt != nil && rList != nil {
				tmp := nestedList{}
				tmp = append(tmp, *lInt)
				correct = tmp.D13_CorrectOrder(*rList)
			} else if lList != nil && rInt != nil {
				tmp := nestedList{}
				tmp = append(tmp, *rInt)
				correct = (*lList).D13_CorrectOrder(tmp)
			} else {
				panic("Expect one of 'lInt', 'rInt' is not nil but they are both nil")
			}
			if correct != 0 {
				return correct
			}
		}
	}
	if len(right) > len(left) {
		// Left has no element while right still has element, correct order
		return 1
	}
	// The case of left shorter than right is returned during above for loop,
	// so the only case here is equal sized which mean order undetermined.
	return 0
}
