package main

import (
	"bufio"
	"fmt"
	"os"
)

//lint:ignore U1000 ignore
func d08() {
	d08_Part1("../data/d08.txt")
	d08_Part2("../data/d08.txt")
}

func d08_Part1(data string) (answer int) {
	f, err := os.Open(data)
	if err != nil {
		panic(fmt.Sprintf("Failed to open file: %s", data))
	}
	defer f.Close()

	input := bufio.NewScanner(f)
	trees, edge := d08_Parse(input)
	fmt.Println("edge: ", edge)

	interior := d08_CountInteriorVisible(trees)
	fmt.Println("interior: ", interior)

	answer = edge + interior
	fmt.Println("[Day08 Part 1] answer: ", answer)
	return
}

func d08_Part2(data string) (x, y int) {
	f, err := os.Open(data)
	if err != nil {
		panic(fmt.Sprintf("Failed to open file: %s", data))
	}
	defer f.Close()

	input := bufio.NewScanner(f)
	trees, _ := d08_Parse(input)

	x, y, max := d08_MaxScenicScore(trees)
	fmt.Printf("[Day 8 Part 2] x: %d, y: %d, max: %d\n", x, y, max)
	return
}

// trees is Y major, meaning to get tree height at (x, y) with trees[y][x].
func d08_Parse(input *bufio.Scanner) (trees [][]int, edge int) {
	for input.Scan() {
		line := input.Text()
		rowLen := len([]rune(line))
		row := make([]int, rowLen)

		for i, r := range line {
			row[i] = int(r - '0')
		}

		trees = append(trees, row)
	}

	height := len(trees)
	width := len(trees[0])
	edge = width*2 + height*2 - 4
	return
}

func d08_CountInteriorVisible(trees [][]int) (visible int) {
	width := len(trees[0])
	height := len(trees)
	for y := 1; y < (height - 1); y++ {
		for x := 1; x < (width - 1); x++ {
			if d08_Visible(trees, x, y) {
				visible++
			}
		}
	}
	return
}

func d08_Visible(trees [][]int, tx, ty int) (visible bool) {
	height := trees[ty][tx]
	//fmt.Printf("x: %d, y: %d, height: %d\n", tx, ty, height)

	// Look from left to right
	visible = true
	for x := 0; x < tx; x++ {
		//fmt.Printf("\tchecking x: %d, y: %d, h: %d\n", x, ty, trees[ty][x])
		if trees[ty][x] >= height {
			//fmt.Println("\tBlocked!!")
			visible = false
			break
		}
	}
	if visible {
		return
	}

	// Look from top to bottom
	visible = true
	for y := 0; y < ty; y++ {
		//fmt.Printf("\tchecking x: %d, y: %d, h: %d\n", tx, y, trees[y][tx])
		if trees[y][tx] >= height {
			//fmt.Println("\tBlocked!!")
			visible = false
			break
		}
	}
	if visible {
		return
	}

	// Look from right to left
	visible = true
	for x := len(trees[ty]) - 1; x > tx; x-- {
		//fmt.Printf("\tchecking x: %d, y: %d, h: %d\n", x, ty, trees[ty][x])
		if trees[ty][x] >= height {
			//fmt.Println("\tBlocked!!")
			visible = false
			break
		}
	}
	if visible {
		return
	}

	// Look from bottom to up
	visible = true
	for y := len(trees) - 1; y > ty; y-- {
		//fmt.Printf("\tchecking x: %d, y: %d, h: %d\n", tx, y, trees[y][tx])
		if trees[y][tx] >= height {
			//fmt.Println("\tBlocked!!")
			visible = false
			break
		}
	}
	return
}

func d08_MaxScenicScore(trees [][]int) (maxX, maxY, max int) {
	for y := 0; y < len(trees); y++ {
		for x := 0; x < len(trees[y]); x++ {
			score := d08_ScenicScore(trees, x, y)
			if score > max {
				max = score
				maxX = x
				maxY = y
			}
		}
	}
	return
}

func d08_ScenicScore(trees [][]int, tx, ty int) (score int) {
	height := trees[ty][tx]
	left := 0
	for x := tx - 1; x >= 0; x-- {
		left++
		if trees[ty][x] >= height {
			break
		}
	}

	top := 0
	for y := ty - 1; y >= 0; y-- {
		top++
		if trees[y][tx] >= height {
			break
		}
	}

	right := 0
	for x := tx + 1; x < len(trees[ty]); x++ {
		right++
		if trees[ty][x] >= height {
			break
		}
	}

	down := 0
	for y := ty + 1; y < len(trees); y++ {
		down++
		if trees[y][tx] >= height {
			break
		}
	}

	score = left * top * right * down
	return
}
