package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func d07() {
	d07_Part1("../data/d07.txt")
	d07_Part2("../data/d07.txt")
}

type d07_Node struct {
	Name string
	Type d07_Type
	// For directory, size means total size of files under this node.
	Size     int
	Parent   *d07_Node
	Children []*d07_Node
}

type d07_Type int

const (
	d07_File = iota
	d07_Dir
)

func (n *d07_Node) FindChild(name string) (child *d07_Node) {
	for _, c := range n.Children {
		if c.Name == name {
			return c
		}
	}
	panic(fmt.Sprintf("Cannot find child with name: %s in node: %s", name, n.Name))
}

func (n *d07_Node) Print(prefix string) {
	if n.Type == d07_Dir {
		fmt.Printf("%s %s (dir, size=%d)\n", prefix, n.Name, n.Size)
	} else {
		fmt.Printf("%s %s (file, size=%d)\n", prefix, n.Name, n.Size)
	}
	for _, c := range n.Children {
		c.Print("  " + prefix)
	}
}

func (n *d07_Node) SumChildrenSizes() {
	for _, c := range n.Children {
		n.Size += c.Size
	}
}

func d07_Part1(data string) (answer int) {
	f, err := os.Open(data)
	if err != nil {
		panic(fmt.Sprintf("Failed to open file: %s", data))
	}
	defer f.Close()

	root := &d07_Node{
		Name:     "/",
		Type:     d07_Dir,
		Parent:   nil,
		Children: []*d07_Node{},
	}
	input := bufio.NewScanner(f)
	// Skip first line for it always is '$cd /' which has no effects.
	input.Scan()
	d07_Parse(input, root, root)
	// root.Print("- ")
	answer = d07_SumDirSizes(root, 100000)
	fmt.Println("[Day07 Part 1] answer: ", answer)
	return
}

func d07_Part2(data string) (answer int) {
	f, err := os.Open(data)
	if err != nil {
		panic(fmt.Sprintf("Failed to open file: %s", data))
	}
	defer f.Close()

	root := &d07_Node{
		Name:     "/",
		Type:     d07_Dir,
		Parent:   nil,
		Children: []*d07_Node{},
	}
	input := bufio.NewScanner(f)
	// Skip first line for it always is '$cd /' which has no effects.
	input.Scan()
	d07_Parse(input, root, root)

	const (
		total = 70000000
		need  = 30000000
	)
	free := total - root.Size
	if free >= need {
		return
	}

	dirSizeLimit := need - free
	fmt.Println("dirSizeLimit: ", dirSizeLimit)
	dir := d07_FindMinDir(root, root, dirSizeLimit)
	answer = dir.Size
	fmt.Println("[Day 7 Part 2] answer: ", answer)
	return
}

func d07_Parse(input *bufio.Scanner, root, curr *d07_Node) {
	for input.Scan() {
		line := input.Text()

		if strings.HasPrefix(line, "$ cd") {
			if strings.HasSuffix(line, "..") {
				// Before go up, sum sizes of children.
				curr.SumChildrenSizes()
				return
			} else {
				// cd to directory
				dirName := line[5:]
				dir := curr.FindChild(dirName)
				d07_Parse(input, root, dir)
			}
		} else if strings.HasPrefix(line, "$ ls") {
			// keep parsing
		} else if strings.HasPrefix(line, "dir ") {
			dir := &d07_Node{
				Name:     line[4:],
				Type:     d07_Dir,
				Parent:   curr,
				Children: []*d07_Node{},
			}
			curr.Children = append(curr.Children, dir)
		} else {
			// Should be file: '123 abc'
			ss := strings.Split(line, " ")
			size, err := strconv.Atoi(ss[0])
			if err != nil {
				panic(fmt.Sprintf("Failed to parse file size: %s", ss[0]))
			}
			file := &d07_Node{
				Name:     ss[1],
				Type:     d07_File,
				Size:     size,
				Parent:   curr,
				Children: []*d07_Node{},
			}
			curr.Children = append(curr.Children, file)
		}
	}
	// After parsed last line, sum up children's file sizes.
	if curr.Type == d07_Dir {
		curr.SumChildrenSizes()
	}
	return
}

func d07_SumDirSizes(node *d07_Node, limit int) (sum int) {
	if node.Type == d07_Dir && node.Size <= limit {
		sum += node.Size
	}
	for _, c := range node.Children {
		if c.Type == d07_Dir {
			sum += d07_SumDirSizes(c, limit)
		}
	}
	return
}
func d07_FindMinDir(node, currMin *d07_Node, sizeLimit int) (min *d07_Node) {
	if node.Size >= sizeLimit && node.Size < currMin.Size {
		currMin = node
	}
	for _, c := range node.Children {
		if c.Type == d07_Dir {
			currMin = d07_FindMinDir(c, currMin, sizeLimit)
		}
	}
	return currMin
}
