package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
	"util"
)

func main() {
	leftLocationIdList, rightLocationIdList, err := readLocationIds("../input.txt")
	util.Check(err)
	slices.Sort(leftLocationIdList)
	slices.Sort(rightLocationIdList)
	sum := 0
	for i := 0; i < len(leftLocationIdList); i++ {
		sum += int(math.Abs(float64(leftLocationIdList[i] - rightLocationIdList[i])))
	}
	fmt.Println("Distance between the lists:", sum)
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
