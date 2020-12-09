package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/a-harvie/advent-2020/gameboi"
	"github.com/a-harvie/advent-2020/input"
)

func main() {
	inputTest := getDay8Input("./input_test")
	fmt.Println(inputTest)
	inputReal := getDay8Input("./input")

	fmt.Printf("Day 8 Part 1 Test: %v\n", day8Part1Solution(inputTest))
	fmt.Printf("Day 8 Part 1 Real: %v\n", day8Part1Solution(inputReal))

	fmt.Printf("Day 8 Part 2 Test: %v\n", day8Part2Solution(inputTest))
	fmt.Printf("Day 8 Part 2 Real: %v\n", day8Part2Solution(inputReal))
}

func day8Part1Solution(input []string) int64 {
	boi := gameboi.BoiMe(input)
	seen := make(map[int64]bool, 0)
	seen[boi.GetInstructionPointer()] = true
	var acc int64
	for {
		boi.Step()
		_, found := seen[boi.GetInstructionPointer()]
		if found {
			acc = boi.GetAccumulator()
			break
		}
		seen[boi.GetInstructionPointer()] = true
	}

	return acc
}

func day8Part2Solution(input []string) int64 {
	var acc int64
	fmt.Printf("Old instructions: %v\n", input)
InputMasher:
	for i := 0; i < len(input); i++ {
		newInstructions := mangleInstructions(input, i)
		boi := gameboi.BoiMe(newInstructions)
		seen := make(map[int64]bool, 0)
		seen[boi.GetInstructionPointer()] = true
		for {
			terminated := boi.Step()
			if terminated {
				acc = boi.GetAccumulator()
				break InputMasher
			}
			_, found := seen[boi.GetInstructionPointer()]
			if found {
				continue InputMasher
			}
			seen[boi.GetInstructionPointer()] = true
		}
	}

	return acc
}

func mangleInstructions(instructions []string, targetIndex int) []string {
	split := strings.Split(instructions[targetIndex], " ")
	var newInstruction string
	switch split[0] {
	case "acc":
		return instructions
	case "jmp":
		newInstruction = fmt.Sprintf("nop %s", split[1])
	case "nop":
		newInstruction = fmt.Sprintf("jmp %s", split[1])
	}
	newInstructions := make([]string, len(instructions))
	copy(newInstructions, instructions)
	newInstructions[targetIndex] = newInstruction
	return newInstructions
}

func getDay8Input(filePath string) []string {
	input, err := input.ReadInputFileLines(filePath)
	if err != nil {
		log.Fatal("Oh noes", err)
	}
	return input
}
