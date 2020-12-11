package input

import (
	"bufio"
	"os"
	"strconv"
)

// ReadInputFileLines will read an input file and return a string slice containing the lines from the file
func ReadInputFileLines(filePath string) ([]string, error) {
	inputLines := make([]string, 0)

	file, err := os.Open(filePath)
	if err != nil {
		return inputLines, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		inputLines = append(inputLines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return inputLines, err
	}

	return inputLines, nil
}

// ReadInputInts will read lines from a file and give you ints, probably
func ReadInputInts(filePath string) ([]int, error) {
	input, err := ReadInputFileLines(filePath)
	if err != nil {
		return []int{}, err
	}
	output := make([]int, len(input))
	for i, line := range input {
		integer, err := strconv.Atoi(line)
		if err != nil {
			return []int{}, err
		}
		output[i] = integer
	}
	return output, nil
}
