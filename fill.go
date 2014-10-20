package main

import (
	"bufio"
	"flag"
	"ioutil"
	"log"
	"math"
	"os"
)

const DEBUG bool = false

var triangular bool

func main() {
	var arraySize int
	flag.IntVar(&arraySize, "s", -1, "Array dimension length")
	flag.BoolVar(&triangular, "t", false, `Specifies if input array should be
	treated like upper-triangular matrix`)
	flag.Parse()

	if arraySize < 0 {
		flag.Usage()
		os.Exit(1)
	}

	in := bufio.NewScanner(os.Stdin)

	inputNumbers := ioutil.ReadDistancesArray(arraySize, *in)
	fillingResult := fillMissingDistances(inputNumbers)
	ioutil.PrintDistancesArray(fillingResult)
}

func fillMissingDistances(distancesArray [][]int) (result [][]int) {
	result = make([][]int, len(distancesArray))
	for i := range result {
		result[i] = make([]int, len(distancesArray))
	}

	for i := range distancesArray {
		var initJ int
		if triangular {
			initJ = i
		} else {
			initJ = 0
		}
		for j := initJ; j < len(distancesArray); j++ {
			var newDistance int
			if distancesArray[i][j] == -1 {
				newDistance = calculateMissingDistance(distancesArray, i, j)
			} else {
				newDistance = distancesArray[i][j]
			}
			result[i][j] = newDistance
			if triangular {
				result[j][i] = newDistance
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
			if j == i { // TODO: clever data partitioning is needed for this type of guessing
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
