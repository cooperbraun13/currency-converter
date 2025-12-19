package converter

import (
	"fmt"
	"math"
)

// GetRates maps currencies to their exchange rate using USD as a "base"
func GetRates() map[string]float64 {
	// Define currencies
	currencies := make(map[string]float64)

	// Assign currencies based on USD
	currencies["USD"] = 1.0    // U.S. dollar
	currencies["EUR"] = 0.85   // Euro
	currencies["JPY"] = 155.84 // Japanese yen
	currencies["GBP"] = 0.75   // Pound sterling
	currencies["CHF"] = 0.80   // Swiss franc

	return currencies
}

// Convert converts a given amount from one currency to another
func Convert(amount float64, from, to string) (float64, error) {
	// Get the rates
	rates := GetRates()

	// Define what currency we are converting from and to
	fromRate, fromValid := rates[from]
	toRate, toValid := rates[to]

	// Make sure the currencies are valid (exist within our current map)
	if !fromValid || !toValid {
		return 0, fmt.Errorf("Unknown currency")
	}

	// Standardizes everything back to USD
	usd := amount / fromRate
	// Converts from USD to what currency we want, rounded to 2 decimal places
	result := math.Round(usd*toRate*100) / 100

	// Validate amount entered
	if result < 0 {
		return 0, fmt.Errorf("Invalid amounts entered")
	}

	return result, nil
}
