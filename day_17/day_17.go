package main

import (
	"fmt"
	"strings"

	"github.com/a-harvie/advent-2020/errors"
	"github.com/a-harvie/advent-2020/input"
	"github.com/a-harvie/advent-2020/nlifeboi"
)

func main() {
	inputTest := getDay17Input1("./input_test")
	inputReal := getDay17Input1("./input")

	fmt.Printf("Day 17 Part 1 Test: %v\n", day17Part1Solution(inputTest))
	fmt.Printf("Day 17 Part 1 Real: %v\n", day17Part1Solution(inputReal))

	inputTest2 := getDay17Input2("./input_test")
	inputReal2 := getDay17Input2("./input")
	fmt.Printf("Day 17 Part 2 Test: %v\n", day17Part2Solution(inputTest2))
	fmt.Printf("Day 17 Part 2 Real: %v\n", day17Part2Solution(inputReal2))
}

func day17Part1Solution(input nlifeboi.ThrGrid) int {
	b := nlifeboi.BoiMe(input)
	// render(0, b.GetGrid())
	for i := 0; i < 6; i++ {
		b.Step()
		// render(i+1, b.GetGrid())
	}
	return b.GetPopulation()
}

func day17Part2Solution(input nlifeboi.ForGrid) int {
	b := nlifeboi.HyperBoiMe(input)
	// render(0, b.GetGrid())
	for i := 0; i < 6; i++ {
		b.Step()
		// render(i+1, b.GetGrid())
	}
	return b.GetPopulation()
}

func render(step int, grid nlifeboi.ThrGrid) {
	fmt.Println("Step", step)
	for i := range grid {
		fmt.Println("z ", i)
		for j := range grid[i] {
			for k := range grid[i][j] {
				if grid[i][j][k] > 0 {
					fmt.Print("#")
				} else {
					fmt.Print(".")
				}
			}
			fmt.Print("\n")
		}
		fmt.Print("\n")
	}
}

func getDay17Input1(filePath string) nlifeboi.ThrGrid {
	input, err := input.ReadInputFileLines(filePath)
	errors.Check(err)

	// shooooould make this dynamically resize instead…
	size := len(input) * 6 // 3 is enough for the real input, but not test :p
	offset := size / 2
	grid := make(nlifeboi.ThrGrid, size)
	for i := 0; i < size; i++ {
		plane := make(nlifeboi.TwoGrid, size)
		for j := 0; j < size; j++ {
			plane[j] = make(nlifeboi.OneGrid, size)
		}
		grid[i] = plane
	}
	for i, line := range input {
		chars := strings.Split(line, "")
		for j, char := range chars {
			if char == "#" {
				grid[offset][i+offset][j+offset] = 1
			}
		}
	}

	return grid
}

func getDay17Input2(filePath string) nlifeboi.ForGrid {
	input, err := input.ReadInputFileLines(filePath)
	errors.Check(err)

	size := len(input) * 6 // ¯\_(ツ)_/¯
	offset := size / 2
	grid := make(nlifeboi.ForGrid, size)
	for i := 0; i < size; i++ {
		cube := make(nlifeboi.ThrGrid, size)
		for j := 0; j < size; j++ {
			plane := make(nlifeboi.TwoGrid, size)
			for k := 0; k < size; k++ {
				plane[k] = make(nlifeboi.OneGrid, size)
			}
			cube[j] = plane
		}
		grid[i] = cube
	}
	for i, line := range input {
		chars := strings.Split(line, "")
		for j, char := range chars {
			if char == "#" {
				grid[offset][offset][i+offset][j+offset] = 1
			}
		}
	}

	return grid
}
