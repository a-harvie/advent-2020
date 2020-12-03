package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/a-harvie/advent-2020/input"
)

func main() {
	testInput, err := input.ReadInputFileLines("./input_test")
	if err != nil {
		log.Fatal("Oh noes", err)
	}
	parsedTestInput := parseDay3Input(testInput)
	// fmt.Println(parsedTestInput)
	input, err := input.ReadInputFileLines("./input")
	if err != nil {
		log.Fatal("Oh noes", err)
	}
	parsedInput := parseDay3Input(input)
	// fmt.Println(parsedInput)

	fmt.Printf("Day 2 Part 1 Test: %v\n", day3Part1Solution(parsedTestInput, slope{3, 1}))
	fmt.Printf("Day 2 Part 1 Real: %v\n", day3Part1Solution(parsedInput, slope{3, 1}))

	slopeList := []slope{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	}
	testProduct := day3Part2Solution(parsedTestInput, slopeList)
	realProduct := day3Part2Solution(parsedInput, slopeList)

	fmt.Printf("Day 2 Part 2 Test: %v\n", testProduct)
	fmt.Printf("Day 2 Part 2 Real: %v\n", realProduct)
}

func day3Part1Solution(input day3Input, s slope) int {
	count := 0
	currentX := 0
	currentY := 0
	width := len(input[0])
	for {
		currentX = (currentX + s.X) % width
		currentY += s.Y
		if currentY >= len(input) {
			break
		}
		if input[currentY][currentX] {
			count++
		}
	}

	return count
}

func day3Part2Solution(input day3Input, slopeList []slope) int {
	product := 1
	for _, s := range slopeList {
		product *= day3Part1Solution(input, s)
	}
	return product
}

type day3Input [][]bool
type slope struct {
	X int
	Y int
}

func parseDay3Input(input []string) day3Input {
	parsed := make([][]bool, len(input))

	for i, line := range input {
		symbols := strings.Split(line, "")
		parsedLine := make([]bool, len(symbols))
		for j, symbol := range symbols {
			if symbol == "#" {
				parsedLine[j] = true
			}
		}
		parsed[i] = parsedLine
	}

	return parsed
}
