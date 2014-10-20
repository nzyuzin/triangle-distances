package main

import (
	"bufio"
	"fmt"
	"ioutil"
	"log"
	"math"
	"os"
	"strconv"
)

const DEBUG bool = false

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

	inputNumbers := ioutil.ReadDistancesArray(resultDimension, *in)
	fillingResult := fillMissingDistances(inputNumbers)
	ioutil.PrintDistancesArray(fillingResult)
}

func fillMissingDistances(distancesArray [][]int) (result [][]int) {
	result = make([][]int, len(distancesArray))
	for i := range result {
		result[i] = make([]int, len(distancesArray))
	}

	for i := range distancesArray {
		for j := range distancesArray[i] {
			if distancesArray[i][j] == -1 {
				result[i][j] = calculateMissingDistance(distancesArray, i, j)
			} else {
				result[i][j] = distancesArray[i][j]
			}
		}
	}

	return
}

func calculateMissingDistance(distances [][]int, row int, col int) (result int) {
	if DEBUG {
		log.Printf("Calculating distance for [%d, %d]", row, col)
	}

	if row == col {
		return 0 // We assume that distance between same objects is 0
	}

	var differentDistances []int
	result = -1

	for i := range distances {
		toSourceByRow := distances[row][i]
		toAdjacentByRow := distances[col][i]
		if toSourceByRow != -1 && toAdjacentByRow != -1 && differentEnough(toSourceByRow, toAdjacentByRow) {
			differentDistances = append(differentDistances, toSourceByRow, toAdjacentByRow)
		}
		toSourceByCol := distances[i][row]
		toAdjacentByCol := distances[i][col]
		if toSourceByCol != -1 && toAdjacentByCol != -1 && differentEnough(toSourceByCol, toAdjacentByCol) {
			differentDistances = append(differentDistances, toSourceByCol, toAdjacentByCol)
		}
	}

	if DEBUG {
		log.Printf("Known distances: %v", differentDistances)
	}

	for i := range differentDistances {
		for j := range differentDistances {
			if j == i {
				continue
			}
			firstDistance := differentDistances[i]
			secondDistance := differentDistances[j]
			if differentEnough(firstDistance, secondDistance) {
				return firstDistance
			}
		}
	}

	if DEBUG {
		log.Printf("result is: %d", result)
	}

	return
}

func differentEnough(firstNumber int, secondNumber int) bool {
	return math.Abs(float64(firstNumber-secondNumber)) > float64(firstNumber+secondNumber)/2*0.05
}
