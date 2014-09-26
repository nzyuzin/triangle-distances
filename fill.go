package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		printUsage(os.Args[0])
		os.Exit(1)
	}
	matrixDimension, convError := strconv.Atoi(os.Args[1])

	var inputWord string
	matrix := make([][]int, matrixDimension)
	for i := range matrix {
		matrix[i] = make([]int, matrixDimension)
	}
	amountOfNumbers := 0

	if convError != nil {
		log.Fatal(convError)
	}

	in := bufio.NewScanner(os.Stdin)
	in.Split(bufio.ScanWords)

	for in.Scan() {
		inputWord = in.Text()
		col := amountOfNumbers % matrixDimension
		row := amountOfNumbers / matrixDimension
		if inputWord == "?" {
			matrix[row][col] = -1
		} else {
			matrix[row][col], convError = strconv.Atoi(inputWord)
		}
		if convError != nil {
			log.Fatal(convError)
		}
		amountOfNumbers++
	}

	fmt.Printf("read %d words\n", amountOfNumbers)
}

func readInputArray() [][]int {
	return nil
}

func printUsage(commandName string) {
	fmt.Printf("Usage: %s [array dimension length]\n", commandName)
}
