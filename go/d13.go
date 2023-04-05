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
		left := d13_Packet{}
		row := input.Text()
		left.Parse(row, 0, len(row))

		right := d13_Packet{}
		input.Scan()
		row = input.Text()
		right.Parse(row, 0, len(row))

		if left.Less(right) == 1 {
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

type d13_Packet struct {
	NestedList
}

func (p d13_Packet) Int(i int) *int {
	return p.NestedList.Int(i)
}

func (p d13_Packet) List(i int) *d13_Packet {
	l := p.NestedList.List(i)
	if l == nil {
		return nil
	}
	return &d13_Packet{
		NestedList: *l,
	}
}

func (p *d13_Packet) Parse(row string, start, end int) (next int) {
	return p.NestedList.Parse(row, '[', ']', start, end)
}

func (p *d13_Packet) Append(v interface{}) {
	p.NestedList = append(p.NestedList, v)
}

func (p d13_Packet) Len() int {
	return len(p.NestedList)
}

func (p *d13_Packet) Swap(i, j int) {
	p.NestedList[i], p.NestedList[j] = p.NestedList[j], p.NestedList[i]
}

// Returns:
// 1 if left < right,
// -1 if left > right,
// 0 if left == right.
func (left d13_Packet) Less(right d13_Packet) int {
	for i := 0; i < left.Len(); i++ {
		if right.Len() <= i {
			// Right run out of elements, left is NOT less than right.
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
			less := (*lList).Less(*rList)
			if less != 0 {
				return less
			}
		} else {
			less := 0
			if lInt != nil && rList != nil {
				tmp := d13_Packet{}
				tmp.Append(*lInt)
				less = tmp.Less(*rList)
			} else if lList != nil && rInt != nil {
				tmp := d13_Packet{}
				tmp.Append(*rInt)
				less = (*lList).Less(tmp)
			} else {
				panic("Expect one of 'lInt', 'rInt' is not nil but they are both nil")
			}
			if less != 0 {
				return less
			}
		}
	}
	if right.Len() > left.Len() {
		// Left has no element while right still has element, left is less than right.
		return 1
	}
	// The case of left shorter than right is returned during above for loop,
	// so the only case here is equal sized which mean order undetermined.
	return 0
}
