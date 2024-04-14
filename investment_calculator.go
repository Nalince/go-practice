package main

import (
	"fmt"
	"math"
)

func main() {

	const inflationRate = 2.5
	var investmentAmount float64
	var years float64
	var expectedReturnRate float64

	fmt.Print("Investment amount : ")
	fmt.Scan(&investmentAmount)
	fmt.Print("Years : ")
	fmt.Scan(&years)
	fmt.Print("Expected return rate : ")
	fmt.Scan(&expectedReturnRate)

	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	// fmt.Println("Future value : ", futureValue)
	fmt.Printf("Future value : %v\n", futureValue)
	fmt.Println(futureRealValue)
	outputText("dsdsds")
}

func outputText(text string) {
	fmt.Println(text)
}
