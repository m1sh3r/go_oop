package main

import (
	"fmt"
	"math"
)

type Calculate interface {
	Formula() float64
}

type Y struct {
	alf float64
}

type Z struct {
	alf float64
}

func (y *Y) Formula() float64 {
	return math.Cos(y.alf) + math.Sin(y.alf) + math.Cos(3*y.alf) + math.Sin(3*y.alf)
}

func (z *Z) Formula() float64 {
	return 2 * math.Sqrt(2) * math.Cos(z.alf) * math.Sin(math.Pi/4+2*z.alf)
}

func main() {
	var alpha float64
	fmt.Print("Введите значение alpha в радианах: ")
	_, err := fmt.Scanln(&alpha)
	if err != nil {
		fmt.Println("Ошибка ввода:", err)
		return
	}

	var calcY Calculate = &Y{alf: alpha}
	var calcZ Calculate = &Z{alf: alpha}

	yVal := calcY.Formula()
	zVal := calcZ.Formula()
	diff := yVal - zVal

	fmt.Printf("Y = %f\n", yVal)
	fmt.Printf("Z = %f\n", zVal)
	fmt.Printf("Разность (Y - Z) = %g\n", diff)
}
