package utils

import (
	"bufio"
	"fmt"
	"github.com/aufafaza/tucil1-stima.git/src/models"
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

func WriteFile(filename string, board *models.Board) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	size := board.Size
	queens := board.Solutions[0]

	for r := 0; r < size; r++ {
		line := ""
		for c := 0; c < size; c++ {
			if queens[r] == c {
				line += "# "
			} else {
				line += board.Grid[r][c] + " "
			}
		}
		fmt.Fprintln(file, line)
	}
	return nil
}
