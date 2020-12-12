package main

import (
	"fmt"
	"log"
	"math"
	"regexp"
	"strconv"

	"github.com/a-harvie/advent-2020/input"
)

var (
	north = vec{0, -1}
	east  = vec{1, 0}
	south = vec{0, 1}
	west  = vec{-1, 0}
)

type nav struct {
	cmd string
	mag int
}

type ferry struct {
	position vec
	heading  vec
	waypoint vec
}

type vec struct {
	X int
	Y int
}

func main() {
	inputTest := getDay12Input("./input_test")
	inputReal := getDay12Input("./input")

	fmt.Printf("Day 12 Part 1 Test: %v\n", day12Part1Solution(inputTest))
	fmt.Printf("Day 12 Part 1 Real: %v\n", day12Part1Solution(inputReal))

	fmt.Printf("Day 12 Part 2 Test: %v\n", day12Part2Solution(inputTest))
	fmt.Printf("Day 12 Part 2 Real: %v\n", day12Part2Solution(inputReal))
}

func day12Part1Solution(input []nav) int {
	f := ferry{
		vec{0, 0},
		vec{1, 0},
		vec{10, -1},
	}
	for _, n := range input {
		switch n.cmd {
		case "N":
			f.position.translate(vec{north.X, north.Y * n.mag})
		case "E":
			f.position.translate(vec{east.X * n.mag, east.Y})
		case "S":
			f.position.translate(vec{south.X, south.Y * n.mag})
		case "W":
			f.position.translate(vec{west.X * n.mag, west.Y})
		case "L":
			f.heading.rotate(-n.mag)
		case "R":
			f.heading.rotate(n.mag)
		case "F":
			f.position.translate(vec{f.heading.X * n.mag, f.heading.Y * n.mag})
		}
	}
	return manhattan(f.position, vec{0, 0})
}

func day12Part2Solution(input []nav) int {
	f := ferry{
		vec{0, 0},
		vec{1, 0},
		vec{10, -1},
	}
	for _, n := range input {
		switch n.cmd {
		case "N":
			f.waypoint.translate(vec{north.X, north.Y * n.mag})
		case "E":
			f.waypoint.translate(vec{east.X * n.mag, east.Y})
		case "S":
			f.waypoint.translate(vec{south.X, south.Y * n.mag})
		case "W":
			f.waypoint.translate(vec{west.X * n.mag, west.Y})
		case "L":
			f.waypoint.rotate(-n.mag)
		case "R":
			f.waypoint.rotate(n.mag)
		case "F":
			f.position.translateN(vec{f.waypoint.X, f.waypoint.Y}, n.mag)
		}
	}
	return manhattan(f.position, vec{0, 0})
}

func (v *vec) translate(v2 vec) {
	v.X += v2.X
	v.Y += v2.Y
}

func (v *vec) translateN(v2 vec, n int) {
	for i := 0; i < n; i++ {
		v.translate(v2)
	}
}

func (v *vec) rotate(d int) {
	xp := v.X*cos(d) - v.Y*sin(d)
	yp := v.X*sin(d) + v.Y*cos(d)
	v.X = xp
	v.Y = yp
}

func cos(d int) int {
	r := float64(d) * math.Pi / 180.0
	return int(math.Cos(r))
}

func sin(d int) int {
	r := float64(d) * math.Pi / 180.0
	return int(math.Sin(r))
}

func manhattan(a, b vec) int {
	return abs(a.X-b.X) + abs(a.Y-b.Y)
}

func abs(a int) int {
	if a < 0 {
		return a * -1
	}
	return a
}

func getDay12Input(filePath string) []nav {
	input, err := input.ReadInputFileLines(filePath)
	if err != nil {
		log.Fatal("Oh noes", err)
	}
	r := regexp.MustCompile("([NSEWLRF])(\\d+)")
	output := make([]nav, len(input))
	for i, line := range input {
		parsed := r.FindStringSubmatch(line)
		mag, err := strconv.Atoi(parsed[2])
		if err != nil {
			log.Fatal("Oh noes", err)
		}
		output[i] = nav{
			cmd: parsed[1],
			mag: mag,
		}
	}
	return output
}
