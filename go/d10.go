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
	d10_Part2("../data/d10.txt")
}

func d10_Part1(data string) (answer int) {
	device := d10_NewDevice(data)

	answer = device.RunForSignalStrength(20, 40)
	fmt.Println("[Day10 Part 1] answer: ", answer)
	return
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

func d10_NewCPU(data string) (cpu *d10_CPU) {
	cpu = &d10_CPU{
		X:   1,
		Mem: []d10_Instruction{},
	}
	cpu.LoadInstructions(data)
	cpu.SetPC(0)
	return
}

func (cpu *d10_CPU) toString() string {
	return fmt.Sprintf("{ X: %d, PC: %d, Inst: %s, ExecCycle: %d}\n", cpu.X, cpu.PC, cpu.Mem[cpu.PC].toString(), cpu.ExecCycle)
}

func (cpu *d10_CPU) LoadInstructions(data string) {
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

func d10_Part2(data string) {
	device := d10_NewDevice(data)
	device.RunForCRT()
	device.CRT.Display()
}

type d10_CRT struct {
	Rows       [6][]rune
	RowIndex   int
	PixelIndex int
}

func d10_NewCRT() (crt *d10_CRT) {
	crt = &d10_CRT{
		Rows: [6][]rune{
			[]rune("########################################"),
			[]rune("########################################"),
			[]rune("########################################"),
			[]rune("########################################"),
			[]rune("########################################"),
			[]rune("########################################"),
		},
		RowIndex:   0,
		PixelIndex: 0,
	}
	return
}

func (crt *d10_CRT) Draw(sprite d10_Sprite) {
	if crt.RowIndex >= len(crt.Rows) {
		fmt.Printf("Out of buffer, don't draw. Sprite: [%d,%d,%d]\n", sprite[0], sprite[1], sprite[2])
		return
	}

	row := crt.Rows[crt.RowIndex]

	overlap := false
	for _, px := range sprite {
		if px == crt.PixelIndex {
			overlap = true
			break
		}
	}

	if !overlap {
		row[crt.PixelIndex] = '.'
	}
	fmt.Printf("row: %s, sprite: [%d,%d,%d]\n", string(row), sprite[0], sprite[1], sprite[2])

	crt.PixelIndex++
	if crt.PixelIndex >= 40 {
		crt.RowIndex++
		crt.PixelIndex = 0
	}
}

func (crt *d10_CRT) Display() {
	for _, row := range crt.Rows {
		fmt.Println(string(row))
	}
}

type d10_Device struct {
	CPU            *d10_CPU
	CRT            *d10_CRT
	Cycle          int
	SignalStrength int
}

func d10_NewDevice(data string) *d10_Device {
	return &d10_Device{
		CPU:            d10_NewCPU(data),
		CRT:            d10_NewCRT(),
		Cycle:          1,
		SignalStrength: 0,
	}
}

// strOffset & strInterval controls when device calculate and accumulate signal
// strength.
func (device *d10_Device) RunForSignalStrength(strOffset, strInterval int) (strength int) {
	done := false
	for !done {
		// fmt.Println(cycle, ":", cpu.toString())
		if device.Cycle == strOffset || (device.Cycle-strOffset)%strInterval == 0 {
			strength := device.Cycle * device.CPU.X
			device.SignalStrength += strength
			fmt.Printf("**** cycle: %d, strength: %d, answer: %d\n", device.Cycle, strength, device.SignalStrength)
		}

		done = device.CPU.Tick()
		device.Cycle++
	}
	return device.SignalStrength
}

type d10_Sprite [3]int

func (device *d10_Device) RunForCRT() {
	done := false
	sprite := d10_Sprite{0, 1, 2}
	for !done {
		x := device.CPU.X
		sprite[0] = x - 1
		sprite[1] = x
		sprite[2] = x + 1

		device.CRT.Draw(sprite)
		done = device.CPU.Tick()
	}
}
