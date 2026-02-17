package solver

import (
	"math"

	"github.com/aufafaza/tucil1-stima.git/src/models"
)

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
