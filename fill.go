package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func printUsage(commandName string) {
	fmt.Printf("Usage: %s [array dimension length]\n", commandName)
}

func main() {
	if len(os.Args) != 2 {
		printUsage(os.Args[0])
		os.Exit(1)
	}
	resultDimension, convError := strconv.Atoi(os.Args[1])

	if convError != nil {
		log.Fatal(convError)
	}

	in := bufio.NewScanner(os.Stdin)

	inputNumbers := readDistancesArray(resultDimension, *in)
	fillMissingDistances(inputNumbers)
}

func readDistancesArray(size int, input bufio.Scanner) (result [][]int) {
	var inputDistance string
	result = make([][]int, size)
	for i := range result {
		result[i] = make([]int, size)
	}
	amountOfNumbers := 0
	var convError error

	input.Split(bufio.ScanWords)

	for input.Scan() {
		inputDistance = input.Text()
		col := amountOfNumbers % size
		row := amountOfNumbers / size
		if inputDistance == "?" {
			result[row][col] = -1
		} else {
			result[row][col], convError = strconv.Atoi(inputDistance)
		}
		if convError != nil {
			log.Fatal(convError)
		}
		amountOfNumbers++
	}
	log.Printf("read %d words\n", amountOfNumbers)
	return
}

func fillMissingDistances(distancesArray [][]int) (result [][]int) {
	result = make([][]int, len(distancesArray))
	for i := range result {
		result[i] = make([]int, len(distancesArray))
	}

	for i := range distancesArray {
		for j := range distancesArray[i] {
			if i == j {
				continue
			}
			if distancesArray[i][j] == -1 {
				calculateMissingDistance(distancesArray, i, j)
				return
			} else {
				result[i][j] = distancesArray[i][j]
			}
		}
	}

	return
}

func calculateMissingDistance(distancesArray [][]int, row int, col int) (result int) {
	log.Printf("Calculating distance for [%d, %d]", row, col)

	var differentDistances []int

	for i := range distancesArray {
		if i == row {
			continue // distance should always be 0
		}

		distanceToSource := distancesArray[row][i]
		distanceToAnother := distancesArray[col][i]

		if distanceToSource != -1 && distanceToAnother != -1 && distanceToSource-distanceToAnother > 200 {
			differentDistances = append(differentDistances, i)
			log.Printf("distances [%d, %d] = %d, [%d, %d] = %d", col, i, distanceToSource, row, i, distanceToAnother)
		}
	}

	log.Printf("Known distance rows = %v", differentDistances)

	return
}
