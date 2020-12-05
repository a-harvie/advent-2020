package main

import (
	"fmt"
	"log"
	"math"
	"sort"
	"strings"

	"github.com/a-harvie/advent-2020/input"
)

func main() {
	/*
		BFFFBBFRRR: row 70, column 7, seat ID 567.
		FFFBBBFRRR: row 14, column 7, seat ID 119.
		BBFFBBFRLL: row 102, column 4, seat ID 820
	*/
	testInput := getInput("./input_test")
	realInput := getInput("./input")
	fmt.Printf("%#v\n", testInput)

	fmt.Printf("Day 4 Part 1 Test: %v\n", findGreatest(day5Part1Solution(testInput, true)))
	fmt.Printf("Day 4 Part 1 Real: %v\n", findGreatest(day5Part1Solution(realInput, false)))

	fmt.Printf("Day 4 Part 2 Real: %v\n", day5Part2Solution(realInput))
}

func day5Part1Solution(input []string, verbose bool) []int {
	ids := make([]int, 0)
	for _, seat := range input {
		chars := strings.Split(seat, "")

		row := binarySearch(chars[0:7], "F", 0.0, 127.0)
		col := binarySearch(chars[7:10], "L", 0.0, 7.0)

		id := row*8 + col
		if verbose {
			fmt.Printf("Row: %v col: %v seat id: %v\n", row, col, id)
		}

		ids = append(ids, int(id))
	}

	return ids
}

func binarySearch(searchSlice []string, lowerSymbol string, lower float64, upper float64) float64 {
	r := upper - lower
	for _, c := range searchSlice {
		if c == lowerSymbol {
			upper -= math.Round(r / 2)
		} else {
			lower += math.Round(r / 2)
		}
		r = math.Round(upper - lower)
	}
	return upper
}

func day5Part2Solution(input []string) int {
	ids := day5Part1Solution(input, false)
	sort.Ints(ids)

	for i, id := range ids {
		if i < len(ids)-1 && ids[i+1]-1 > id {
			return id + 1
		}
	}

	return 0
}

func findGreatest(ints []int) int {
	max := 0
	for _, i := range ints {
		if i > max {
			max = i
		}
	}
	return max
}

func getInput(fileName string) []string {
	input, err := input.ReadInputFileLines(fileName)
	if err != nil {
		log.Fatal("Oh noes", err)
	}
	return input
}
