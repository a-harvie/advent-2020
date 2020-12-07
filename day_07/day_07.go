package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/a-harvie/advent-2020/input"
)

type day7input map[string]bag

type bag struct {
	label    string
	children map[string]int
}

const target = "shiny gold"

func main() {
	inputTest := getDay7Input("./input_test")
	fmt.Println(inputTest)
	inputReal := getDay7Input("./input")

	fmt.Printf("Day 7 Part 1 Test: %v\n", day7Part1Solution(inputTest, target))
	fmt.Printf("Day 7 Part 1 Real: %v\n", day7Part1Solution(inputReal, target))

	fmt.Printf("Day 7 Part 2 Test: %v\n", day7Part2Solution(inputTest, target))
	fmt.Printf("Day 7 Part 2 Real: %v\n", day7Part2Solution(inputReal, target))

}

func day7Part1Solution(input day7input, target string) int {
	count := 0
	for _, b := range input {
		if bagContains(b, target, input) {
			count++
		}
	}

	return count
}

func bagContains(b bag, target string, bagList day7input) bool {
	if len(b.children) == 0 {
		return false
	}

	found := false
	for childName := range b.children {
		if childName == target {
			return true
		}
		found = found || bagContains(bagList[childName], target, bagList)
	}

	return found
}

func day7Part2Solution(input day7input, target string) int {
	return bagCount(input[target], input)
}

func bagCount(b bag, bagList day7input) int {
	if len(b.children) == 0 {
		return 0
	}

	count := 0
	for childName, childCount := range b.children {
		count += childCount + (childCount * bagCount(bagList[childName], bagList))
	}
	return count
}

func getDay7Input(filePath string) day7input {
	input, err := input.ReadInputFileLines(filePath)
	if err != nil {
		log.Fatal("Oh noes", err)
	}
	output := make(day7input, 0)

	/*
		light red bags contain 1 bright white bag, 2 muted yellow bags.
		dark orange bags contain 3 bright white bags, 4 muted yellow bags.
		bright white bags contain 1 shiny gold bag.
		muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.
		shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.
		dark olive bags contain 3 faded blue bags, 4 dotted black bags.
		vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.
		faded blue bags contain no other bags.
		dotted black bags contain no other bags.
	*/

	for _, line := range input {
		b := bag{}

		lineParts := strings.Split(line, " contain ")
		r := regexp.MustCompile("^([a-z ]+) bags$")
		parsedName := r.FindStringSubmatch(lineParts[0])
		b.label = parsedName[1]

		lineParts[1] = strings.TrimRight(lineParts[1], ".")
		childClauses := strings.Split(lineParts[1], ", ")
		b.children = make(map[string]int, 0)
		childRegex := regexp.MustCompile("^(\\d) (\\w+ \\w+) bag(s)?$")
		for _, childClause := range childClauses {
			if childClause == "no other bags" {
				break
			}
			parsedChildClause := childRegex.FindStringSubmatch(childClause)
			count, _ := strconv.Atoi(parsedChildClause[1])
			b.children[parsedChildClause[2]] = count
		}

		output[b.label] = b
	}

	return output
}
