package main

import (
	"fmt"
	"math"

	"example.com/investmentcalculator/fileops"
)

const inflationRate = 2.5

const calculatedFutureValuesFile = "future.txt"

/*
func getFloatFromFile(fileName string) (float64, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return 0, errors.New("Failed to find file")
	}
	valueText := string(data)
	value, err := strconv.ParseFloat(valueText, 64)
	if err != nil {
		return 0, errors.New("Failed to parse value to float")
	}
	return value, nil
}
*/

func main() {
	var investmentAmount float64
	var years float64
	expectedReturnRate := 5.5

	// fmt.Print("Investment Amount: ")
	outputText("Investment Amount: ")
	fmt.Scan(&investmentAmount) // Gets the user input

	// fmt.Print("Investment horizon: ")
	outputText("Investment horizon: ")
	fmt.Scan(&years)

	// fmt.Print("Expected return rate: ")
	outputText("Expected return rate: ")
	fmt.Scan(&expectedReturnRate)

	futureValue, futureRealValue := calculateFutureValues(investmentAmount, expectedReturnRate, years)
	// futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	// futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	formattedFV := fmt.Sprintf("Future value: %.1f\n", futureValue)
	formattedRFV := fmt.Sprintf("Future value (adjusted for inflation: %.1f\n", futureRealValue)
	// Outputs information
	// fmt.Println("Future value:", futureValue)
	// fmt.Printf(`Future value: %.1f

	// Future value (adjusted for inflation): %.1f`, futureValue, futureRealValue)
	// fmt.Println("Future value (adjusted for inflation):", futureRealValue)
	fmt.Print(formattedFV, formattedRFV)
}

func outputText(text string) {
	fmt.Print(text)

}

func calculateFutureValues(investmentAmount, expectedReturnRate, years float64) (fv float64, rfv float64) {
	// return investmentAmount * math.Pow(1+expectedReturnRate/100, years), futureValue / math.Pow(1+inflationRate/100, years)
	fv = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	rfv = fv / math.Pow(1+inflationRate/100, years)

	fileops.WriteFloatToFile(calculatedFutureValuesFile, fv, rfv)
	return fv, rfv
	// return
}
