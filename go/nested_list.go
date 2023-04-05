package main

import (
	"fmt"
	"strconv"
)

type NestedList []interface{}

func (l *NestedList) Parse(row string, left, right byte, start, end int) (next int) {
	end = l.findPair(row, left, right, start, end)
	next = end + 1
	i := start + 1
	notNumber := func(b byte) bool {
		return b < '0' || '9' < b
	}
	for i < end {
		if '0' <= row[i] && row[i] <= '9' {
			j := l.find(row, notNumber, i+1, end)
			s := row[i:j]
			i = j + 1

			v, err := strconv.Atoi(s)
			if err != nil {
				panic(fmt.Sprintf("Failed to parse number: %s", s))
			}
			*l = append(*l, v)
		} else if row[i] == '[' {
			v := NestedList{}
			i = v.Parse(row, left, right, i, end)

			*l = append(*l, v)
		} else {
			i++
		}
	}
	return
}

// Given "...[...]...", find index of corresponding right bracket.
// right = left + 4 in this case.
func (l NestedList) findPair(row string, left, right byte, start, end int) (index int) {
	pairs := 1
	for i := start + 1; i < end; i++ {
		if row[i] == left {
			pairs++
		} else if row[i] == right {
			pairs--
			if pairs == 0 {
				return i
			}
		}
	}
	panic(fmt.Sprintf("Unable to find pair of (%v and %v) in %s", left, right, row[start:end]))
}

// Notice if row = [123,456], target = ',', start = 5, end = 8
// We won't find ',', in this case we should return end so caller can
// treat index 5 - 8 as valid segment.
func (l NestedList) find(row string, matcher func(byte) bool, start, end int) (index int) {
	for i := start; i < end; i++ {
		if matcher(row[i]) {
			return i
		}
	}
	return end
}

func (l *NestedList) Clear() {
	*l = (*l)[:0]
}

func (l NestedList) Int(i int) *int {
	v, ok := l[i].(int)
	if !ok {
		return nil
	}
	return &v
}

func (l NestedList) List(i int) *NestedList {
	v, ok := l[i].(NestedList)
	if !ok {
		return nil
	}
	return &v
}

func (l NestedList) String() string {
	s, sep := "[", ""
	for i := 0; i < len(l); i++ {
		pInt := l.Int(i)
		pList := l.List(i)

		if pInt != nil {
			s += fmt.Sprintf("%s%d", sep, *pInt)
		} else if pList != nil {
			s += fmt.Sprintf("%s%s", sep, pList.String())
		} else {
			s += fmt.Sprintf("%s%v", sep, l[i])
		}
		sep = ","
	}
	return s + "]"
}
