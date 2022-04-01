package main

import (
	"fmt"
	"math"
)

func GoldenRatioRaisedToNthPowerOf2(n int) map[string]float64 {
	// Returns golden ratio (Φ) Raised to nth power of 2

	// When Golden Ratio is Raised to the power of 2
	var PreviousCoefficientOfGoldenRatio float64 = 1
	var PreviousConstant float64 = 1


	// Computing Golden Ratio Raised to the power of nth power of 2
	for i:=0; i<(n-1); i++{
		CoefficientOfGoldenRatio := math.Pow(PreviousCoefficientOfGoldenRatio, 2) + (2 * PreviousCoefficientOfGoldenRatio * PreviousConstant)
		Constant := math.Pow(PreviousCoefficientOfGoldenRatio, 2) + math.Pow(PreviousConstant, 2)

		PreviousCoefficientOfGoldenRatio = CoefficientOfGoldenRatio
		PreviousConstant = Constant
	}

	var result = map[string]float64{
		"CoefficientOfGoldenRatio": PreviousCoefficientOfGoldenRatio,
		"Constant": PreviousConstant,
	}

	return result
}


func main() {
	const NumberOfIterations = 8

	for n:=1; n<NumberOfIterations; n++{
		Result := GoldenRatioRaisedToNthPowerOf2(n)
		fmt.Printf("Φ^%v = %vΦ + %v\n",math.Pow(2.0, float64(n)), Result["CoefficientOfGoldenRatio"], Result["Constant"])
	}
}