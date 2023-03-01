package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func d10() {
	d10_Part1("../data/d10.txt")
	// d10_Part2("../data/d10.txt")
}

type d10_Instruction struct {
	Inst int
	X    int
}

func (inst d10_Instruction) toString() string {
	if inst.Inst == d10_Noop {
		return "noop"
	}
	return fmt.Sprintf("addx %d", inst.X)
}

const (
	d10_Noop = iota
	d10_Addx
)

type d10_Memory []d10_Instruction
type d10_Address int

type d10_CPU struct {
	X   int
	Mem d10_Memory
	PC  d10_Address

	ExecCycle int
}

func (cpu *d10_CPU) toString() string {
	return fmt.Sprintf("{ X: %d, PC: %d, Inst: %s, ExecCycle: %d}\n", cpu.X, cpu.PC, cpu.Mem[cpu.PC].toString(), cpu.ExecCycle)
}

func (cpu *d10_CPU) Tick() (done bool) {
	if cpu.ExecCycle == 0 {
		if cpu.PC == d10_Address(len(cpu.Mem)-1) {
			done = true
			return
		}
		// Next instruction
		cpu.SetPC(cpu.PC + 1)
	}

	cpu.ExecCycle--
	inst := cpu.Mem[cpu.PC]
	if inst.Inst == d10_Addx && cpu.ExecCycle == 0 {
		cpu.X += inst.X
	}
	return
}

func (cpu *d10_CPU) SetPC(pc d10_Address) {
	cpu.PC = pc
	inst := cpu.Mem[pc]

	m := map[int]int{
		d10_Noop: 1,
		d10_Addx: 2,
	}
	cpu.ExecCycle = m[inst.Inst]
}

const (
	d10_FirstCycle    = 20
	d10_CycleInterval = 40
)

func d10_Part1(data string) (answer int) {
	// Prepare data
	cpu := &d10_CPU{
		X:   1,
		Mem: []d10_Instruction{},
	}

	// Open file
	f, err := os.Open(data)
	if err != nil {
		panic(fmt.Sprintf("Failed to open file: %s", data))
	}
	defer f.Close()

	// Scan line by line
	input := bufio.NewScanner(f)
	for input.Scan() {
		ss := strings.Split(input.Text(), " ")
		if ss[0] == "noop" {
			cpu.Mem = append(cpu.Mem, d10_Instruction{Inst: d10_Noop})
		} else if ss[0] == "addx" {
			x, err := strconv.Atoi(ss[1])
			if err != nil {
				panic(fmt.Sprintf("Failed to parse addx 'X': %s", ss[1]))
			}
			cpu.Mem = append(cpu.Mem, d10_Instruction{Inst: d10_Addx, X: x})
		}
	}

	cpu.SetPC(0)
	for cycle, done := 1, false; !done; cycle++ {
		// fmt.Println(cycle, ":", cpu.toString())
		if cycle == d10_FirstCycle || (cycle-d10_FirstCycle)%d10_CycleInterval == 0 {
			strength := cycle * cpu.X
			answer += strength
			fmt.Printf("**** cycle: %d, strength: %d, answer: %d\n", cycle, strength, answer)
		}
		done = cpu.Tick()
	}
	fmt.Println("[Day10 Part 1] answer: ", answer)
	return
}

func d10_Part2(data string) (answer int) {
	fmt.Println("[Day10 Part 2] answer: ", answer)
	return
}
