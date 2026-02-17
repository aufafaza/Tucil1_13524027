package main

import (
	"fmt"
	"github.com/aufafaza/tucil1-stima.git/src/models"
	"github.com/aufafaza/tucil1-stima.git/src/solver"
	"github.com/aufafaza/tucil1-stima.git/src/utils"
	"log"
	"time"
)

func main() {
	grid, err := utils.ReadFile("../test/test1.txt")
	if err != nil {
		log.Fatal(err)
	}
	N := len(grid)
	if N == 0 {
		log.Fatal("empty board")
	}
	log.Printf("Board size: %dx%d\n", N, N)
	board := &models.Board{
		Size: N,
		Grid: grid,
		Q:    make([]int, N),
		Iter: 0,
	}
	start := time.Now()
	solver.Solver2(board)
	duration := time.Since(start)
	fmt.Printf("solution found in %s\n", duration)
	fmt.Printf("solution found in %d iterations\n", board.Iter)
	fmt.Printf("number of solution: %d\n", board.SolCount)
	if board.SolCount > 0 {
		for i, sol := range board.Solutions {
			fmt.Printf("solution %d\n", i+1)
			printSolution(N, grid, sol)
		}
	} else {
		log.Println("solution not found")
	}
}

func printSolution(size int, grid [][]string, queens []int) {
	for r := 0; r < size; r++ {
		for c := 0; c < size; c++ {
			if queens[r] == c {
				fmt.Print("Q ") // Print Queen
			} else {
				fmt.Printf("%s ", grid[r][c])
			}
		}
		fmt.Println()
	}
}
