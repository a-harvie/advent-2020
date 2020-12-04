package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/a-harvie/advent-2020/input"
)

type passport struct {
	Byr string
	Iyr string
	Eyr string
	Hgt string
	Hcl string
	Ecl string
	Pid string
	Cid string
}

func main() {
	testInput := getInput("./input_test")
	realInput := getInput("./input")
	fmt.Printf("%#v\n", testInput)
	fmt.Printf("Day 4 Part 1 Test (%v): %v\n", len(testInput), day4Part1Solution(testInput))
	fmt.Printf("Day 4 Part 1 Real (%v): %v\n", len(realInput), day4Part1Solution(realInput))

	testInputPt2 := getInput("./input_test_pt2")
	fmt.Printf("%#v\n", testInputPt2)
	fmt.Printf("Day 4 Part 2 Test (%v): %v\n", len(testInputPt2), day4Part2Solution(testInputPt2))
	fmt.Printf("Day 4 Part 2 Real (%v): %v\n", len(realInput), day4Part2Solution(realInput))
}

func day4Part1Solution(input []passport) int {
	count := 0
	for _, pp := range input {
		if pp.Byr != "" &&
			pp.Iyr != "" &&
			pp.Eyr != "" &&
			pp.Hgt != "" &&
			pp.Hcl != "" &&
			pp.Ecl != "" &&
			pp.Pid != "" {
			count++
		}
	}

	return count
}

func day4Part2Solution(input []passport) int {
	count := 0
	for i, pp := range input {

		if !validateYearString(pp.Byr, 1920, 2002) {
			fmt.Printf("Passport %v failed %s with %v\n", i, "Byr", pp.Byr)
			continue
		}
		if !validateYearString(pp.Iyr, 2010, 2020) {
			fmt.Printf("Passport %v failed %s with %v\n", i, "Iyr", pp.Iyr)
			continue
		}
		if !validateYearString(pp.Eyr, 2020, 2030) {
			fmt.Printf("Passport %v failed %s with %v\n", i, "Eyr", pp.Eyr)
			continue
		}
		if !validateHeight(pp.Hgt) {
			fmt.Printf("Passport %v failed %s with %v\n", i, "Hgt", pp.Hgt)
			continue
		}
		if !validateHcl(pp.Hcl) {
			fmt.Printf("Passport %v failed %s with %v\n", i, "Hcl", pp.Hcl)
			continue
		}
		if !validateEcl(pp.Ecl) {
			fmt.Printf("Passport %v failed %s with %v\n", i, "Ecl", pp.Ecl)
			continue
		}
		if !validatePid(pp.Pid) {
			fmt.Printf("Passport %v failed %s with %v\n", i, "Pid", pp.Pid)
			continue
		}

		count++
	}

	return count
}

func validateYearString(input string, min int, max int) bool {
	year, err := strconv.Atoi(input)
	if err != nil {
		return false
	}
	if year < min || year > max {
		return false
	}
	return true
}

func validateHeight(height string) bool {
	r, _ := regexp.Compile("(\\d+)(cm|in)")
	elems := r.FindStringSubmatch(height)
	if len(elems) != 3 {
		return false
	}
	units := elems[2]
	h, err := strconv.Atoi(elems[1])
	if err != nil ||
		(units == "cm" && (h < 150 || h > 193)) ||
		(units == "in" && (h < 59 || h > 76)) {
		return false
	}

	return true
}

func validateHcl(hcl string) bool {
	r, _ := regexp.Compile("^#[a-f0-9A-F]{6}$")
	if !r.MatchString(hcl) {
		return false
	}
	return true
}

func validateEcl(ecl string) bool {
	if ecl != "amb" && ecl != "blu" && ecl != "brn" && ecl != "gry" && ecl != "grn" && ecl != "hzl" && ecl != "oth" {
		return false
	}
	return true
}

func validatePid(pid string) bool {
	r, _ := regexp.Compile("^\\d{9}$")
	if !r.MatchString(pid) {
		return false
	}
	return true
}

func getInput(fileName string) []passport {
	testInput, err := input.ReadInputFileLines(fileName)
	if err != nil {
		log.Fatal("Oh noes", err)
	}
	parsedInput := parseDay4Input(testInput)
	return parsedInput
}

func parseDay4Input(input []string) []passport {
	passports := make([]passport, 0)
	currentPassport := passport{}
	for _, line := range input {
		if line == "" {
			passports = append(passports, currentPassport)
			currentPassport = passport{}
		}
		fields := parseDay4Line(line)
		for _, field := range fields {
			switch field[0] {
			case "byr":
				currentPassport.Byr = field[1]
			case "iyr":
				currentPassport.Iyr = field[1]
			case "eyr":
				currentPassport.Eyr = field[1]
			case "hgt":
				currentPassport.Hgt = field[1]
			case "hcl":
				currentPassport.Hcl = field[1]
			case "ecl":
				currentPassport.Ecl = field[1]
			case "pid":
				currentPassport.Pid = field[1]
			case "cid":
				currentPassport.Cid = field[1]
			}
		}
	}
	passports = append(passports, currentPassport)

	return passports
}

func parseDay4Line(line string) [][]string {
	parsed := make([][]string, 0)
	fields := strings.Split(line, " ")
	for _, field := range fields {
		parsed = append(parsed, strings.Split(field, ":"))
	}
	return parsed
}
