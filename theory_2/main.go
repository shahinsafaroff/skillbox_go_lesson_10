package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Enter A: ")
	var a float64
	fmt.Scan(&a)
	fmt.Println("Enter B: ")
	var b float64
	fmt.Scan(&b)
	fmt.Println("Enter C: ")
	var c float64
	fmt.Scan(&c)
	d := math.Pow(b, 2) -4*a*c
	if d > 0 {
		x1 := (-b + math.Sqrt(d))/(2*a)
		x2 := (-b - math.Sqrt(d))/(2*a)
		fmt.Printf("x1 = %v, x2 = %v", x1, x2)
	} else if d == 0 {
		x := (-b)/(2*a)
		fmt.Printf("x = %v\n", x)
	} else {
		realPart := -b / (2 * a)
		imaginaryPart := math.Sqrt(math.Abs(d))/(2*a)
		x1 := complex(realPart,imaginaryPart)
		x2 := complex(realPart,-imaginaryPart)
		fmt.Printf("x1 = %v, x2 = %v", x1, x2)
	}

}
