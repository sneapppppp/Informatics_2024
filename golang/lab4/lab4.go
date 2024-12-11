package Lab4

import (
	"fmt"
	"math"
)

func Calculate(x, a float64) (float64, error) {
	if x <= 1 {
		return 0, fmt.Errorf(x)
	}

	y := math.Pow(a, math.Sqrt(x)-1) - math.Log10(math.Sqrt(x)-1) + math.Pow(math.Sqrt(x)-1, 1.0/3.0)
	return y, nil
}

func TaskA(b, Xn, Xk, delX float64) []float64 {
	var y []float64
	for x := Xn; x <= Xk; x += delX {
		result, err := Calculate(x, b)
		if err != nil {
			fmt.Printf(x, err)
			continue
		}
		y = append(y, result)
	}
	return y
}

func TaskB(b float64, x []float64) []float64 {
	var y []float64
	for _, value := range x {
		result, err := Calculate(value, b)
		if err != nil {
			fmt.Printf(, value, err)
			continue
		}
		y = append(y, result)
	}
	return y
}

func RunLab4() {
	b := 2.5
	fmt.Println(TaskA(b, 1.2, 3.7, 0.5))

	var s = []float64{1.28, 1.36, 2.47, 3.68, 4.56}
	fmt.Println(TaskB(b, s))
}
