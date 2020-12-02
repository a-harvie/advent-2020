package input

import (
	"bufio"
	"os"
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
