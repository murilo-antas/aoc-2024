package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"util"
)

func main() {
	leftLocationIdList, rightLocationIdList, err := readLocationIds("../input.txt")
	util.Check(err)
	rightLocationCount := make(map[int]int)
	for i := 0; i < len(rightLocationIdList); i++ {
		rightLocationCount[rightLocationIdList[i]] += 1
	}

	similarityScore := 0
	for i := 0; i < len(leftLocationIdList); i++ {
		similarityScore += leftLocationIdList[i] * rightLocationCount[leftLocationIdList[i]]
	}

	fmt.Println("Similarity score:", similarityScore)
}

func readLocationIds(fileName string) ([]int, []int, error) {
	file, err := os.Open(fileName)
	util.Check(err)
	defer file.Close()

	var leftList []int
	var rightList []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		numbers := strings.Split(line, "   ")
		leftId, err := strconv.Atoi(numbers[0])
		util.Check(err)
		rightId, err := strconv.Atoi(numbers[1])
		util.Check(err)
		leftList = append(leftList, leftId)
		rightList = append(rightList, rightId)
	}
	return leftList, rightList, scanner.Err()
}
