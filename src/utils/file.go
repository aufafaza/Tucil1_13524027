package utils

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func ReadFile(path string) ([][]string, error) {
	file, err := os.Open(path)
	if err != nil {
		log.Printf("couldn't open file %s\n", path)
	}
	defer file.Close()

	var data [][]string

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		rowChars := strings.Split(line, "")

		if len(rowChars) > 0 {
			data = append(data, rowChars)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return data, nil

}
