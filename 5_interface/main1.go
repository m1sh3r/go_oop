package main

import (
	"fmt"
	"math"
)

func calculateY(alpha float64) float64 {
	return math.Cos(alpha) + math.Sin(alpha) + math.Cos(3*alpha) + math.Sin(3*alpha)
}

func calculateZ(alpha float64) float64 {
	return 2 * math.Sqrt(2) * math.Cos(alpha) * math.Sin(math.Pi/4+2*alpha)
}

func main() {
	var alpha float64
	fmt.Print("Введите значение alpha в радианах: ")
	_, err := fmt.Scanln(&alpha)
	if err != nil {
		fmt.Println("Ошибка ввода:", err)
		return
	}

	y := calculateY(alpha)
	z := calculateZ(alpha)
	diff := y - z

	fmt.Printf("y = %f\n", y)
	fmt.Printf("z = %f\n", z)
	fmt.Printf("Разность (y - z) = %g\n", diff)
}
