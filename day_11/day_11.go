package main

import (
	"fmt"
	"log"

	"github.com/a-harvie/advent-2020/input"

	"github.com/a-harvie/advent-2020/lifeboi"
)

func main() {
	inputTest := getDay11Input("./input_test")
	testBoi := getBoi(inputTest)
	fmt.Println(inputTest)
	inputReal := getDay11Input("./input")
	realBoi := getBoi(inputReal)

	fmt.Printf("Day 11 Part 1 Test: %v\n", day11Part1Solution(testBoi))
	fmt.Printf("Day 11 Part 1 Real: %v\n", day11Part1Solution(realBoi))

	testBoi = getBoi(inputTest)
	fmt.Printf("Day 11 Part 2 Test: %v\n", day11Part2Solution(testBoi))
	realBoi = getBoi(inputReal)
	fmt.Printf("Day 11 Part 2 Real: %v\n", day11Part2Solution(realBoi))
}

func day11Part1Solution(boi *lifeboi.Lifeboi) int {
	lastPop := 0
	stableFor := 0
	for {
		boi.Step(4, "neighbour")
		if boi.GetPopulation() == lastPop {
			if stableFor == 0 {
				break
			}
			stableFor--
		}
		lastPop = boi.GetPopulation()
	}

	return boi.GetPopulation()
}

func day11Part2Solution(boi *lifeboi.Lifeboi) int {
	lastPop := 0
	stableFor := 0
	for {
		boi.Step(5, "viewable")
		if boi.GetPopulation() == lastPop {
			if stableFor == 0 {
				break
			}
			stableFor--
		}
		lastPop = boi.GetPopulation()
	}

	return boi.GetPopulation()
}

func getDay11Input(filePath string) []string {
	input, err := input.ReadInputFileLines(filePath)
	if err != nil {
		log.Fatal("Oh noes", err)
	}
	return input
}

func getBoi(input []string) *lifeboi.Lifeboi {
	states := map[string]int{
		".": lifeboi.Barren,
		"L": lifeboi.Empty,
		"#": lifeboi.Full,
	}

	b, err := lifeboi.BoiMe(input, states)
	if err != nil {
		log.Fatal("Oh noes", err)
	}
	return b
}
