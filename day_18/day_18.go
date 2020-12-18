package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/a-harvie/advent-2020/errors"
	"github.com/a-harvie/advent-2020/input"
)

type tree struct {
	op    string
	left  *tree
	right *tree
}

func main() {
	inputTest := getDay18Input("./input_test", false)
	// inputReal := getDay18Input("./input", false)

	fmt.Printf("Day 18 Part 1 Test: %v\n", day18Part1Solution(inputTest))
	// fmt.Printf("Day 18 Part 1 Real: %v\n", day18Part1Solution(inputReal))

	// inputTest = getDay18Input("./input_test_2", true)
	// fmt.Printf("Day 18 Part 2 Test: %v\n", day18Part2Solution(inputTest))
	// fmt.Printf("Day 18 Part 2 Real: %v\n", day18Part2Solution(inputReal))
}

func day18Part1Solution(input [][]string) int {
	total := 0
	for _, line := range input {
		total += mathTime(line)
	}
	return 0
}

func day18Part2Solution(input []string) int {
	return 0
}

func mathTime(eq []string) int {
	t := makeTree(eq)
	return t.math()
}

func makeTree(eq []string) *tree {
	fmt.Println("maketree", eq)
	t := tree{}
	// r := regexp.MustCompile("\\d+")
	if len(eq) == 1 {
		t.op = eq[0]
		fmt.Println("tree len 1", t)
		return &t
	}
	for i, sym := range eq {
		if sym == "(" {
			subeq := eq[i+1 : findClosing(eq)]
			fmt.Println("subtree", subeq)
			return makeTree(subeq)
		}
		if (i + 1) == len(eq) {
			break
		}
		next := eq[i+1]
		if next == "*" || next == "+" {
			t.op = next
			fmt.Println("optree", t)
			t.left = makeTree(eq[:i+1])
			t.right = makeTree(eq[i+2:])
		}

	}
	fmt.Println("endtree", t)
	return &t
}

func findClosing(eq []string) int {
	stack := 1
	for i := 1; i < len(eq); i++ {
		fmt.Println(stack, eq[i])

		switch eq[i] {
		case "(":
			stack++
		case ")":
			stack--
		default:
			continue

		}
		if stack == 0 {
			return i
		}
	}
	return 0
}

func (t *tree) math() int {
	i := 0
	switch t.op {
	case "+":
		i = t.left.math() + t.right.math()
	case "*":
		i = t.left.math() * t.right.math()
	default:
		n, _ := strconv.Atoi(t.op)
		i = n
	}
	return i
}

func getClosing(symbols []string) int {
	stack := 0
	i := 0
	for {
		if i == len(symbols) || (i > 1 && stack == 0) {
			break
		}
		switch symbols[i] {
		case "(":
			stack++
		case ")":
			stack--
		default:
			continue
		}

		i++
	}
	return i
}

func getDay18Input(filePath string, getFloating bool) [][]string {
	input, err := input.ReadInputFileLines(filePath)
	errors.Check(err)
	parsed := make([][]string, 0)
	for _, line := range input {
		split := strings.Split(line, " ")
		symbols := make([]string, 0)
		for i := 0; i < len(split); i++ {
			if len(split[i]) > 1 {
				subs := strings.Split(split[i], "")
				for j := 0; j < len(subs); j++ {
					symbols = append(symbols, subs[j])
				}
			} else {
				symbols = append(symbols, split[i])
			}
		}
		parsed = append(parsed, symbols)
	}
	return parsed
}
