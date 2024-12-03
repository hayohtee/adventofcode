package main

import (
	"bufio"
	"cmp"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	inputFile, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("error opening input.txt")
	}
	defer inputFile.Close()

	totalSafe := 0

	inputScanner := bufio.NewScanner(inputFile)

	for inputScanner.Scan() {
		line := inputScanner.Text()
		levels := toIntSlice(strings.Split(line, " "))
		if isSorted(levels) && checkDiff(levels, 3) {
			totalSafe++
		}
	}

	fmt.Println(totalSafe)
}

func toIntSlice(value []string) []int {
	result := make([]int, len(value))
	for i, v := range value {
		intValue, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
		result[i] = intValue
	}
	return result
}

func isSorted(levels []int) bool {
	isAscending := slices.IsSortedFunc(levels, func(a, b int) int { return cmp.Compare(a, b) })
	if isAscending {
		return true
	}

	isDescending := slices.IsSortedFunc(levels, func(a, b int) int { return cmp.Compare(b, a) })
	return isDescending
}

func abs(value int) int {
	if value < 0 {
		return value * -1
	}
	return value
}

func checkDiff(levels []int, max int) bool {
	for i := 1; i < len(levels); i++ {
		diff := abs(levels[i] - levels[i-1])
		if diff == 0 || diff > max {
			return false
		}
	}

	return true
}
