package utils

import (
	"bufio"
	"os"
)

type InputParser struct {
}

func (inputParser *InputParser) Execute(filepath string) ([]string, error) {
	// Open the file
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, nil
}
