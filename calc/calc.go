package calc

import "fmt"

type calculator struct {
	}

func NewCalculator () *calculator {
	return new(calculator)
}

func (c *calculator) sum (first,second float64) float64{
	result:=first+second
	return result
}

func (c *calculator) diff (first,second float64) float64{
	result:=first-second
	return result
}

func (c *calculator) mult (first,second float64) float64{
	result:=first*second
	return result
}

func (c *calculator) div (first,second float64) (result float64) {
	if second == 0 {
		fmt.Println("Делить на 0 нельзя")
	} else {
		result = first / second
		}
	return result
}

func (c *calculator) Calculate (num1, num2 float64, oper string) (result float64) {
	switch oper {
	case "+":
		result = c.sum(num1, num2)
	case "-":
		result = c.diff(num1, num2)
	case "*":
		result = c.mult(num1, num2)
	case "/":
		result = c.div(num1, num2)
	default:
		fmt.Println("Невозможно выполнить данную операцию")
	}
	return result
}