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

func buildDistancesArray(distancesArray [][]int) (result [][]ioutil.Distance) {
	result = make([][]ioutil.Distance, len(distancesArray))
	for i := range result {
		result[i] = make([]ioutil.Distance, len(distancesArray))
	}

	for i := range distancesArray {
		for j := range distancesArray {
			result[i][j] = buildDistance(distancesArray[i][j])
		}
	}

	return
}

func buildDistance(distance int) (result ioutil.Distance) {
	result.Value = distance
	if distance == -1 {
		result.LevelOfTrust = -1
	} else {
		result.LevelOfTrust = 0
	}
	return
}

func fillMissingDistances(distancesArray [][]ioutil.Distance) (result [][]int) {
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
			if !distancesArray[i][j].IsKnown() {
				distancesArray[i][j] = calculateMissingDistance(distancesArray, i, j)
				distancesArray[i][j].LevelOfTrust += 1
			}
			result[i][j] = distancesArray[i][j].Value
			if triangular {
				result[j][i] = distancesArray[i][j].Value
				distancesArray[j][i] = distancesArray[i][j]
			}
		}
	}

	return
}

func calculateMissingDistance(distances [][]ioutil.Distance, row int, col int) (result ioutil.Distance) {
	if row == col {
		return ioutil.Distance{0, 0} // We assume that distance between same objects is 0
	}
	if DEBUG {
		log.Printf("Calculating distance for [%d, %d]", row, col)
	}
	var differentDistances []ioutil.Distance = getDifferentDistances(distances, row, col)
	if len(differentDistances) == 0 {
		if DEBUG {
			log.Printf("No different distances found")
		}
		return ioutil.Distance{-1, -1}
	}
	if DEBUG {
		log.Printf("Known distances: %v", differentDistances)
	}
	var equivalences = getEquivalentCountArray(differentDistances)
	var frequencyCount = getDistanceToFrequence(differentDistances, equivalences)
	var possibleDistance = getMostFrequentDistance(differentDistances, frequencyCount)
	result = ioutil.Distance{possibleDistance, 1}
	if DEBUG {
		log.Printf("guessed distance = %v", result)
	}
	return
}

func getDifferentDistances(distances [][]ioutil.Distance, row int, col int) (result []ioutil.Distance) {
	for i := max(row, col); i < len(distances); i++ {
		toSourceByRow := distances[row][i]
		toAdjacentByRow := distances[col][i]
		if toSourceByRow.IsKnown() && toAdjacentByRow.IsKnown() && toSourceByRow.IsDifferentFrom(toAdjacentByRow) {
			result = append(result, toSourceByRow, toAdjacentByRow)
		}
	}
	for i := 0; i < max(row, col); i++ {
		toSourceByCol := distances[i][row]
		toAdjacentByCol := distances[i][col]
		if toSourceByCol.IsKnown() && toAdjacentByCol.IsKnown() && toSourceByCol.IsDifferentFrom(toAdjacentByCol) {
			result = append(result, toSourceByCol, toAdjacentByCol)
		}
	}
	return
}

func getEquivalentCountArray(distances []ioutil.Distance) (equivalentDistances []int) {
	equivalentDistances = make([]int, len(distances))
	for i := range equivalentDistances {
		equivalentDistances[i] = i
	}
	for i := 0; i < len(distances)-1; i++ {
		for j := i + 1; j < len(distances); j++ {
			if !distances[i].IsDifferentFrom(distances[j]) {
				equivalentDistances[j] = equivalentDistances[i]
			}
		}
	}
	return
}

func getDistanceToFrequence(distances []ioutil.Distance, equivalences []int) (result map[int]int) {
	result = make(map[int]int)
	for _, v := range equivalences {
		result[distances[v].Value] += 1
	}
	return
}

func getMostFrequentDistance(distances []ioutil.Distance, distanceFrequency map[int]int) (result int) {
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
