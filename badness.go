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

// a, b, c -- sides
// alpha, beta, gamma -- angles (radians) opposite to corresponding side
// a: alpha, b: beta, c: gamma
type Triangle struct {
	a, b, c            int
	alpha, beta, gamma float64
}

func BuildTriangle(a int, b int, c int) Triangle {
	return Triangle{a, b, c,
		getAngle(a, b, c), getAngle(b, a, c), getAngle(c, a, b)}
}

func getAngle(toSide int, oneSide int, anotherSide int) float64 {
	return math.Acos(float64(square(oneSide)+square(anotherSide)-square(toSide)) /
		float64(2*oneSide*anotherSide))
}

func square(x int) int {
	return x * x
}

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

	fmt.Printf("%f\n", ComputeBadness(distancesArray, arrayWidth, triangular, badness))
}

func badness(triangle Triangle) float64 {
	return float64(triangle.a-triangle.b) / float64(triangle.a)
}

func ComputeBadness(distancesArray [][]int, arrayWidth int, triangular bool, getBadness func(Triangle) float64) float64 {
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

				triangle := BuildTriangle(a, b, c)
				badness := getBadness(triangle)
				averageBadness += badness
				amountOfTriangles++
			}
		}
	}

	return averageBadness / amountOfTriangles
}

func max(f int, s int, t int) int {
	intMax := floatFuncToIntFunc(math.Max)
	return intMax(intMax(f, s), t)
}

func min(f int, s int, t int) int {
	intMin := floatFuncToIntFunc(math.Min)
	return intMin(intMin(f, s), t)
}

func floatFuncToIntFunc(fn func(float64, float64) float64) func(int, int) int {
	return func(a int, b int) int {
		return int(fn(float64(a), float64(b)))
	}
}
