package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"util"
)

func main() {
	memory, err := os.ReadFile("../input.txt")
	util.Check(err)
	sum := 0
	enabled := true
	pattern := regexp.MustCompile(`mul\((?P<first>\d{1,3})\,(?P<second>\d{1,3})\)|do\(\)|don't\(\)`)
	expressions := pattern.FindAllSubmatch(memory, -1)
	for _, expr := range expressions {
		if string(expr[0]) == "do()" {
			enabled = true
		} else if string(expr[0]) == "don't()" {
			enabled = false
		} else if enabled {
			term1, err := strconv.Atoi(string(expr[1]))
			util.Check(err)

			term2, err := strconv.Atoi(string(expr[2]))
			util.Check(err)
			sum += term1 * term2
		}
	}
	fmt.Println("Sum:", sum)
}
