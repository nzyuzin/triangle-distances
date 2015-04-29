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
	return math.Abs(float64(first.value-second.value)) > float64(first.value+second.value)/2*0.08
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
			if !distancesArray[i][j].isKnown() {
				distancesArray[i][j] = calculateMissingDistance(distancesArray, i, j)
				distancesArray[i][j].levelOfTrust += 1
			}
			result[i][j] = distancesArray[i][j].value
			if triangular {
				result[j][i] = distancesArray[i][j].value
				distancesArray[j][i] = distancesArray[i][j]
			}
		}
	}

	return
}

func calculateMissingDistance(distances [][]Distance, row int, col int) (result Distance) {
	if row == col {
		return Distance{0, 0} // We assume that distance between same objects is 0
	}
	if DEBUG {
		log.Printf("Calculating distance for [%d, %d]", row, col)
	}
	var differentDistances []Distance = getDifferentDistances(distances, row, col)
	if len(differentDistances) == 0 {
		if DEBUG {
			log.Printf("No different distances found")
		}
		return Distance{-1, -1}
	}
	if DEBUG {
		log.Printf("Known distances: %v", differentDistances)
	}
	var equivalences = getEquivalentCountArray(differentDistances)
	var frequencyCount = getDistanceToFrequence(differentDistances, equivalences)
	var possibleDistance = getMostCommonDistance(differentDistances, frequencyCount)
	result = Distance{possibleDistance, 1}
	if DEBUG {
		log.Printf("guessed distance = %v", result)
	}
	return
}

func getDifferentDistances(distances [][]Distance, row int, col int) (result []Distance) {
	for i := max(row, col); i < len(distances); i++ {
		toSourceByRow := distances[row][i]
		toAdjacentByRow := distances[col][i]
		if toSourceByRow.isKnown() && toAdjacentByRow.isKnown() && toSourceByRow.differentFrom(toAdjacentByRow) {
			result = append(result, toSourceByRow, toAdjacentByRow)
		}
	}
	for i := 0; i < max(row, col); i++ {
		toSourceByCol := distances[i][row]
		toAdjacentByCol := distances[i][col]
		if toSourceByCol.isKnown() && toAdjacentByCol.isKnown() && toSourceByCol.differentFrom(toAdjacentByCol) {
			result = append(result, toSourceByCol, toAdjacentByCol)
		}
	}
	return
}

func getEquivalentCountArray(distances []Distance) (equivalentDistances []int) {
	equivalentDistances = make([]int, len(distances))
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
	return
}

func getDistanceToFrequence(distances []Distance, equivalences []int) (result map[int]int) {
	result = make(map[int]int)
	for _, v := range equivalences {
		result[distances[v].value] += 1
	}
	return
}

func getMostCommonDistance(distances []Distance, distanceFrequency map[int]int) (result int) {
	maxCount := 0
	for k, v := range distanceFrequency {
		if v > maxCount {
			maxCount = v
			result = k
		}
	}
	return
}

func validTriangle(a int, b int, c int) bool {
	_, err := ioutil.BuildTriangle(a, b, c)
	if err != nil {
		log.Printf("error for %d %d %d", a, b, c)
		return false
	}
	return true
}

func max(f int, s int) int {
	return int(math.Max(float64(f), float64(s)))
}
