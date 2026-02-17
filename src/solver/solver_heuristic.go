package solver

import (
	"fmt"
	"github.com/aufafaza/tucil1-stima.git/src/models"
	"github.com/aufafaza/tucil1-stima.git/src/utils"
	"math"
	"os"
	"path/filepath"
	"strings"
)

func ColorList(board [][]string) map[string]bool {
	// create a map of colors in the data
	colors := make(map[string]bool)
	for _, row := range board {
		for _, color := range row {
			colors[color] = true
		}
	}
	return colors
}

// params
// grid : the given board read by utils.ReadFile
// queens : a map showing the column location of a queen in a given row
// row, col: the given row and  column in the recursion step

func IsSafe(grid [][]string, queens []int, row, col int) bool {
	//
	targetColor := grid[row][col]

	for i := 0; i < row; i++ {
		tempCol := queens[i]

		// if current col = existing, return false
		if tempCol == col {
			return false
		}

		// check colors map, if the current color of the grid is the same as the one in the queens list, false.
		if grid[i][tempCol] == targetColor {
			return false
		}

		// check adjacency
		// given its a check for adjacency, just check for rows that is just 1 distance away from the current check
		rowDiff := int(math.Abs(float64(row - i)))
		colDiff := int(math.Abs(float64(col - tempCol)))

		if rowDiff <= 1 && colDiff <= 1 {
			return false
		}
	}
	return true
}

func Solver(board *models.Board, filePath string) bool {
	if err := ValidateUnsolvable(board); err != nil {
		fmt.Printf("UNSOLVABLE: %v\n", err)
		return false
	}
	row := 0

	for i := range board.Q {
		board.Q[i] = -1
	}

	for row >= 0 {

		// check if sollution reach, i.e. reached end of row
		if row == board.Size {
			// copy solution to the struct
			board.SolCount = 1
			solCopy := make([]int, board.Size)
			copy(solCopy, board.Q)
			board.Solutions = append(board.Solutions, solCopy)

			baseName := filepath.Base(filePath)
			nameOnly := strings.TrimSuffix(baseName, filepath.Ext(baseName))
			outputName := "solution_" + nameOnly + ".txt"

			outputDir := "output_alternative"
			os.MkdirAll(outputDir, 0755)
			finalPath := filepath.Join(outputDir, outputName)

			utils.WriteFile(finalPath, board)
			return true
		}

		startCol := board.Q[row] + 1

		found := false

		for col := startCol; col < board.Size; col++ {
			board.Iter++
			if IsSafe(board.Grid, board.Q, row, col) {
				board.Q[row] = col
				found = true
				break
			}
		}

		if found {
			row++

			if row < board.Size {
				board.Q[row] = -1
			}
		} else {
			board.Q[row] = -1
			row--
		}
	}

	return false
}

// derived from the update
func ValidateUnsolvable(board *models.Board) error {
	n := board.Size

	if len(board.Grid) != n {
		return fmt.Errorf("grid height (%d) does not match board size (%d)", len(board.Grid), n)
	}
	for i, row := range board.Grid {
		if len(row) != n {
			return fmt.Errorf("row %d has length %d, expected %d", i, len(row), n)
		}
	}

	colorMap := make(map[string]bool)
	for _, row := range board.Grid {
		for _, cell := range row {
			if cell != "" {
				colorMap[cell] = true
			}
		}
	}
	if len(colorMap) < n {
		return fmt.Errorf("only %d colors found, but %d are required", len(colorMap), n)
	}

	return nil
}
