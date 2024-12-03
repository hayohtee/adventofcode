package main

import (
	"bufio"
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

	pair1 := make([]int, 0, 1000)
	pair2 := make([]int, 0, 1000)

	inputScanner := bufio.NewScanner(inputFile)
	for inputScanner.Scan() {
		// Read values of input.txt line by line
		line := inputScanner.Text()

		// The values are separated by 3 spaces
		pairs := strings.Split(line, "   ")

		// Convert the value for each pairs into int
		value1, err := strconv.Atoi(pairs[0])
		if err != nil {
			log.Fatal(err)
		}

		value2, err := strconv.Atoi(pairs[1])
		if err != nil {
			log.Fatal(err)
		}

		pair1 = append(pair1, value1)
		pair2 = append(pair2, value2)
	}

	if len(pair1) != len(pair2) {
		panic("the pairs are not of equal size")
	}

	// Sort the pair in ascending order
	slices.Sort(pair1)
	slices.Sort(pair2)

	
	fmt.Println(calculateTotalDistance(pair1, pair2))
}

func calculateTotalDistance(pair1, pair2 []int) int  {
	totalDistance := 0

	for i := range len(pair1) {
		distance := pair1[i] - pair2[i]
		if distance < 0 {
			distance = distance * -1
		}
		totalDistance += distance
	}

	return totalDistance
}
