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

func printUsage(commandName string) {
	fmt.Printf("Usage: %s [array dimension length] [first array file]\n", commandName)
}

func main() {
	if len(os.Args) != 3 {
		printUsage(os.Args[0])
		os.Exit(1)
	}
	arrayWidth, convError := strconv.Atoi(os.Args[1])
	fileName := os.Args[2]

	if convError != nil {
		log.Fatal(convError)
	}

	file, err := os.Open(fileName)

	if err != nil {
		panic(err)
	}

	firstArray := ioutil.ReadDistancesArray(arrayWidth, *bufio.NewScanner(os.Stdin))
	secondArray := ioutil.ReadDistancesArray(arrayWidth, *bufio.NewScanner(file))

	defer func() {
		if err := file.Close(); err != nil {
			panic(err)
		}
	}()

	differences := make([][]int, arrayWidth)
	for i := range differences {
		differences[i] = make([]int, arrayWidth)
	}

	amountOfUknown := 0
	var sumOfSquareDifferences int64 = 0
	for i := 0; i < arrayWidth; i++ {
		for j := 0; j < arrayWidth; j++ {
			firstNumber := firstArray[i][j]
			secondNumber := secondArray[i][j]
			if firstNumber == -1 || secondNumber == -1 {
				differences[i][j] = -1
				amountOfUknown++
			} else {
				difference := int(math.Abs(float64(firstNumber - secondNumber)))
				differences[i][j] = difference
				sumOfSquareDifferences += int64(difference * difference)
			}
		}
	}

	numbersInArray := arrayWidth * arrayWidth
	fmt.Printf("couldn't guess %d numbers of %d, ratio is: %f\n", amountOfUknown, numbersInArray, float64(amountOfUknown)/float64(numbersInArray))
	fmt.Printf("difference is: %0.2f\n", math.Sqrt(float64(sumOfSquareDifferences)))
	ioutil.PrintDistancesArray(differences)
}
