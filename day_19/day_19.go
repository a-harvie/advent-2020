package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	//
	"github.com/a-harvie/advent-2020/errors"
	"github.com/a-harvie/advent-2020/input"
)

type rules map[string]string

var verbose = false

func main() {
	rulesTest, inputTest := getDay19Input("./input_test", false)
	rulesReal, inputReal := getDay19Input("./input", false)

	fmt.Printf("Day 19 Part 1 Test: %v\n", day19Part1Solution(rulesTest, inputTest))
	fmt.Printf("Day 19 Part 1 Real: %v\n", day19Part1Solution(rulesReal, inputReal))

	// fmt.Printf("Day 19 Part 2 Test: %v\n", day19Part2Solution(inputTest))
	fmt.Printf("Day 19 Part 2 Real: %v\n", day19Part2Solution(rulesReal, inputReal))
}

func day19Part1Solution(rules rules, input []string) int {
	count := 0
	regex := parseRegex(rules, false)
	log("regex", regex)
	r := regexp.MustCompile(regex)
	for _, line := range input {
		log(line)
		if r.MatchString(line) {
			count++
		}
	}
	return count
}

func day19Part2Solution(rules rules, input []string) int {
	regex := parseRegex(rules, true)
	// log("regex", regex)
	copy := strings.ReplaceAll(regex, "x", strconv.Itoa(1))
	count := getMatchCount(copy, input)
	prevCount := 0
	i := 2
	for {
		if prevCount == count {
			break
		}
		copy := strings.ReplaceAll(regex, "x", strconv.Itoa(i))
		prevCount = count
		count += getMatchCount(copy, input)
		i++
	}

	return count
}

func getMatchCount(regex string, input []string) int {
	count := 0
	r := regexp.MustCompile(regex)
	for _, line := range input {
		log(line)
		if r.MatchString(line) {
			count++
		}
	}
	return count
}

func parseRegex(rules rules, part2 bool) string {
	terminals := make(map[string]bool)
	for _, v := range rules {
		if v == "a" || v == "b" {
			terminals[v] = true
		}
	}

	r := ""
	split := strings.Split(rules["0"], " ")
	log("first split", split)
	for _, s := range split {
		r += parser(rules[s], rules, terminals, s, part2, 0)
	}
	return fmt.Sprintf("^%s$", r)
}
func parser(pattern string, rules rules, terminals map[string]bool, currentRule string, part2 bool, indent int) string {

	log(fmt.Sprintf("%*s%s %s %s %d", indent, "", "Parse", pattern, currentRule, indent))

	if pattern == "" {
		return pattern
	}
	isTerminal := terminals[pattern]
	log(fmt.Sprintf("%*s%s %v", indent, "", "isTerminal", isTerminal))
	if isTerminal {
		return pattern
	}

	r := make([]string, 0)
	orParts := strings.Split(pattern, " | ")
	for _, orPart := range orParts {
		log(fmt.Sprintf("%*s%s %s", indent, "", "orpart", orPart))
		sequenceParts := strings.Split(orPart, " ")
		sequence := ""
		for _, part := range sequenceParts {
			log(fmt.Sprintf("%*s%s %s", indent, "", "seqpart", part))
			p := parser(rules[part], rules, terminals, part, part2, (indent + 1))
			if part2 {
				if currentRule == "8" && part == "42" {
					p = p + "+"
				}
				if currentRule == "11" && (part == "42" || part == "31") {
					p = p + "{x}"
				}
			}
			sequence += p
		}
		r = append(r, sequence)
	}
	if len(orParts) > 1 {
		return fmt.Sprintf("(%s)", strings.Join(r, "|"))
	}
	return r[0]
}

func log(s ...interface{}) {
	if verbose {
		fmt.Println(s...)
	}
}

func getDay19Input(filePath string, getFloating bool) (rules, []string) {
	input, err := input.ReadStringChunks(filePath, "")
	errors.Check(err)
	rules := make(rules)
	for _, s := range input[0] {
		split := strings.Split(s, ": ")
		rules[split[0]] = strings.Trim(split[1], "\"")
	}
	return rules, input[1]
}
