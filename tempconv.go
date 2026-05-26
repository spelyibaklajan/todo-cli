package main

import "fmt"

func celsiusToFahrenheit(c float64) float64 {
	return c*9/5 + 32
}

func fahrenheitToCelsius(f float64) float64 {
	return (f - 32) * 5 / 9
}

func main() {
	temps := []float64{0, 100, 37, -40}

	fmt.Println("Celsius → Fahrenheit:")
	for _, c := range temps {
		fmt.Printf("%.1f°C = %.1f°F\n", c, celsiusToFahrenheit(c))
	}

	fmt.Println("\nFahrenheit → Celsius:")
	for _, f := range temps {
		fmt.Printf("%.1f°F = %.1f°C\n", f, fahrenheitToCelsius(f))
	}
}