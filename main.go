package main

import (
	"7_7_1/calc"
	"fmt"
)

func main() {
	var (
		num1, num2 float64
		oper       string
	)

	fmt.Print("Введите 1е число:")
	fmt.Scanln(&num1)
	fmt.Print("Введите арифметический оператор:")
	fmt.Scanln(&oper)
	fmt.Print("Введите 2е число:")
	fmt.Scanln(&num2)

	calculator := calc.NewCalculator()
	result := calculator.Calculate(num1, num2, oper)

	fmt.Printf("Результат вычисления: %f", result)
}
