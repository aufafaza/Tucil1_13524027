package solver

import (
	"math"

	"github.com/aufafaza/tucil1-stima.git/src/models"
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

func CheckValid(board *models.Board) bool {
	n := board.Size
	q := board.Q
	grid := board.Grid

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			// if columns are the same -> false
			if q[i] == q[j] {
				return false
			}

			// if same color -> false
			if grid[i][q[i]] == grid[j][q[j]] {
				return false
			}

			//check adjacency
			if math.Abs(float64(i-j)) == 1 && math.Abs(float64(q[i]-q[j])) <= 1 {
				return false
			}
		}
	}
	return true
}

func NextState(board *models.Board) bool {
	n := board.Size
	q := board.Q

	for i := n - 1; i >= 0; i-- {
		if q[i] < n-1 {
			q[i]++ // check for each column in a row
			return true
		} else {
			q[i] = 0
		}
	}
	return false
}

func Solver(board *models.Board) bool {
	row := 0

	for i := range board.Q {
		board.Q[i] = -1
	}

	for row >= 0 {

		// check if sollution reach, i.e. reached end of row
		if row == board.Size {
			// copy solution to the struct
			board.SolCount++
			solCopy := make([]int, board.Size)
			copy(solCopy, board.Q)
			board.Solutions = append(board.Solutions, solCopy)

			row--
			continue
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

func Solver2(board *models.Board) bool {
	for i := range board.Q {
		board.Q[i] = 0
	}

	for {
		board.Iter++
		if CheckValid(board) {
			board.SolCount++

			solCopy := make([]int, board.Size)
			copy(solCopy, board.Q)
			board.Solutions = append(board.Solutions, solCopy)

		}
		// have checked all states of a NxN board (N^N)
		if !NextState(board) {
			break
		}

	}
	return board.SolCount > 0

}
