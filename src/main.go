package main

import (
	"fmt"
	"github.com/aufafaza/tucil1-stima.git/src/gui"
	"github.com/aufafaza/tucil1-stima.git/src/utils"
	"log"
)

func main() {
	var fileName string
	fmt.Print("Enter the input file path (e.g., test/test1.txt): ")
	fmt.Scanln(&fileName)

	grid, err := utils.ReadFile(fileName)
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}

	gui.StartGame(grid, fileName)
}
