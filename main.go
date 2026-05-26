package main

import "fmt"

func add(a, b float64) float64 { return a + b }
func sub(a, b float64) float64 { return a - b }
func mul(a, b float64) float64 { return a * b }
func div(a, b float64) float64 { 
	if b == 0 {
		fmt.Println("Error")
		return 0
	} else {
		return a / b
	}
  }

func main() {
	a := 1.0
	b := 2.0
	op := "+"

	switch op {
	case "+":
		fmt.Println(add(a, b))
	case "-":
		fmt.Println(sub(a, b))
	case "*":
		fmt.Println(mul(a, b))
	case "/":
		fmt.Println(div(a, b))
	}
}