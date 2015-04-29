package ioutil

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"math"
	"strconv"
)

// a, b, c -- sides
// alpha, beta, gamma -- angles (radians) opposite to corresponding side
// a: alpha, b: beta, c: gamma
type Triangle struct {
	A, B, C            int
	Alpha, Beta, Gamma float64
}

func BuildTriangle(a int, b int, c int) (result Triangle, err error) {
	if a+b < c || a+c < b || b+c < a {
		return Triangle{a, b, c,
				getAngle(a, b, c), getAngle(b, a, c), getAngle(c, a, b)},
			errors.New(fmt.Sprintf("Triangle inequality doesn't hold for [%d %d %d]!", a, b, c))
	}
	return Triangle{a, b, c,
		getAngle(a, b, c), getAngle(b, a, c), getAngle(c, a, b)}, nil
}

func getAngle(toSide int, oneSide int, anotherSide int) float64 {
	return math.Acos(float64(square(oneSide)+square(anotherSide)-square(toSide)) /
		float64(2*oneSide*anotherSide))
}

func square(x int) int {
	return x * x
}

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
