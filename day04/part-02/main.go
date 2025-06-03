package main

import (
	"bufio"
	"fmt"
	"os"
	"util"
)

func main() {
	file, err := os.Open("../input.txt")
	util.Check(err)
	defer file.Close()

	var grid []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		grid = append(grid, scanner.Text())
	}
	err = scanner.Err()
	util.Check(err)

	result := 0

	findWord := func(row, col int) {
		upperLeft := string(grid[row-1][col-1])
		upperRight := string(grid[row-1][col+1])
		lowerLeft := string(grid[row+1][col-1])
		lowerRight := string(grid[row+1][col+1])

		if (upperLeft == "M" && lowerRight == "S" || upperLeft == "S" && lowerRight == "M") && (upperRight == "M" && lowerLeft == "S" || upperRight == "S" && lowerLeft == "M") {
			result++
		}
	}

	for r := 1; r < len(grid)-1; r++ {
		for c := 1; c < len(grid[r])-1; c++ {
			if string(grid[r][c]) == "A" {
				findWord(r, c)
			}
		}
	}

	fmt.Println(result)
}
