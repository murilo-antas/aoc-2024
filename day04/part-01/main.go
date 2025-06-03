package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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

	const word = "XMAS"
	result := 0
	var findWordInDirection func(int, int, int, int, string)
	findWordInDirection = func(r, c, deltaR, deltaC int, currentWord string) {
		if r < 0 || r >= len(grid) || c < 0 || c >= len(grid[r]) || len(currentWord) > len(word) {
			return
		}
		currentWord = currentWord + string(grid[r][c])
		fmt.Println(currentWord)
		if !strings.HasPrefix(word, currentWord) {
			return
		}
		if currentWord == word {
			result++
			return
		}
		findWordInDirection(r+deltaR, c+deltaC, deltaR, deltaC, currentWord)
	}

	for r := range len(grid) {
		for c := range len(grid[r]) {
			if string(grid[r][c]) == "X" {
				findWordInDirection(r, c, -1, 0, "")
				findWordInDirection(r, c, -1, 1, "")
				findWordInDirection(r, c, 0, 1, "")
				findWordInDirection(r, c, 1, 1, "")
				findWordInDirection(r, c, 1, 0, "")
				findWordInDirection(r, c, 1, -1, "")
				findWordInDirection(r, c, 0, -1, "")
				findWordInDirection(r, c, -1, -1, "")
			}
		}
	}
	fmt.Println(result)
}
