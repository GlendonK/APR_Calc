package main

import (
	"fmt"
)

/**
change compounder(_depositPrice float64, _apr float64, _gasTotal float64, _days float64)
to your own price
*/

func main() {
	var compoundDaysList [365]float64
	for i := 1; i <= 365; i++ {
		compoundDaysList[i-1] = compounder(527.98, 147.70, 2, float64(i))
	}
	var days, max = findMax(compoundDaysList)
	var apy = ((max- 527.98)/527.98) *100
	fmt.Printf("Optimal days: %d of %f%% apy at $%f \n", days, apy, max)
}

func findMax(list [365]float64) (int, float64) {
	var max float64 = list[0]
	var days int = 0
	for i := 0; i < 365; i++ {
		if list[i] > max {
			max = list[i]
			days = i + 1
		}
	}
	return days, max
}

func compounder(_depositPrice float64, _apr float64, _gasTotal float64, _days float64) float64 {

	var accumulatedYield float64 = 0
	for i := 0; i < 365; {
		i++
		var yearlyYield float64 = _depositPrice * (_apr / 100)
		var dailyYield float64 = yearlyYield / 365
		accumulatedYield = accumulatedYield + dailyYield

		if i%int(_days) == 0 && i != 0 {
			_depositPrice = compoundAndDeposit(_depositPrice, accumulatedYield, _gasTotal)
			accumulatedYield = 0
		}

		if i == 365 {
			_depositPrice = _depositPrice + accumulatedYield
		}
	}

	fmt.Printf("Compounded %f days: %f\n", _days, _depositPrice)
	return _depositPrice

}

func compoundAndDeposit(_depositPrice float64, accumulatedYield float64, _gasTotal float64) float64 {
	return _depositPrice + accumulatedYield - _gasTotal
}
