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

type Distance struct {
	value        int
	levelOfTrust int
}

func (distance Distance) isKnown() bool {
	return distance.value != -1
}

func (first Distance) differentFrom(second Distance) bool {
	return math.Abs(float64(first.value-second.value)) > float64(first.value+second.value)/2*0.05
}

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
	distances := buildDistancesArray(inputNumbers)
	fillingResult := fillMissingDistances(distances)
	ioutil.PrintDistancesArray(fillingResult)
}

func buildDistancesArray(distancesArray [][]int) (result [][]Distance) {
	result = make([][]Distance, len(distancesArray))
	for i := range result {
		result[i] = make([]Distance, len(distancesArray))
	}

	for i := range distancesArray {
		for j := range distancesArray {
			result[i][j] = buildDistance(distancesArray[i][j])
		}
	}

	return
}

func buildDistance(distance int) (result Distance) {
	result.value = distance
	if distance == -1 {
		result.levelOfTrust = -1
	} else {
		result.levelOfTrust = 0
	}
	return
}

func fillMissingDistances(distancesArray [][]Distance) (result [][]int) {
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
			var newDistance Distance
			var levelOfTrust int
			if !distancesArray[i][j].isKnown() {
				newDistance = calculateMissingDistance(distancesArray, i, j)
				levelOfTrust = newDistance.levelOfTrust + 1
			} else {
				newDistance = distancesArray[i][j]
				levelOfTrust = newDistance.levelOfTrust
			}
			result[i][j] = newDistance.value
			distancesArray[i][j].levelOfTrust = levelOfTrust
			if triangular {
				result[j][i] = newDistance.value
				distancesArray[j][i].levelOfTrust = levelOfTrust
			}
		}
	}

	return
}

func calculateMissingDistance(distances [][]Distance, row int, col int) Distance {
	if DEBUG {
		log.Printf("Calculating distance for [%d, %d]", row, col)
	}

	if row == col {
		return Distance{0, 0} // We assume that distance between same objects is 0
	}

	var differentDistances []Distance

	for i := range distances { // FIXME: this loop doesn't consider that matrix can be upper-triangular
		toSourceByRow := distances[row][i]
		toAdjacentByRow := distances[col][i]
		if toSourceByRow.isKnown() && toAdjacentByRow.isKnown() && toSourceByRow.differentFrom(toAdjacentByRow) {
			differentDistances = append(differentDistances, toSourceByRow, toAdjacentByRow)
		}
		toSourceByCol := distances[i][row]
		toAdjacentByCol := distances[i][col]
		if toSourceByCol.isKnown() && toAdjacentByCol.isKnown() && toSourceByCol.differentFrom(toAdjacentByCol) {
			differentDistances = append(differentDistances, toSourceByCol, toAdjacentByCol)
		}
	}

	return findBestGuess(differentDistances)
}

func findBestGuess(distances []Distance) Distance {
	if DEBUG {
		log.Printf("Known distances: %v", distances)
	}

	for i := range distances {
		for j := range distances {
			if j == i { // TODO: clever data partitioning is needed for this type of guessing
				continue
			}
			firstDistance := distances[i]
			secondDistance := distances[j]
			if firstDistance.differentFrom(secondDistance) {
				return firstDistance
			}
		}
	}

	return Distance{-1, -1}
}
