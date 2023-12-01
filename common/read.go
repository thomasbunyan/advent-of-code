package common

import (
	"bufio"
	"io"
	"os"
)

func Read(input string) ([]string, error) {
	f, err := os.Open(input)
	if err != nil {
		return nil, err
	}

	reader := bufio.NewReader(f)

	var output []string

	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}

		if err != nil {
			return nil, err
		}

		output = append(output, string(line))
	}

	return output, nil
}

func ReadFast(input string) []string {
	lines, err := Read(input)
	if err != nil {
		panic(err)
	}
	return lines
}
