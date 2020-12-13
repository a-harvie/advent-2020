package main

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"

	"github.com/a-harvie/advent-2020/input"
)

func main() {
	startTimeTest, busIDsTest := getDay13Input("./input_test")
	fmt.Println(startTimeTest, busIDsTest)
	startTimeTreal, busIDsReal := getDay13Input("./input")

	fmt.Printf("Day 13 Part 1 Test: %v\n", day13Part1Solution(startTimeTest, busIDsTest))
	fmt.Printf("Day 13 Part 1 Real: %v\n", day13Part1Solution(startTimeTreal, busIDsReal))

	fmt.Printf("Day 13 Part 2 Test: %v\n", day13Part2Solution(busIDsTest))
	fmt.Printf("Day 13 Part 2 Real: %v\n", day13Part2Solution(busIDsReal))
}

func day13Part1Solution(startTime int, busIDs []int) int {
	minTime := math.MaxInt64
	chosenBus := 0
	for _, busID := range busIDs {
		if busID < 0 {
			continue
		}
		time := startTime
		for {
			if time%busID == 0 {
				if time < minTime {
					minTime = time
					chosenBus = busID
				}
				break
			}
			time++
		}
	}
	return (minTime - startTime) * chosenBus
}

// I absolutely looked up hints for this one
// Approaching the number of buses iteratively is what helped
func day13Part2Solution(busIDs []int) int {
	time := busIDs[0]
	busLimit := 2 // start with the first two buses
	timeIterator := busIDs[0]
	for {
		thisIsItHooray := true
		for i := 0; i < busLimit; i++ {
			busID := busIDs[i]
			if busID >= 0 && (time+i)%busID != 0 {
				thisIsItHooray = false
				break
			}
		}
		if thisIsItHooray {
			if busLimit == len(busIDs) { // we've accounted for all buses
				break
			}
			// since all inputs are prime, we iterate by the product of all
			// the previously processed buses
			timeIterator = getTimeIterator(busIDs, busLimit)
			// this probably doesn't work if the last bus is an 'x'?
			// Worked for my input though, may circle back and make it less shonky
			for {
				busLimit++
				if busIDs[busLimit-1] > 0 || busLimit == len(busIDs) {
					break
				}
			}
		}
		time += timeIterator
	}
	return time
}

func getTimeIterator(busIDs []int, busLimit int) int {
	limit := 1
	for i := 0; i < busLimit; i++ {
		if busIDs[i] > 0 {
			limit *= busIDs[i]
		}
	}
	return limit
}

func getDay13Input(filePath string) (int, []int) {
	input, err := input.ReadInputFileLines(filePath)
	if err != nil {
		log.Fatal("Oh noes", err)
	}
	startTime, err := strconv.Atoi(input[0])
	if err != nil {
		log.Fatal("Oh noes", err)
	}
	split := strings.Split(input[1], ",")
	busIDs := make([]int, 0)
	for _, el := range split {
		if el == "x" {
			busIDs = append(busIDs, -1)
			continue
		}
		busID, err := strconv.Atoi(el)
		if err != nil {
			log.Fatal("Oh noes", err)
		}
		busIDs = append(busIDs, busID)
	}

	return startTime, busIDs
}
