package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//lint:ignore U1000 ignore
func d04() {
	d04_Part1and2("../data/d04.txt")
}

func d04_Part1and2(data string) (contains, overlap int) {
	f, err := os.Open(data)
	if err != nil {
		panic(fmt.Sprintf("Failed to open file: %s", data))
	}

	input := bufio.NewScanner(f)
	for input.Scan() {
		pp := d04_ParsePairs(input.Text())

		if d04_FullyContains(pp[0], pp[1]) {
			contains++
			// fmt.Printf("%v,%v\n", *pp[0], *pp[1])
		}

		if d04_Overlap(pp[0], pp[1]) {
			overlap++
		}
	}
	fmt.Println("Number of fully contains: ", contains)
	fmt.Println("Number of overlap: ", overlap)
	return
}

type d04_Pair struct {
	s, t int
}

func (p d04_Pair) length() int {
	return p.t - p.s + 1
}

func d04_ParsePairs(line string) (pp []*d04_Pair) {
	ss := strings.Split(line, ",")
	for _, s := range ss {
		ss := strings.Split(s, "-")
		if len(ss) != 2 {
			panic(fmt.Sprintf("Expect 2 pairs, but got: %v", ss))
		}

		s, err := strconv.Atoi(ss[0])
		if err != nil {
			panic(fmt.Sprintf("Failed to convert: %s to string", ss[0]))
		}

		t, err := strconv.Atoi(ss[1])
		if err != nil {
			panic(fmt.Sprintf("Failed to convert: %s to string", ss[0]))
		}

		pp = append(pp, &d04_Pair{s: s, t: t})
	}
	return
}

func d04_FullyContains(p1, p2 *d04_Pair) bool {
	if p1.length() > p2.length() {
		// Make sure p1 shorter or equal to p2, to simplify comparison.
		p1, p2 = p2, p1
	}

	return p2.s <= p1.s && p1.t <= p2.t
}

func d04_Overlap(p1, p2 *d04_Pair) bool {
	if p2.s < p1.s {
		// Make sure p1.s <= p2s, to simplify comparison.
		p1, p2 = p2, p1
	}

	return p2.s <= p1.t
}
