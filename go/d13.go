package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func d13() {
	d13_Part1("../data/d13.txt")
	// d13_Part2("../data/d13.txt")
}

func d13_Part1(data string) (answer int) {
	for i, pair := range d13_Load(data) {
		order := pair.Left().CorrectOrder(pair.Right(), 0)
		if order == 1 {
			answer += (i + 1)
		}
		// fmt.Printf("order: %d, answer:%d\n\n", order, answer)
	}
	fmt.Println("[Day13 Part 1] answer: ", answer)
	return
}

func d13_Part2(data string) (answer int) {
	fmt.Println("[Day13 Part 2] answer: ", answer)
	return
}

type list struct {
	// Element of Value can be either int or another list
	Data []interface{}
}

type d13_Pair vec2

func (p *d13_Pair) Left() *list {
	return p.X.(*list)
}

func (p *d13_Pair) Right() *list {
	return p.Y.(*list)
}

func (p *d13_Pair) String() string {
	return fmt.Sprintf("%s\n%s", p.Left().String(), p.Right().String())
}

// Return index of correct ordered pair.
func d13_Load(data string) (pairs []d13_Pair) {
	f, err := os.Open(data)
	if err != nil {
		panic(fmt.Sprintf("Failed to open file: %s", data))
	}
	defer f.Close()

	input := bufio.NewScanner(f)
	for input.Scan() {
		pair := d13_Pair{
			X: &list{},
			Y: &list{},
		}
		pair.Left().Parse(input.Text())

		input.Scan()
		pair.Right().Parse(input.Text())

		pairs = append(pairs, pair)
		// Eat blank line
		input.Scan()
	}
	return
}

func (l *list) Parse(row string) {
	l.Clear()
	row = row[1 : len(row)-1]

	for len(row) > 0 {
		if '0' <= row[0] && row[0] <= '9' {
			end := strings.IndexByte(row, ',')
			if end < 0 {
				end = len(row)
			}
			element, err := strconv.Atoi(row[:end])
			if err != nil {
				panic(err)
			}
			l.Data = append(l.Data, element)

			if end < len(row) {
				row = row[end+1:]
			} else {
				row = row[end:]
			}
		} else if row[0] == '[' {
			end := -1
			brackets := 0
			for i, b := range row {
				if b == '[' {
					brackets++
				} else if b == ']' {
					brackets--
					if brackets == 0 {
						end = i + 1
						break
					}
				}
			}
			if end < 0 {
				panic(fmt.Sprintf("Invalid row: %s", row))
			}
			element := &list{}
			element.Parse(row[:end])
			l.Data = append(l.Data, element)
			if end < len(row) {
				row = row[end+1:]
			} else {
				row = row[end:]
			}
		}
	}
}

func (l *list) String() string {
	s, sep := "", ""
	for _, v := range l.Data {
		s += sep
		sep = ","
		if vInt, ok := v.(int); ok {
			s += strconv.Itoa(vInt)
		} else if vList, ok := v.(*list); ok {
			s += vList.String()
		} else {
			panic(fmt.Sprintf("Unknown list element: %v", v))
		}
	}
	return fmt.Sprintf("[%s]", s)
}

func (l *list) Clear() {
	l.Data = nil
}

// Return 1 if in correct order, -1 if not, 0 if unable to determine.
func (left *list) CorrectOrder(right *list, depth int) int {
	prefix := ""
	for i := 0; i < depth; i++ {
		prefix += "  "
	}
	// fmt.Printf("%sCompare %s vs %s\n", prefix, left.String(), right.String())
	prefix += "  "

	for i, ldata := range left.Data {
		if len(right.Data) <= i {
			return -1
		}

		lVal, lInt := ldata.(int)
		rdata := right.Data[i]
		rVal, rInt := rdata.(int)

		if lInt && rInt {
			// Case 1
			// fmt.Printf("%sCompare %d vs %d\n", prefix, lVal, rVal)
			if lVal < rVal {
				return 1
			} else if lVal > rVal {
				return -1
			}
		} else if !lInt && !rInt {
			// Case 2
			order := ldata.(*list).CorrectOrder(rdata.(*list), depth+1)
			if order == 0 {
				continue
			}
			return order
		} else {
			// Case 3
			order := 0
			if lInt {
				tmp := &list{
					Data: []interface{}{lVal},
				}
				order = tmp.CorrectOrder(rdata.(*list), depth+1)
			} else {
				tmp := &list{
					Data: []interface{}{rVal},
				}
				order = ldata.(*list).CorrectOrder(tmp, depth+1)
			}
			if order == 0 {
				continue
			}
			return order
		}
	}
	if len(right.Data) > len(left.Data) {
		// Left has no element while right still has element, correct order
		return 1
	}
	// The case of left shorter than right is returned during above for loop,
	// so the only case here is equal sized which mean order undetermined.
	return 0
}
