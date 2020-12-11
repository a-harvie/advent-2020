package main

import (
	"fmt"
	"log"
	"sort"

	"github.com/a-harvie/advent-2020/input"
)

func main() {
	inputTest := getDay10Input("./input_test")
	fmt.Println(inputTest)
	inputReal := getDay10Input("./input")

	fmt.Printf("Day 10 Part 1 Test: %v\n", day10Part1Solution(inputTest))
	fmt.Printf("Day 10 Part 1 Real: %v\n", day10Part1Solution(inputReal))

	fmt.Printf("Day 10 Part 2 Test: %v\n", day10Part2Solution(inputTest))
	fmt.Printf("Day 10 Part 2 Real: %v\n", day10Part2Solution(inputReal))
}

func day10Part1Solution(jolts []int) int {
	jolts = normalize(jolts)
	ones := 0
	threes := 0
	for i := 0; i < len(jolts)-1; i++ {
		switch d := jolts[i+1] - jolts[i]; d {
		case 1:
			ones++
		case 3:
			threes++
		}
	}
	return ones * threes
}

func day10Part2Solution(jolts []int) int {
	jolts = normalize(jolts)
	combos := 1
	ones := 0
	for i := 0; i < len(jolts)-1; i++ {
		switch d := jolts[i+1] - jolts[i]; d {
		case 1:
			ones++
		case 3:
			combos *= getCombos(ones)
			ones = 0
		}
	}

	return combos
}

func normalize(jolts []int) []int {
	sort.Ints(jolts)
	jolts = append([]int{0}, jolts...)
	jolts = append(jolts, jolts[len(jolts)-1]+3)
	return jolts
}

func getCombos(len int) int {
	switch len {
	case 0, 1: // zero is just to handle the "3" case without special handling in the pt2 function
		return 1
	case 2:
		return 2
	case 3:
		return 4
	default:
		return getCombos(len-1) + getCombos(len-2) + getCombos(len-3)
	}
}

func getDay10Input(filePath string) []int {
	input, err := input.ReadInputInts(filePath)
	if err != nil {
		log.Fatal("Oh noes", err)
	}
	return input
}
