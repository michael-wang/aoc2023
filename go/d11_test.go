package main

import (
	"fmt"
	"sort"
	"testing"
)

func TestD11_Part1(t *testing.T) {
	expected := 10605
	got := d11_Part1("../data/d11_example.txt")
	if got != expected {
		t.Errorf("Day 11 Part 1 failed, expect answer: %d, but got: %d", expected, got)
	}

}

func TestD11_Part2(t *testing.T) {
	expected := map[int]map[int]int{
		1: {
			0: 2,
			1: 4,
			2: 3,
			3: 6,
		},
		20: {
			0: 99,
			1: 97,
			2: 8,
			3: 103,
		},
		1000: {
			0: 5204,
			1: 4792,
			2: 199,
			3: 5192,
		},
		2000: {
			0: 10419,
			1: 9577,
			2: 392,
			3: 10391,
		},
		3000: {
			0: 15638,
			1: 14358,
			2: 587,
			3: 15593,
		},
		4000: {
			0: 20858,
			1: 19138,
			2: 780,
			3: 20797,
		},
		5000: {
			0: 26075,
			1: 23921,
			2: 974,
			3: 26000,
		},
		6000: {
			0: 31294,
			1: 28702,
			2: 1165,
			3: 31204,
		},
		7000: {
			0: 36508,
			1: 33488,
			2: 1360,
			3: 36400,
		},
		8000: {
			0: 41728,
			1: 38268,
			2: 1553,
			3: 41606,
		},
		9000: {
			0: 46945,
			1: 43051,
			2: 1746,
			3: 46807,
		},
		10000: {
			0: 52166,
			1: 47830,
			2: 1938,
			3: 52013,
		},
	}

	keys := []int{}
	for k := range expected {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	for _, round := range keys {
		you := d11_You{}
		you.ParseNote("../data/d11_example.txt")

		for i := 0; i < round; i++ {
			you.Round(you.d11_ModByDivisorProduct)
		}

		fmt.Println("== After round", round, "==")
		for monkey, exp := range expected[round] {
			got := you.Monkeys[monkey].InspectCount
			if got == exp {
				fmt.Printf("Monkey %d inspected item %d times\n", monkey, got)
			} else {
				fmt.Printf("Monkey %d inspected item %d times (expected: %d)\n", monkey, got, exp)
				t.Errorf("Error, expected count: %d, got count: %d\n", exp, got)
			}
		}
	}

	// for round := 0;
}
