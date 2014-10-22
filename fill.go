package main

import (
	"bufio"
	"flag"
	"ioutil"
	"log"
	"math"
	"os"
)

var DEBUG bool

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
	flag.BoolVar(&DEBUG, "d", false, "Enables debug logging")
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
			distancesArray[i][j].value = newDistance.value
			distancesArray[i][j].levelOfTrust = levelOfTrust
			if triangular {
				result[j][i] = newDistance.value
				distancesArray[j][i].value = newDistance.value
				distancesArray[j][i].levelOfTrust = levelOfTrust
			}
		}
	}

	return
}

func calculateMissingDistance(distances [][]Distance, row int, col int) Distance {
	if row == col {
		return Distance{0, 0} // We assume that distance between same objects is 0
	}

	if DEBUG {
		log.Printf("Calculating distance for [%d, %d]", row, col)
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
	if len(distances) == 0 {
		if DEBUG {
			log.Printf("No different distances found")
		}
		return Distance{-1, -1}
	}

	if DEBUG {
		log.Printf("Known distances: %v", distances)
	}
	equivalentDistances := make([]int, len(distances))
	for i := range equivalentDistances {
		equivalentDistances[i] = i
	}

	for i := 0; i < len(distances)-1; i++ {
		for j := i + 1; j < len(distances); j++ {
			if !distances[i].differentFrom(distances[j]) {
				equivalentDistances[j] = equivalentDistances[i]
			}
		}
	}
	return getMostCommonDistance(distances, equivalentDistances)
}

func getMostCommonDistance(distances []Distance, equivalentDistances []int) Distance {
	amountOfEquivalent := make([]int, len(distances))
	for i := range amountOfEquivalent {
		amountOfEquivalent[i] = 0
	}

	for i := range equivalentDistances {
		amountOfEquivalent[equivalentDistances[i]] += 1
	}

	max := 0
	for i := range amountOfEquivalent {
		if amountOfEquivalent[i] > amountOfEquivalent[max] {
			max = i
		}
	}

	result := 0
	for i := range equivalentDistances {
		if equivalentDistances[i] == max {
			result += distances[i].value / amountOfEquivalent[max]
		}
	}

	if DEBUG {
		log.Printf("amountOfEquivalent = %v, equivalentDistances = %v, result = %d",
			amountOfEquivalent, equivalentDistances, result)
	}

	return Distance{result, distances[max].levelOfTrust}
}
