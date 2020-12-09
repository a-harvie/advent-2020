package main

import (
	"fmt"
	"log"
	"math"
	"strconv"

	"github.com/a-harvie/advent-2020/input"
)

func main() {
	inputTest := getDay9Input("./input_test")
	fmt.Println(inputTest)
	inputReal := getDay9Input("./input")

	fmt.Printf("Day 9 Part 1 Test: %v\n", day9Part1Solution(inputTest, 5))
	fmt.Printf("Day 9 Part 1 Real: %v\n", day9Part1Solution(inputReal, 25))

	fmt.Printf("Day 9 Part 2 Test: %v\n", day9Part2Solution(inputTest, 5))
	fmt.Printf("Day 9 Part 2 Real: %v\n", day9Part2Solution(inputReal, 25))
}

func day9Part1Solution(input []int, preamble int) int {
	for i := preamble; i < len(input); i++ {
		if !valid(input[i], input[i-preamble:i]) {
			return input[i]
		}
	}

	return -1
}

func valid(candidate int, precursors []int) bool {
	for i, precursorA := range precursors {
		for j, precursorB := range precursors {
			if i != j && precursorA+precursorB == candidate {
				return true
			}
		}
	}

	return false
}

func day9Part2Solution(input []int, preamble int) int {
	invalid := day9Part1Solution(input, preamble)
	sums := getSums(invalid, input)
	return min(sums) + max(sums)
}

func getSums(target int, candidates []int) []int {
sumsouter:
	for i := 0; i < len(candidates); i++ {
		sum := candidates[i]
		if sum >= target {
			continue
		}
		for j := i + 1; j < len(candidates); j++ {
			sum += candidates[j]
			if sum > target || j == len(candidates)-1 {
				continue sumsouter
			}
			if sum == target {
				return candidates[i : j+1]
			}
		}
	}

	return make([]int, 0)
}

func min(ints []int) int {
	min := math.MaxInt64
	for _, i := range ints {
		if i < min {
			min = i
		}
	}
	return min
}

func max(ints []int) int {
	max := math.MinInt64
	for _, i := range ints {
		if i > max {
			max = i
		}
	}
	return max
}

func getDay9Input(filePath string) []int {
	input, err := input.ReadInputFileLines(filePath)
	if err != nil {
		log.Fatal("Oh noes", err)
	}
	output := make([]int, len(input))
	for i, line := range input {
		integer, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal("Oh noes", err)
		}
		output[i] = integer
	}
	return output
}
