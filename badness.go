package main

import (
	"bufio"
	"flag"
	"fmt"
	"ioutil"
	"math"
	"os"
)

var DEBUG bool

func main() {
	var arrayWidth int
	var triangular bool
	flag.IntVar(&arrayWidth, "s", -1, "Array dimension length")
	flag.BoolVar(&DEBUG, "d", false, "Enables debug logging")
	flag.BoolVar(&triangular, "t", false, `Specifies if input array should be
	treated like upper-triangular matrix`)
	flag.Parse()

	if arrayWidth < 0 {
		flag.Usage()
		os.Exit(1)
	}

	distancesArray := ioutil.ReadDistancesArray(arrayWidth, *bufio.NewScanner(os.Stdin))

	var averageBadness float64 = 0
	var amountOfTriangles float64 = 0

	for i := 0; i < arrayWidth; i++ {
		for j := 0; j < arrayWidth; j++ {
			if triangular && i >= j {
				continue
			}
			for k := 0; k < arrayWidth; k++ {

				if i == j || j == k || i == k { // FIXME: ugly hack
					continue
				}

				// FIXME: triangular matrix isn't considered
				firstSide := distancesArray[i][j]
				secondSide := distancesArray[i][k]
				thirdSide := distancesArray[j][k]

				a := max(firstSide, secondSide, thirdSide)
				c := min(firstSide, secondSide, thirdSide)
				b := firstSide + secondSide + thirdSide - a - c

				// FIXME: pass different badnesses as a function
				badness := (float64(a) - float64(b)) / float64(a)
				averageBadness += badness
				amountOfTriangles++
			}
		}
	}

	fmt.Printf("%f", averageBadness/amountOfTriangles)

}

// FIXME: generalize
func max(f int, s int, t int) int {
	ff := float64(f)
	fs := float64(s)
	ft := float64(t)
	return int(math.Max(math.Max(ff, fs), ft))
}

func min(f int, s int, t int) int {
	ff := float64(f)
	fs := float64(s)
	ft := float64(t)
	return int(math.Min(math.Min(ff, fs), ft))
}
