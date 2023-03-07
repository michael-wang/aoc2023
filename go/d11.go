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
	d11_Part1("../data/d11.txt")
	d11_Part2("../data/d11.txt")
}

func d11_Part1(data string) (answer int) {
	you := d11_You{}
	you.ParseNote(data)

	for i := 0; i < 20; i++ {
		you.Round(you.d11_DivideBy3)
	}
	you.PrintInspectCounts()

	answer = you.LevelOfMonkeyBusiness()
	fmt.Println("[Day10 Part 1] answer: ", answer)
	return
}

// d11_example.txt
// M0: new % 23
// M1: new % 19
// M2: new % 13
// M3: new % 17
func d11_Part2(data string) (answer int) {
	you := d11_You{}
	you.ParseNote(data)

	const rounds = 10000
	for i := 0; i < rounds; i++ {
		you.Round(you.d11_ModByDivisorProduct)
	}
	you.PrintInspectCounts()

	answer = you.LevelOfMonkeyBusiness()
	fmt.Println("[Day10 Part 2] answer: ", answer)
	return
}

type d11_You struct {
	Monkeys         []*d11_Monkey
	DivisorProducts int
}

func (y *d11_You) d11_DivideBy3(old int) (new int) {
	return old / 3
}

func (y *d11_You) d11_ModByDivisorProduct(old int) (new int) {
	return old % y.DivisorProducts
}

func (y *d11_You) ParseNote(data string) {
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
	y.Monkeys = monkeys

	y.DivisorProducts = 1
	for _, m := range monkeys {
		y.DivisorProducts *= m.TestDivisor
	}
}

func (r *d11_You) PrintItems() {
	for i, m := range r.Monkeys {
		fmt.Printf("Monkey %d: %v\n", i, m.Items)
	}
}

func (r *d11_You) PrintInspectCounts() {
	for i, m := range r.Monkeys {
		fmt.Printf("Monkey %d inspected items %d times.\n", i, m.InspectCount)
	}
}

func (r *d11_You) LevelOfMonkeyBusiness() (monkeyBusiness int) {
	ii := []int{}
	for _, m := range r.Monkeys {
		ii = append(ii, m.InspectCount)
	}
	i1, ii := intSlicePopMax(ii)
	i2, _ := intSlicePopMax(ii)
	return i1 * i2
}

func (r *d11_You) Round(worryLevelDown func(old int) (new int)) {
	// moduloProduct := 1
	// for _, m := range r.Monkeys {
	// 	moduloProduct *= m.TestDivisor
	// }
	for _, m := range r.Monkeys {
		for item, more := m.InspectItem(); more; item, more = m.InspectItem() {
			item = worryLevelDown(item)

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
