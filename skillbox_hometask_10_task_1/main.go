package main

import (
	"fmt"
	"math"
)

func main() {
	var x, epsilon float64
	var number int
	fmt.Println("Вставление eˆx в ряд Тейлора")
	fmt.Printf("Просим ввести x: ")
	fmt.Scan(&x)
	fmt.Printf("Просим ввести степень точности расчета: ")
	fmt.Scan(&number)
	epsilon = 1 / math.Pow10(number)
	var result, prevResult = 0.0, 0.0
	factorial := 1
	quantity := 0
	for {
		if quantity > 1 {
			factorial *= quantity
		}
		result += math.Pow(x, float64(quantity)) / float64(factorial)
		if math.Abs(result-prevResult) < epsilon {
			fmt.Println("")
			fmt.Println("Результат:", result)
			break
		}
		quantity++
		prevResult = result
	}
}