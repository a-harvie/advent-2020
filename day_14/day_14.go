package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"

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
	inputTest := getDay14Input("./input_test", false)
	fmt.Println(inputTest)
	inputReal := getDay14Input("./input", false)

	fmt.Printf("Day 14 Part 1 Test: %v\n", day14Part1Solution(inputTest))
	fmt.Printf("Day 14 Part 1 Real: %v\n", day14Part1Solution(inputReal))

	inputTest = getDay14Input("./input_test_2", true)
	fmt.Println(inputTest)
	fmt.Printf("Day 14 Part 2 Test: %v\n", day14Part2Solution(inputTest))
	fmt.Printf("Day 14 Part 2 Real: %v\n", day14Part2Solution(inputReal))
}

func day14Part1Solution(cmds []cmd) uint64 {
	memory := make(map[uint64]uint64)
	currMask := mask{}
	for _, c := range cmds {
		switch c.op {
		case "mas":
			currMask = c.mask
		case "mem":
			v := c.val
			v = v | currMask.one
			v = v &^ currMask.zero
			memory[c.addr] = v
		}
	}
	total := uint64(0)
	for _, v := range memory {
		total += v
	}
	return total
}

func day14Part2Solution(cmds []cmd) uint64 {
	memory := make(map[string]uint64)
	currMask := mask{}
	for _, c := range cmds {
		switch c.op {
		case "mas":
			currMask = c.mask
		case "mem":
			addr := fmt.Sprintf("%036s", strconv.FormatInt(int64(c.addr), 2))
			masked := applyStringMask(currMask.str, addr)
			addrs := generateFloatingAddresses(masked)
			for _, a := range addrs {
				memory[a] = c.val
			}
		}
	}
	total := uint64(0)
	for _, v := range memory {
		total += v
	}
	return total
}

func getDay14Input(filePath string, getFloating bool) []cmd {
	input, err := input.ReadInputFileLines(filePath)
	if err != nil {
		log.Fatal("Oh noes", err)
	}
	cmds := make([]cmd, len(input))
	memr := regexp.MustCompile("mem\\[(\\d+)\\] = (\\d+)")
	masr := regexp.MustCompile("mask = ([01X]+)")
	for i, line := range input {
		var c cmd
		switch prefix := line[0:3]; prefix {
		case "mas":
			elements := masr.FindStringSubmatch(line)
			zero, err := makeZeroMask(elements[1])
			if err != nil {
				log.Fatal("Oh noes", err)
			}
			one, err := makeOneMask(elements[1])
			if err != nil {
				log.Fatal("Oh noes", err)
			}
			c = cmd{
				op: "mas",
				mask: mask{
					zero: zero,
					one:  one,
					str:  elements[1],
				},
			}
		case "mem":
			elements := memr.FindStringSubmatch(line)
			addr, err := strconv.ParseUint(elements[1], 10, 64)
			if err != nil {
				log.Fatal("Oh noes", err)
			}
			val, err := strconv.ParseUint(elements[2], 10, 64)
			if err != nil {
				log.Fatal("Oh noes", err)
			}
			c = cmd{
				op:   "mem",
				addr: addr,
				val:  val,
			}
		}
		cmds[i] = c
	}
	return cmds
}

func makeZeroMask(s string) (uint64, error) {
	m := strings.ReplaceAll(s, "1", "X")
	m = strings.ReplaceAll(m, "0", "1")
	m = strings.ReplaceAll(m, "X", "0")
	b, err := strconv.ParseUint(m, 2, 64)
	if err != nil {
		return 0, err
	}
	return b, nil
}

func makeOneMask(s string) (uint64, error) {
	m := strings.ReplaceAll(s, "X", "0")
	b, err := strconv.ParseUint(m, 2, 64)
	if err != nil {
		return 0, err
	}
	return b, nil
}

func applyStringMask(mask, target string) string {
	result := ""
	for i := len(target); i > 0; i-- {
		switch mask[i-1 : i] {
		case "0":
			result = target[i-1:i] + result
		default:
			result = mask[i-1:i] + result
		}
	}
	return result
}

func generateFloatingAddresses(masked string) []string {
	return generateFloating(masked, []string{""}, len(masked))
}

func generateFloating(s string, f []string, i int) []string {
	if i == 0 {
		return f
	}
	switch s[i-1 : i] {
	case "0", "1":
		for j, str := range f {
			f[j] = s[i-1:i] + str
		}
	case "X":
		f2 := make([]string, 0)
		for _, str := range f {
			f2 = append(f2, "0"+str, "1"+str)
		}
		f = f2
	}
	return generateFloating(s, f, i-1)
}
