package main

import (
	"fmt"
	"math"
	"strconv"
)

const input = "23"

func main() {
	n, err := strconv.Atoi(input)
	if err != nil {
		panic("error converting string")
	}

	_ = getCornerDistance(n)
}

func getCornerDistance(n int) int {
	nextCorner := 1
	currentCorner := 0
	spiralCount := 0
	rowSize := 0

	for i := 1; i <= n; i++ {
		rt := math.Sqrt(float64(i))
		spiralCount++
		if rt == float64(nextCorner) {
			currentCorner++
			nextCorner += 2
			spiralCount = 0
			rowSize += 8
		}
	}

	fmt.Println(currentCorner, rowSize%spiralCount)

	return currentCorner
}
