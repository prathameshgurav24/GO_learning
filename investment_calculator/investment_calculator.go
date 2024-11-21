package main

import (
	"fmt"
	"math"
)

const inflationRate = 2.5 //can not be changed
func main() {

	var investMentAmount float64
	var expectedReturnRate float64 = 5.5
	var years float64

	fmt.Print("Investment Amount : ")
	fmt.Scan(&investMentAmount)

	fmt.Print("Investment Years : ")
	fmt.Scan(&years)

	futureValue := calculateFutureValue(investMentAmount, expectedReturnRate, years)
	futureRealValue := calculateRealFutureValue(futureValue, years)

	formattedFV := fmt.Sprintf("futureValue : %f\n", futureValue)
	formattedRFV := fmt.Sprintf("futureRealValue : %f\n", futureRealValue)

	fmt.Print(formattedFV, formattedRFV)

	//`` allow line breaks in print statments
	// fmt.Printf(`futureValue : %f\n
	// futureRealValue : %f\n`, futureValue, futureRealValue)
	futureValueNew, futureRealValueNew := calculateFVandRVR(investMentAmount, expectedReturnRate, years)

	fmt.Printf(`futureValueNew : %f
	futureRealValueNew : %f`, futureValueNew, futureRealValueNew)
}

func calculateFutureValue(investMentAmount, expectedReturnRate, years float64) float64 {
	return investMentAmount * math.Pow(1+expectedReturnRate/100, years)
}

func calculateRealFutureValue(futureValue, years float64) float64 {
	return futureValue / math.Pow(1+inflationRate/100, years)
}

// func calculateFVandRVR(investMentAmount, expectedReturnRate, years float64) (float64, float64) {
// 	futureValue := investMentAmount * math.Pow(1+expectedReturnRate/100, years)
// 	formattedRFV := futureValue / math.Pow(1+inflationRate/100, years)
// 	return futureValue, formattedRFV
// }

func calculateFVandRVR(investMentAmount, expectedReturnRate, years float64) (futureValueNew float64, futureRealValueNew float64) {
	futureValueNew = investMentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValueNew = futureValueNew / math.Pow(1+inflationRate/100, years)
	return
}
