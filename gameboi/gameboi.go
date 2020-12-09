package gameboi

import (
	"fmt"
	"strconv"
	"strings"
)

// Gameboi is an AoC opcode thing
type Gameboi struct {
	acc          int64
	instrPointer int64
	program      []instruction
}

type instruction struct {
	cmd    string
	params []int64
}

// BoiMe will give you a Gameboi
func BoiMe(program []string) *Gameboi {
	g := Gameboi{}
	instructions := parseInstructions(program)
	g.program = instructions
	return &g
}

func parseInstructions(program []string) []instruction {
	instructions := make([]instruction, len(program))
	for i, line := range program {
		parts := strings.Split(line, " ")
		param, _ := strconv.Atoi(parts[1])
		instructions[i] = instruction{
			cmd:    parts[0],
			params: []int64{int64(param)},
		}
	}
	return instructions
}

// Step will step once through the program, or return a terminated signal if at the end of the program
func (g *Gameboi) Step() bool {
	i := g.program[g.instrPointer]
	switch i.cmd {
	case "acc":
		g.doAcc(i.params[0])
	case "jmp":
		g.doJmp(i.params[0])
	case "nop":
		g.doNop()

	}

	if g.instrPointer == int64(len(g.program)) {
		return true
	}

	return false
}

func (g *Gameboi) doAcc(param int64) {
	g.acc += param
	g.instrPointer++
}

func (g *Gameboi) doJmp(param int64) {
	g.instrPointer += param
}

func (g *Gameboi) doNop() {
	g.instrPointer++
}

// GetInstructionPointer will return the current instruction pointer
func (g *Gameboi) GetInstructionPointer() int64 {
	return g.instrPointer
}

// GetAccumulator will return the current acc value
func (g *Gameboi) GetAccumulator() int64 {
	return g.acc
}

// GetProgram will return the current program formatted to strings
func (g *Gameboi) GetProgram() []string {
	out := make([]string, len(g.program))
	for i, line := range g.program {
		out[i] = fmt.Sprintf("%d: %#v", i, line)
	}

	return out
}

// GetState will return the gameboi's current state
func (g *Gameboi) GetState() string {
	return fmt.Sprintf("Acc: %d Ptr: %d Instr: %#v\n", g.acc, g.instrPointer, g.program[g.instrPointer])
}
