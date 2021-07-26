package main

import (
	"fmt"
	"math"
)

func main() {
	var epsilon float64
	epsilon = 0.0001
	fmt.Println("Please enter X for calculation of it's Sinus: ")
	var x float64
	_, _ = fmt.Scan(&x)
	result := 0.0
	prevResult := 0.0
	factorial := 1
	k := 0
	for {
		if k > 0 {
			factorial *= 2 * k * (2*k +1)
		}
		result += math.Pow(-1, float64(k) * math.Pow(x, float64(2*k + 1))/float64(factorial))
		if math.Abs(result - prevResult) < epsilon {
			fmt.Println(result)
			break
		}
		k++
		prevResult = result
	}
}
