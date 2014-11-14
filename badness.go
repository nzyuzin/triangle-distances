package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"ioutil"
	"log"
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

	badnessValue, invalidTriangles := ComputeBadness(distancesArray, arrayWidth, triangular, badness)
	fmt.Printf("%f %d\n", badnessValue, invalidTriangles)
}

func badness(triangle Triangle) float64 {
	return float64(triangle.a-triangle.b) / float64(triangle.a)
}

func ComputeBadness(distancesArray [][]int, arrayWidth int, triangular bool, getBadness func(Triangle) float64) (float64, int) {
	var averageBadness float64 = 0
	var amountOfTriangles float64 = 0
	var triangleInequalityViolations int = 0

	for i := 0; i < arrayWidth; i++ {
		for j := i + 1; j < arrayWidth; j++ {
			for k := j + 1; k < arrayWidth; k++ {

				// FIXME: non triangular matrix isn't considered
				firstSide := distancesArray[i][j]
				secondSide := distancesArray[i][k]
				thirdSide := distancesArray[j][k]

				if firstSide == -1 || secondSide == -1 || thirdSide == -1 {
					continue
				}

				a := max(firstSide, secondSide, thirdSide)
				c := min(firstSide, secondSide, thirdSide)
				b := firstSide + secondSide + thirdSide - a - c

				triangle, err := BuildTriangle(a, b, c)
				if err != nil {
					triangleInequalityViolations++
					if DEBUG {
						log.Printf("%v", err)
					}
					continue
				}
				badness := getBadness(triangle)
				averageBadness += badness
				amountOfTriangles++
			}
		}
	}
	if DEBUG {
		log.Printf("amountOfTriangles = %f", amountOfTriangles)
	}

	return averageBadness / amountOfTriangles, triangleInequalityViolations
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
