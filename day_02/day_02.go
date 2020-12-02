package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/a-harvie/advent-2020/input"
)

func main() {
	testInput, err := input.ReadInputFileLines("./input_test")
	if err != nil {
		log.Fatal("Oh noes", err)
	}
	parsedTestInput := parseDay2Input(testInput)
	fmt.Println(parsedTestInput)
	input, err := input.ReadInputFileLines("./input")
	if err != nil {
		log.Fatal("Oh noes", err)
	}
	parsedInput := parseDay2Input(input)
	// fmt.Println(parsedInput)

	fmt.Printf("Day 2 Part 1 Test: %v\n", day2Part1Solution(parsedTestInput))
	fmt.Printf("Day 2 Part 1 Real: %v\n", day2Part1Solution(parsedInput))
	fmt.Printf("Day 2 Part 2 Test: %v\n", day2Part2Solution(parsedTestInput))
	fmt.Printf("Day 2 Part 2 Real: %v\n", day2Part2Solution(parsedInput))
}

func day2Part1Solution(input []day2Input) int {
	count := 0
	for _, entry := range input {
		if passwordValidPart1(entry.Password, entry.SearchChar, entry.Min, entry.Max) {
			count++
		}
	}
	return count
}

func day2Part2Solution(input []day2Input) int {
	count := 0
	for _, entry := range input {
		if passwordValidPart2(entry.Password, entry.SearchChar, entry.Min, entry.Max) {
			count++
		}
	}
	return count
}

func passwordValidPart1(password string, requiredChar string, min int, max int) bool {
	count := strings.Count(password, requiredChar)
	return count >= min && count <= max
}

func passwordValidPart2(password string, requiredChar string, pos1 int, pos2 int) bool {
	// (Be careful; Toboggan Corporate Policies have no concept of "index zero"!)
	c1 := password[pos1-1 : pos1]
	c2 := password[pos2-1 : pos2]
	return (c1 == requiredChar && c2 != requiredChar) || (c1 != requiredChar && c2 == requiredChar)
}

type day2Input struct {
	Min        int
	Max        int
	SearchChar string
	Password   string
}

func parseDay2Input(input []string) []day2Input {
	parsed := make([]day2Input, len(input))
	for i, line := range input {
		lineParts := strings.Split(line, ": ")
		metaParts := strings.Split(lineParts[0], " ")
		minMax := strings.Split(metaParts[0], "-")
		min, _ := strconv.Atoi(minMax[0])
		max, _ := strconv.Atoi(minMax[1])
		parsed[i] = day2Input{
			Min:        min,
			Max:        max,
			SearchChar: metaParts[1],
			Password:   lineParts[1],
		}

	}

	return parsed
}
