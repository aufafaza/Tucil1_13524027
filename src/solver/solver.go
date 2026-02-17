package solver

import (
	"log"
	"math"

	"github.com/aufafaza/tucil1-stima.git/src/models"
	"time"
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

func Solver(board *models.Board) bool {
	start := time.Now()
	row := 0

	for i := range board.Q {
		board.Q[i] = -1
	}

	for row >= 0 {

		// check if sollution reach, i.e. reached end of row
		if row == board.Size {
			board.Solutions = 1
			end := time.Since(start)
			log.Printf("algorithm took %s to run\n", end)
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
	end := time.Since(start)
	log.Printf("algorithm took %s to run\n", end)

	return false
}
