package main

import (
	"fmt"
	"log"

	"github.com/a-harvie/advent-2020/input"
)

type cmd struct {
	op   string
	mask mask
	addr uint64
	val  uint64
}
type mask struct {
	zero uint64
	one  uint64
	str  string
}

func main() {
	inputTest := getDay15Input("./input_test", false)
	inputReal := getDay15Input("./input", false)

	fmt.Printf("Day 15 Part 1 Test: %v\n", day15Part1Solution(inputTest, 2020))
	fmt.Printf("Day 15 Part 1 Real: %v\n", day15Part1Solution(inputReal, 2020))

	// 	inputTest = getDay15Input("./input_test_2", true)
	// 	fmt.Println(inputTest)
	fmt.Printf("Day 15 Part 2 Test: %v\n", day15Part1Solution(inputTest, 30000000))
	fmt.Printf("Day 15 Part 2 Real: %v\n", day15Part1Solution(inputReal, 30000000))
}

func day15Part1Solution(input []int, lastTurn int) int {
	seen := make(map[int][]int)
	for i, n := range input {
		seen[n] = []int{i + 1}
	}
	last := input[len(input)-1]
	for turn := len(input) + 1; turn <= lastTurn; turn++ {
		lastSeen, lastWasSeen := seen[last]
		if !lastWasSeen || len(lastSeen) == 1 {
			last = 0
		} else {
			last = lastSeen[len(lastSeen)-1] - lastSeen[len(lastSeen)-2]
		}

		seen[last] = append(seen[last], turn)
	}
	return last
}

func day15Part2Solution(cmds []int) int {
	return 0
}

func getDay15Input(filePath string, getFloating bool) []int {
	input, err := input.ReadInputCSInts(filePath)
	if err != nil {
		log.Fatal("Oh noes", err)
	}
	return input
}
