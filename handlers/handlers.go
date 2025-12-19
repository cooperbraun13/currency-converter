package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/cooperbraun13/currency-converter/converter"
)

// ConvertHandler takes amount, from, and to as query params, validates them, converts the amount between
// currencies, and returns the result as JSON. Returns 400 errors for invalid input
func ConvertHandler(w http.ResponseWriter, r *http.Request) {
	// Extract query parameter
	amount := r.URL.Query().Get("amount")
	from := r.URL.Query().Get("from")
	to := r.URL.Query().Get("to")

	// Parse amount to a float
	f, err := strconv.ParseFloat(amount, 64)

	// Validate input amount
	if err != nil {
		http.Error(w, "Invalid amount", http.StatusBadRequest)
		return
	}

	// Make sure two valid currencies are passed in before converting
	if from == "" || to == "" {
		http.Error(w, "At least 1 currency is missing", http.StatusBadRequest)
		return
	}

	// Convert from the given amount to the specified currency
	result, err := converter.Convert(f, from, to)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Set the Content-Type header before writing the body
	w.Header().Set("Content-Type", "application/json")

	// Set the HTTP status code
	// Writing the header before encoding allows error handling prior to writing the body
	w.WriteHeader(http.StatusOK)

	// Create structured response
	response := map[string]any{
		"from":   from,
		"to":     to,
		"amount": f,
		"result": result,
	}
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Printf("Error encoding response: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// RatesHandler returns all available currencies and their exchange rates as JSON
func RatesHandler(w http.ResponseWriter, r *http.Request) {
	// Instance of the data we want to return
	rates := converter.GetRates()

	// Set the Content-Type header before writing the body
	w.Header().Set("Content-Type", "application/json")

	// Set the HTTP status code
	// Writing the header before encoding allows error handling prior to writing the body
	w.WriteHeader(http.StatusOK)

	// Encode the data to JSON and write it directly to the http.ResponseWriter
	err := json.NewEncoder(w).Encode(rates)
	if err != nil {
		// Log the error and send a 500 Internal Server Error response if encoding fails
		log.Printf("Error encoding response: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
