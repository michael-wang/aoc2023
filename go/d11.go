package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func d11() {
	d11_Part1("../data/d11_example.txt")
	// d11_Part2("../data/d11.txt")
}

func d11_Part1(data string) (answer int) {
	// Open file
	f, err := os.Open(data)
	if err != nil {
		panic(fmt.Sprintf("Failed to open file: %s", data))
	}
	defer f.Close()

	monkeys := []*d11_Monkey{}
	// Scan line by line
	input := bufio.NewScanner(f)
	for input.Scan() {
		// fmt.Println(input.Text())
		m := &d11_Monkey{}
		m.Parse(input)
		// m.Print()
		monkeys = append(monkeys, m)
		// Skip one line
		input.Scan()
	}

	round := Round{Monkeys: monkeys}
	round.Run(20, true)
	round.PrintInspectCounts()
	i1, i2 := round.Top2InspectCounts()
	answer = i1 * i2

	fmt.Println("[Day10 Part 1] answer: ", answer)
	return
}

func d11_Part2(data string) (answer int) {
	return
}

type Round struct {
	Monkeys []*d11_Monkey
}

func (r *Round) PrintItems() {
	for i, m := range r.Monkeys {
		fmt.Printf("Monkey %d: %v\n", i, m.Items)
	}
}

func (r *Round) PrintInspectCounts() {
	for i, m := range r.Monkeys {
		fmt.Printf("Monkey %d inspected items %d times.\n", i, m.InspectCount)
	}
}

func (r *Round) Top2InspectCounts() (i1, i2 int) {
	ii := []int{}
	for _, m := range r.Monkeys {
		ii = append(ii, m.InspectCount)
	}
	i1, ii = intSlicePopMax(ii)
	i2, _ = intSlicePopMax(ii)
	return
}

func (r *Round) Run(round int, worryLevelDown bool) {
	// moduloProduct := 1
	// for _, m := range r.Monkeys {
	// 	moduloProduct *= m.TestDivisor
	// }
	for i := 0; i < round; i++ {
		for _, m := range r.Monkeys {
			for item, more := m.InspectItem(); more; item, more = m.InspectItem() {
				// Worry level down
				if worryLevelDown {
					item = item / 3
					// } else {
					// item = item % moduloProduct
				}

				// Monkey test item
				divisible := item%m.TestDivisor == 0
				if divisible {
					// True
					r.Monkeys[m.TrueMonkey].AcceptItem(item)
				} else {
					// False
					r.Monkeys[m.FalseMonkey].AcceptItem(item)
				}
			}
		}
	}
}

const (
	d11_op_add  = "+"
	d11_op_mul  = "*"
	d11_opr_num = "num"
	d11_opr_old = "old"
)

type d11_Monkey struct {
	Items []int `json:"items"`
	// Operation
	Oprator  string `json:"oprator"`
	Oprand   string `json:"operand"`
	OpNumber int    `json:"op_number"`
	// Test
	TestDivisor int `json:"test_divisor"`
	TrueMonkey  int `json:"true_monkey"`
	FalseMonkey int `json:"false_monkey"`

	InspectCount int `json:"inspect_count"`
}

func (m *d11_Monkey) NextItem() (item int) {
	item, items := intSlicePop(m.Items, 0)
	m.Items = items
	return
}

func (m *d11_Monkey) InspectItem() (item int, more bool) {
	if len(m.Items) == 0 {
		more = false
		return
	}
	more = true
	m.InspectCount++
	item = m.NextItem()

	switch m.Oprator {
	case d11_op_add:
		if m.Oprand == d11_opr_num {
			item += m.OpNumber
		} else {
			item += item
		}
	case d11_op_mul:
		if m.Oprand == d11_opr_num {
			item *= m.OpNumber
		} else {
			item *= item
		}
	default:
		panic(fmt.Sprintf("Invalid operator: %s", m.Oprator))
	}
	return
}

func (m *d11_Monkey) AcceptItem(item int) {
	m.Items = append(m.Items, item)
}

func (m *d11_Monkey) Print() {
	bb, err := json.MarshalIndent(m, "", "    ")
	if err != nil {
		panic(fmt.Sprintf("Failed to print monkey: %v, err: %v", m, err))
	}
	fmt.Println(string(bb))
}

func (m *d11_Monkey) Parse(input *bufio.Scanner) {
	// items
	if !input.Scan() {
		panic("Expecting items but reach EOF")
	}
	ss := strings.Split(input.Text(), ": ")
	if len(ss) != 2 {
		panic(fmt.Sprintf("Invalid items: %s", input.Text()))
	}
	ss = strings.Split(ss[1], ", ")
	for _, s := range ss {
		item, err := strconv.Atoi(s)
		if err != nil {
			panic(fmt.Sprintf("Invalid item: %s", s))
		}
		m.Items = append(m.Items, item)
	}

	// Operation
	if !input.Scan() {
		panic("Expecting items but reach EOF")
	}
	ss = strings.Split(input.Text(), " = ")
	if len(ss) != 2 {
		panic(fmt.Sprintf("Invalid operation: %s", input.Text()))
	}
	ss = strings.Split(ss[1], " ")
	if len(ss) != 3 {
		panic(fmt.Sprintf("Unexpected operation: %s", input.Text()))
	}

	switch ss[1] {
	case "*":
		m.Oprator = d11_op_mul
		if ss[2] == "old" {
			m.Oprand = d11_opr_old
		} else {
			m.Oprand = d11_opr_num
			num, err := strconv.Atoi(ss[2])
			if err != nil {
				panic(fmt.Sprintf("Invalid number: %s", ss[2]))
			}
			m.OpNumber = num
		}
	case "+":
		m.Oprator = d11_op_add
		if ss[2] == "old" {
			m.Oprand = d11_opr_old
		} else {
			m.Oprand = d11_opr_num
			num, err := strconv.Atoi(ss[2])
			if err != nil {
				panic(fmt.Sprintf("Invalid number: %s", ss[2]))
			}
			m.OpNumber = num
		}
	default:
		panic(fmt.Sprintf("Invalid operation: %s", input.Text()))
	}

	// Test
	if !input.Scan() {
		panic("Expecting test but reach EOF")
	}
	ss = strings.Split(input.Text(), " by ")
	if len(ss) != 2 {
		panic(fmt.Sprintf("Invalid test: %s", input.Text()))
	}
	divisor, err := strconv.Atoi(ss[1])
	if err != nil {
		panic(fmt.Sprintf("Invalid test divisor: %s", input.Text()))
	}
	m.TestDivisor = divisor

	// True monkey
	if !input.Scan() {
		panic("Expecting true monkey but reach EOF")
	}
	m.TrueMonkey = int(input.Text()[len(input.Text())-1] - '0')
	// False monkey
	if !input.Scan() {
		panic("Expecting false monkey but reach EOF")
	}
	m.FalseMonkey = int(input.Text()[len(input.Text())-1] - '0')
}
