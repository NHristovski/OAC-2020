package main

import (
	"bufio"
	"fmt"
	"os"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func countTrees(incrementX, incrementY int) int {
	const input_file = "day3/day3_input"

	file, err := os.Open(input_file)
	check(err)

	// close the file after the function ends
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var matrix [][]bool

	for scanner.Scan() {
		line := scanner.Text()

		var row []bool

		for _, char := range line {
			row = append(row, string(char) == "#")
		}

		matrix = append(matrix, row)
	}

	startingPositionX := 0
	startingPositionY := 0

	currentPositionX, currentPositionY := startingPositionX, startingPositionY

	endingPosition := len(matrix) - 1
	rowLength := len(matrix[0])

	treesCounter := 0

	for currentPositionY < endingPosition {
		currentPositionX = (currentPositionX + incrementX) % rowLength
		currentPositionY = currentPositionY + incrementY

		if matrix[currentPositionY][currentPositionX] {
			treesCounter = treesCounter + 1
		}
	}

	return treesCounter
}
func main() {
	result := countTrees(1, 1) *
		countTrees(3, 1) *
		countTrees(5, 1) *
		countTrees(7, 1) *
		countTrees(1, 2)

	fmt.Println(result)
}
