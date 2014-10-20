package main

import (
	"bufio"
	"fmt"
	"ioutil"
	"log"
	"os"
	"strconv"
)

func printUsage(commandName string) {
	fmt.Printf("Usage: %s [array dimension length]\n, [subtrahend]", commandName)
}

func main() {
	if len(os.Args) != 3 {
		printUsage(os.Args[0])
		os.Exit(1)
	}
	arrayDimension, convError := strconv.Atoi(os.Args[1])
	if convError != nil {
		log.Fatal(convError)
	}
	subtrahend, convError := strconv.Atoi(os.Args[2])
	if convError != nil {
		log.Fatal(convError)
	}

	in := bufio.NewScanner(os.Stdin)

	inputNumbers := ioutil.ReadDistancesArray(arrayDimension, *in)
	for i := 0; i < arrayDimension; i++ {
		for j := 0; j < arrayDimension; j++ {
			inputNumbers[i][j] = subtrahend - inputNumbers[i][j]
		}
	}
	ioutil.PrintDistancesArray(inputNumbers)
}
