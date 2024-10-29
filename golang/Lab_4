go
package main

import (
	"fmt"
	"math"
)

func CalculateExpression(a float64, b float64, x float64) float64 {
	return (math.Cbrt(x-a) + math.Pow(x+b, 1.0/5.0)) /
		(math.Pow(x, 1.0/7.0) - math.Pow(x*x-(a+b)*(a+b), 1.0/9.0))
}

func TaskA(a float64, b float64, Xn float64, Xk float64, delX float64) []float64 {
	var results []float64
	for x := Xn; x <= Xk; x += delX {
		result := CalculateExpression(a, b, x)
		results = append(results, result)
		fmt.Printf("При x = %.2f: результат y = %.4f\n", x, result)
	}
	return results
}

func TaskB(a float64, b float64, x [5]float64) []float64 {
	var results []float64
	for _, value := range x {
		result := CalculateExpression(a, b, value)
		results = append(results, result)
		fmt.Printf("При x = %.2f: результат y = %.4f\n", value, result)
	}
	return results
}

func main() {
	a := 2.0
	b := 3.0
	Xn := 1.0
	Xk := 10.0
	delX := 1.0
	xValues := [5]float64{1, 2, 3, 4, 5}

	resultsA := TaskA(a, b, Xn, Xk, delX)
	fmt.Println("Результаты TaskA:", resultsA)

	resultsB := TaskB(a, b, xValues)
	fmt.Println("Результаты TaskB:", resultsB)
}