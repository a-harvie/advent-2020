package input

import (
	"bufio"
	"os"
	"strconv"
	"strings"
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

// ReadInputCSInts if it's comma separated ints on one line
func ReadInputCSInts(filePath string) ([]int, error) {
	input, err := ReadInputFileLines(filePath)
	if err != nil {
		return []int{}, err
	}
	split := strings.Split(input[0], ",")
	output := make([]int, len(split))
	for i, item := range split {
		integer, err := strconv.Atoi(item)
		if err != nil {
			return []int{}, err
		}
		output[i] = integer
	}
	return output, nil
}

// ReadStringChunks is a slice of slices of strings split by the splitter
func ReadStringChunks(filePath string, splitter string) ([][]string, error) {
	input, err := ReadInputFileLines(filePath)
	if err != nil {
		return [][]string{}, err
	}
	split := make([][]string, 0)
	currSlice := make([]string, 0)
	for _, str := range input {
		if str == splitter {
			split = append(split, currSlice)
			currSlice = make([]string, 0)
			continue
		}
		currSlice = append(currSlice, str)
	}
	split = append(split, currSlice)

	return split, nil
}
