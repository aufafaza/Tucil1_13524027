package main

import (
	"fmt"
	"log"
	"time"

	"github.com/aufafaza/tucil1-stima.git/src/gui"
	"github.com/aufafaza/tucil1-stima.git/src/models"
	"github.com/aufafaza/tucil1-stima.git/src/solver"
	"github.com/aufafaza/tucil1-stima.git/src/utils"
)

func main() {
	var fileName string
	fmt.Print("Enter the input file path (e.g., test/files/test1.txt): ")
	fmt.Scanln(&fileName)

	grid, err := utils.ReadFile(fileName)
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}

	fmt.Println("solve using pure brute force (with GUI) or with heuristic (more fun)?")
	fmt.Println("1 for brute force (+GUI), 2 for heuristic")
	var input int
	fmt.Scanln(&input)
	if input == 1 {
		gui.StartGame(grid, fileName)
	} else {
		board := &models.Board{
			Grid: grid,
			Size: len(grid),
			Q:    make([]int, len(grid)),
		}
		start := time.Now()
		success := solver.Solver(board, fileName)
		end := time.Since(start)
		fmt.Printf("algorithm ran for %s", end)
		if success {
			fmt.Println("no solution")
		} else {
			fmt.Println("found solution")
		}
	}
}
