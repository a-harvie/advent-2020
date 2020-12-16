package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/a-harvie/advent-2020/errors"
	"github.com/a-harvie/advent-2020/input"
)

type day16Input struct {
	fields   ticketFields
	myTicket []int
	tickets  [][]int
}

type ticketFields map[string][]int

func main() {
	inputTest := getDay16Input("./input_test", false)
	inputReal := getDay16Input("./input", false)

	fmt.Printf("Day 16 Part 1 Test: %v\n", day16Part1Solution(inputTest))
	fmt.Printf("Day 16 Part 1 Real: %v\n", day16Part1Solution(inputReal))

	inputTest = getDay16Input("./input_test_2", true)
	fmt.Printf("Day 16 Part 2 Test: %v\n", day16Part2Solution(inputTest))
	fmt.Printf("Day 16 Part 2 Real: %v\n", day16Part2Solution(inputReal))
}

func day16Part1Solution(input day16Input) int {
	invalid := make([]int, 0)

	for _, ticket := range input.tickets {
		for _, val := range ticket {
			if !ticketValueValidForAnyField(val, input.fields) {
				invalid = append(invalid, val)
			}
		}
	}

	errorRate := 0
	for _, i := range invalid {
		errorRate += i
	}
	return errorRate
}

func day16Part2Solution(input day16Input) int {
	tickets := getValidTickets(input.tickets, input.fields)
	fieldSlots := make(map[string]int)
	fieldsToSlot := make(map[string]bool)
	for fieldName := range input.fields {
		fieldsToSlot[fieldName] = false
	}
	for {
		if len(fieldsToSlot) == 0 {
			break
		}
		testColumn := 0
	columns:
		for {
			testFields := make(ticketFields)
			for fieldName := range fieldsToSlot {
				testFields[fieldName] = input.fields[fieldName]
			}
			for _, ticket := range tickets {
				for fieldName, fieldVals := range testFields {
					if !ticketValueValidForField(ticket[testColumn], fieldVals) {
						delete(testFields, fieldName)
					}
				}
			}
			if len(testFields) == 1 {
				for k := range testFields {
					fieldSlots[k] = testColumn
					delete(fieldsToSlot, k)
					break columns
				}
			}
			testColumn = (testColumn + 1) % len(tickets[0])
		}
	}

	fmt.Printf("Field slots: %#v\n", fieldSlots)
	total := 1
	for name, slot := range fieldSlots {
		if len(name) >= 9 && name[0:9] == "departure" {
			total *= input.myTicket[slot]
		}
	}
	return total
}

func ticketValueValidForAnyField(val int, fields ticketFields) bool {
	for _, fieldVals := range fields {
		if (val >= fieldVals[0] && val <= fieldVals[1]) ||
			(val >= fieldVals[2] && val <= fieldVals[3]) {
			return true
		}
	}
	return false
}

func ticketValueValidForField(val int, fieldVals []int) bool {
	if (val >= fieldVals[0] && val <= fieldVals[1]) ||
		(val >= fieldVals[2] && val <= fieldVals[3]) {
		return true
	}
	return false
}

func ticketValid(ticket []int, fields ticketFields) bool {
	for _, val := range ticket {
		if !ticketValueValidForAnyField(val, fields) {
			return false
		}
	}
	return true
}

func getValidTickets(tickets [][]int, fields ticketFields) [][]int {
	valid := make([][]int, 0)
	for _, ticket := range tickets {
		if ticketValid(ticket, fields) {
			valid = append(valid, ticket)
		}
	}
	return valid
}

func splitStrings(s []string, splitter string) [][]string {
	split := make([][]string, 0)
	currSlice := make([]string, 0)
	for _, str := range s {
		if str == splitter {
			split = append(split, currSlice)
			currSlice = make([]string, 0)
			continue
		}
		currSlice = append(currSlice, str)
	}
	split = append(split, currSlice)

	return split
}

func getDay16Input(filePath string, getFloating bool) day16Input {
	input, err := input.ReadInputFileLines(filePath)
	if err != nil {
		log.Fatal("Oh noes", err)
	}
	split := splitStrings(input, "")
	r := regexp.MustCompile("([a-z ]+): (\\d+)-(\\d+) or (\\d+)-(\\d+)")
	fields := make(ticketFields)
	for _, line := range split[0] {
		elements := r.FindStringSubmatch(line)
		low1, err := strconv.Atoi(elements[2])
		errors.Check(err)
		high1, err := strconv.Atoi(elements[3])
		errors.Check(err)
		low2, err := strconv.Atoi(elements[4])
		errors.Check(err)
		high2, err := strconv.Atoi(elements[5])
		errors.Check(err)
		fields[elements[1]] = []int{low1, high1, low2, high2}
	}

	myTicket := make([]int, 0)
	myTicketParts := strings.Split(split[1][1], ",")
	for _, part := range myTicketParts {
		i, err := strconv.Atoi(part)
		errors.Check(err)
		myTicket = append(myTicket, i)
	}

	tickets := make([][]int, 0)
	for i := 1; i < len(split[2]); i++ {
		ticket := make([]int, 0)
		ticketParts := strings.Split(split[2][i], ",")
		for _, part := range ticketParts {
			i, err := strconv.Atoi(part)
			errors.Check(err)
			ticket = append(ticket, i)
		}
		tickets = append(tickets, ticket)
	}
	return day16Input{
		fields:   fields,
		myTicket: myTicket,
		tickets:  tickets,
	}
}
