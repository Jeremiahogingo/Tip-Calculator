package main

import (
	"fmt"
)

func main() {
	var (
		billAmount    float64
		tipPercentage float64
	)

	fmt.Println("Enter th bill amount: $")
	fmt.Scanln(&billAmount)
	fmt.Println("Enter the tip percentage")
	fmt.Scanln(&tipPercentage)

	tipAmount := (tipPercentage / 100) * billAmount
	totalBillAmount := tipAmount + billAmount

	fmt.Printf("Tip amount: $%.2f\n", tipAmount)
	fmt.Printf("The total bill amont(Tip included): $%.2f\n", totalBillAmount)

}
