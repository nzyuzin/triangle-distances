package ioutil

import (
	"bufio"
	"fmt"
	"log"
	"strconv"
)

func ReadDistancesArray(size int, input bufio.Scanner) (result [][]int) {
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

func PrintDistancesArray(distancesArray [][]int) {
	for i := 0; i < len(distancesArray); i++ {
		for j := 0; j < len(distancesArray); j++ {
			processedDistance := distancesArray[i][j]
			if processedDistance != -1 {
				fmt.Printf("%d", processedDistance)
			} else {
				fmt.Printf("?")
			}
			if j == len(distancesArray)-1 {
				fmt.Printf("\n")
			} else {
				fmt.Printf("\t")
			}
		}
	}
}
