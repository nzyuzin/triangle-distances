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

func main() {
	var arrayWidth int
	flag.IntVar(&arrayWidth, "s", -1, "Array dimension length")
	flag.BoolVar(&DEBUG, "d", false, "Enables debug logging")
	flag.Parse()
	fileName := flag.Args()[0]

	if arrayWidth < 0 {
		flag.Usage()
		os.Exit(1)
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

	numbersInArray := arrayWidth * arrayWidth
	amountOfUknown := 0
	var sumOfSquareDifferences int64 = 0
	var averageDifference float64 = 0
	var differentOnlyAverageDifference float64 = 0
	var differentNumbersAmount int = 0
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
				if DEBUG {
					sumOfSquareDifferences += int64(difference * difference)
					averageDifference += float64(difference) / float64(numbersInArray)
					if difference != 0 {
						differentOnlyAverageDifference += float64(difference)
						differentNumbersAmount++
					}
				}
			}
		}
	}

	if DEBUG {
		differentOnlyAverageDifference /= float64(differentNumbersAmount)
		log.Printf("couldn't guess %d numbers out of %d, ratio is: %f\n", amountOfUknown, numbersInArray, float64(amountOfUknown)/float64(numbersInArray))
		log.Printf("different numbers amount is: %d", differentNumbersAmount)
		log.Printf("difference is: %0.2f\n", math.Sqrt(float64(sumOfSquareDifferences)))
		log.Printf("average difference is: %0.2f\n", averageDifference)
		log.Printf("different numbers only average difference is: %0.2f\n", differentOnlyAverageDifference)
	}
	ioutil.PrintDistancesArray(differences)
}
