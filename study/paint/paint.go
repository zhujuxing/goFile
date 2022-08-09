package main

import "fmt"

func paintNeeded(width float64, height float64) float64 {
	area := width * height
	return float64(area / 10.0)
}

func main() {
	var amount, total float64
	amount = paintNeeded(4.2, 3)
	fmt.Printf("%0.2f liters is needed.\n", amount)
	total += amount
	amount = paintNeeded(5.2, 3.5)
	fmt.Printf("%0.2f liters is needed.\n", amount)
	total += amount
	fmt.Printf("Total: %0.2f liters is needed.\n", total)
}
