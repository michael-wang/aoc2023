package main

import (
	"testing"
)

func TestD10_Part1(t *testing.T) {
	expected := 13140
	got := d10_Part1("../data/d10_example.txt")
	if got != expected {
		t.Errorf("Day 10 Part 1 failed, expect answer: %d, but got: %d", expected, got)
	}

}

func TestD10_Part2(t *testing.T) {
	exps := make([]string, 6)
	exps[0] = "##..##..##..##..##..##..##..##..##..##.."
	exps[1] = "###...###...###...###...###...###...###."
	exps[2] = "####....####....####....####....####...."
	exps[3] = "#####.....#####.....#####.....#####....."
	exps[4] = "######......######......######......####"
	exps[5] = "#######.......#######.......#######....."

	device := d10_NewDevice("../data/d10_example.txt")
	device.RunForCRT()
	gots := device.CRT.Rows

	for i := 0; i < len(exps); i++ {
		got := string(gots[i])
		if got != exps[i] {
			t.Errorf("Day 10 Part 2 failed, row[%d] expect: %s, but got: %s", i, exps[i], got)
		}
	}
}
