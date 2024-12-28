package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"util"
)

const (
	MinDiff = 1
	MaxDiff = 3
)

func main() {
	reports, err := readLevels("../input.txt")
	util.Check(err)

	safeReports := 0
	for _, report := range reports {
		fmt.Println(report)
		if checkSafeReport(report) {
			safeReports += 1
		}
	}
	fmt.Println("Safe reports:", safeReports)
}

func checkSafeReport(report []int) bool {
	lastValue := report[0]
	increasing := report[0]-report[1] >= 0
	for index := 1; index < len(report); index++ {
		diff := lastValue - report[index]
		if diff == 0 {
			return false
		}

		if increasing && (diff < MinDiff || diff > MaxDiff) {
			return false
		}

		if !increasing && (diff > -MinDiff || diff < -MaxDiff) {
			return false
		}

		lastValue = report[index]
	}
	return true
}

func readLevels(fileName string) ([][]int, error) {
	file, err := os.Open(fileName)
	util.Check(err)
	defer file.Close()

	var reports [][]int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		numbers := strings.Split(line, " ")
		var report []int
		for _, value := range numbers {
			level, err := strconv.Atoi(value)
			util.Check(err)
			report = append(report, level)
		}
		reports = append(reports, report)
	}
	return reports, scanner.Err()
}
