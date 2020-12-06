package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/a-harvie/advent-2020/input"
)

type day6input [][]string

func main() {
	inputTest := getDay6Input("./input_test")
	fmt.Println(inputTest)
	inputReal := getDay6Input("./input")

	fmt.Printf("Day 6 Part 1 Test: %v\n", day6Part1Solution(inputTest))
	fmt.Printf("Day 6 Part 1 Real: %v\n", day6Part1Solution(inputReal))

	fmt.Printf("Day 6 Part 2 Test: %v\n", day6Part2Solution(inputTest))
	fmt.Printf("Day 6 Part 2 Real: %v\n", day6Part2Solution(inputReal))

}

func day6Part1Solution(input day6input) int {
	count := 0
	for _, group := range input {
		seen := getAllSeen(group)
		count += len(seen)
	}

	return count
}

func day6Part2Solution(input day6input) int {
	count := 0
	for _, group := range input {
		seen := getAllSeen(group)
		seenByAll := make([]string, 0)
		for _, item := range seen {
			missing := false
			for _, answer := range group {
				if !strings.Contains(answer, item) {
					missing = true
				}
			}
			if !missing {
				seenByAll = append(seenByAll, item)
			}
		}
		count += len(seenByAll)
	}

	return count
}

func getAllSeen(group []string) []string {
	seen := make([]string, 0)
	for _, answer := range group {
		for _, item := range strings.Split(answer, "") {
			if !wasSeen(seen, item) {
				seen = append(seen, item)
			}
		}
	}
	return seen
}

func wasSeen(seen []string, answer string) bool {
	for _, s := range seen {
		if s == answer {
			return true
		}
	}
	return false
}

func getDay6Input(filePath string) day6input {
	input, err := input.ReadInputFileLines(filePath)
	if err != nil {
		log.Fatal("Oh noes", err)
	}

	processed := make(day6input, 0)
	current := make([]string, 0)

	for _, line := range input {
		if line == "" {
			processed = append(processed, current)
			current = make([]string, 0)
			continue
		}

		current = append(current, line)
	}

	processed = append(processed, current)

	return processed
}
