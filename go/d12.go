package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"math"
	"os"
)

func d12() {
	d12_Part1("../data/d12.txt")
	// d12_Part2("../data/d12.txt")
}

func d12_Part1(data string) (answer int) {
	start, end, m := d12_Load(data)
	answer = d12_Dijkstra(start, end, m)
	fmt.Println("[Day10 Part 1] answer: ", answer)
	return
}

func d12_Part2(data string) (answer int) {
	return
}

type point struct {
	X      int
	Y      int
	Height int
	Dist   int
}

func (p point) ToString() string {
	return fmt.Sprintf("(%d, %d, %d, %d)", p.X, p.Y, p.Height, p.Dist)
}

func (p *point) Set(a, b, c, d int) {
	p.X = a
	p.Y = b
	p.Height = c
	p.Dist = d
}

func (p point) Equals(other point) bool {
	return p.X == other.X &&
		p.Y == other.Y &&
		p.Height == other.Height &&
		p.Dist == other.Dist
}

type d12_Path []point

func (pq d12_Path) Len() int {
	return len(pq)
}

func (pq d12_Path) Less(i, j int) bool {
	return pq[i].Dist < pq[j].Dist
}

func (pq d12_Path) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *d12_Path) Push(x interface{}) {
	*pq = append(*pq, x.(point))
}

func (pq *d12_Path) Pop() interface{} {
	old := *pq
	n := len(old)
	x := old[n-1]
	*pq = old[0 : n-1]
	return x
}

func d12_Load(data string) (start, end point, m int2D) {
	// Open file
	f, err := os.Open(data)
	if err != nil {
		panic(fmt.Sprintf("Failed to open file: %s", data))
	}
	defer f.Close()

	input := bufio.NewScanner(f)
	for y := 0; input.Scan(); y++ {
		row := input.Text()
		m = append(m, make([]int, len(row)))
		for x := 0; x < len(row); x++ {
			if row[x] == 'S' {
				m[y][x] = 0
				start.Set(x, y, 0, 0)
			} else if row[x] == 'E' {
				dH := int('z' - 'a')
				m[y][x] = dH
				end.Set(x, y, dH, 0)
			} else {
				m[y][x] = int(row[x] - 'a')
			}
		}
	}
	return
}

// Return shorted path's distance.
// Assume heightMap is Y major.
func d12_Dijkstra(start, end point, heightMap int2D) int {
	fmt.Printf("start: %s, end: %s", start.ToString(), end.ToString())
	distances := make(int2D, len(heightMap))
	for y := range distances {
		distances[y] = make([]int, len(heightMap[y]))
		for x := range distances[y] {
			distances[y][x] = math.MaxInt32
		}
	}
	distances[start.X][start.Y] = 0

	path := make(d12_Path, 1)
	path[0] = start
	heap.Init(&path)

	for len(path) > 0 {
		curr := heap.Pop(&path).(point)
		if curr.X == end.X && curr.Y == end.Y {
			return curr.Dist
		}

		for _, dir := range [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}} {
			newX := curr.X + dir[0]
			newY := curr.Y + dir[1]
			if newX < 0 || newX >= len(heightMap[curr.Y]) || newY < 0 || newY >= len(heightMap) {
				continue
			}

			newH := heightMap[newY][newX]
			if newH > curr.Height+1 {
				continue
			}
			newDist := curr.Dist + 1
			if newDist < distances[newY][newX] {
				distances[newY][newX] = newDist
				heap.Push(&path, point{newX, newY, newH, newDist})
			}
		}
	}
	fmt.Println("height map:")
	heightMap.Print()
	fmt.Println("distances:")
	distances.Print()
	return -1
}
