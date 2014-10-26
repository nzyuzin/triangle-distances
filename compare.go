package main

import (
	"bufio"
	"flag"
	"fmt"
	"ioutil"
	"log"
	"math"
	"os"
)

var DEBUG bool

func main() {
	var arrayWidth int
	var threshold int
	var triangular bool
	flag.IntVar(&arrayWidth, "s", -1, "Array dimension length")
	flag.BoolVar(&DEBUG, "d", false, "Enables debug logging")
	flag.BoolVar(&triangular, "t", false, `Specifies if input array should be
	treated like upper-triangular matrix`)
	flag.IntVar(&threshold, "td", -1, "Print distances which value exceeds given threshold. Negative value is ignored")
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

	sourceArray := ioutil.ReadDistancesArray(arrayWidth, *bufio.NewScanner(file))
	comparedArray := ioutil.ReadDistancesArray(arrayWidth, *bufio.NewScanner(os.Stdin))

	defer func() {
		if err := file.Close(); err != nil {
			panic(err)
		}
	}()

	differences := make([][]int, arrayWidth)
	for i := range differences {
		differences[i] = make([]int, arrayWidth)
	}

	var numbersInArray int
	if triangular {
		numbersInArray = arrayWidth * (arrayWidth - 1) / 2
	} else {
		numbersInArray = arrayWidth * arrayWidth
	}
	amountOfUknown := 0
	var sumOfSquareDifferences int64 = 0
	var averageDifference float64 = 0
	var differentOnlyAverageDifference float64 = 0
	var differentNumbersAmount int = 0
	for i := 0; i < arrayWidth; i++ {
		for j := 0; j < arrayWidth; j++ {
			if triangular && i >= j {
				continue
			}
			source := sourceArray[i][j]
			compared := comparedArray[i][j]
			if source == -1 || compared == -1 {
				differences[i][j] = -1
				amountOfUknown++
			} else {
				difference := int(math.Abs(float64(source - compared)))

				if threshold >= 0 && difference > threshold {
					fmt.Printf("[%d, %d] source = %d, compared = %d, difference = %d\n", i, j, source, compared, difference)
					continue
				}

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

	if threshold >= 0 {
		os.Exit(0)
	}

	if DEBUG {
		differentOnlyAverageDifference /= float64(differentNumbersAmount)
		log.Printf("couldn't guess %d numbers out of %d, ratio is: %f\n", amountOfUknown, numbersInArray, float64(amountOfUknown)/float64(numbersInArray))
		log.Printf("different numbers amount is: %d", differentNumbersAmount)
		log.Printf("difference is: %0.2f\n", math.Sqrt(float64(sumOfSquareDifferences/int64(numbersInArray))))
		log.Printf("average difference is: %0.2f\n", averageDifference)
		log.Printf("different numbers only average difference is: %0.2f\n", differentOnlyAverageDifference)
	}
	ioutil.PrintDistancesArray(differences)
}
