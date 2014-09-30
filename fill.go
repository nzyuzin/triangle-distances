package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
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
	fillingResult := fillMissingDistances(inputNumbers)
	for i := 0; i < len(fillingResult); i++ {
		for j := 0; j < len(fillingResult)-1; j++ {
			processedDistance := fillingResult[i][j]
			if processedDistance != -1 {
				fmt.Printf("%d\t", processedDistance)
			} else {
				fmt.Printf("?\t")
			}
		}
		fmt.Printf("%d\n", fillingResult[i][len(fillingResult)-1])
	}
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
	//log.Printf("read %d words\n", amountOfNumbers)
	return
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
	//log.Printf("Calculating distance for [%d, %d]", row, col)

	var differentDistances []int
	result = -1

	for i := range distances {
		toSource := distances[row][i]
		toAnother := distances[col][i]
		if i == row || toSource == -1 || toAnother == -1 {
			continue
		}

		differentEnough := float64(toSource)-float64(toAnother) > float64(toSource+toAnother)/2.0*0.05 // TODO: replace with function that calculates average and replace 5% (0.05) with meaningful constant

		if differentEnough {
			differentDistances = append(differentDistances, toSource, toAnother)
			//log.Printf("distances [%d, %d] = %d, [%d, %d] = %d", col, i, toSource, row, i, toAnother)
		}
	}

	for i := range differentDistances {
		// TODO: find most common range of distances, i.e. distances within found
		// range should occur most frequently in differentDistances
		for j := range differentDistances {
			if j == i {
				continue
			}
			firstDistance := differentDistances[i]
			secondDistance := differentDistances[j]
			if math.Abs(float64(firstDistance-secondDistance)) < float64(firstDistance+secondDistance)/2*0.05 {
				result = firstDistance
			}
		}
	}

	//log.Printf("Known distances: %v", differentDistances)

	//log.Printf("result is: %d", result)

	return
}
