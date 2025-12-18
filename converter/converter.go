package converter

import "fmt"

func getRates() map[string]float64 {
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

func convert(amount float64, from, to string) (float64, error) {
	// Get the rates
	rates = getRates()

	// Define what currency we are converting from and to
	fromRate, fromValid := rates[from]
	toRate, toValid := rates[to]

	// Make sure the currencies are valid (exist within our current map)
	if !fromValid || !toValid {
		return 0, fmt.Errorf("Unknown currency")
	}

	// Standardizes everything back to USD
	usd := amount / fromRate
	// Converts from USD to what currency we want
	return usd * toRate
}
