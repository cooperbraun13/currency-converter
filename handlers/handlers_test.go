package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestConvertHandlerValidConversion creates a fake request and captures the output. It then calls the handler
// directly and asserts tthe status code, headers, and body
func TestConvertHandlerValidConversion(t *testing.T) {
	// Create a request with query params
	req := httptest.NewRequest("GET", "/api/convert?amount=100&from=USD&to=EUR", nil)

	// Create a ResponseRecorder to capture the response
	rr := httptest.NewRecorder()

	// Call the handler
	ConvertHandler(rr, req)

	// Check status code
	if rr.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", rr.Code)
	}

	// Check content type
	if ct := rr.Header().Get("Content-Type"); ct != "application/json" {
		t.Errorf("Expected Content-Type application/json, got %s", ct)
	}

	// Decode and verify response body
	var resp map[string]any
	if err := json.NewDecoder(rr.Body).Decode(&resp); err != nil {
		t.Fatalf("Failed to decode response: %v", err)
	}

	if resp["from"] != "USD" || resp["to"] != "EUR" {
		t.Errorf("Unexpected currency values in response: %v", resp)
	}
}

func TestConvertMissingParameters(t *testing.T) {

}

func TestConvertInvalidAmount(t *testing.T) {

}

func TestConvertInvalidCurrency(t *testing.T) {

}

func TestRatesHandler(t *testing.T) {
	// Create a request for all the currencies and their exchange rates
	req := httptest.NewRequest("GET", "/api/rates", nil)

	// Create a ResponseRecorder to capture the response
	rr := httptest.NewRecorder()

	// Call the handler
	RatesHandler(rr, req)

	// Check status code
	if rr.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", rr.Code)
	}

	// Check content type
	if ct := rr.Header().Get("Content-Type"); ct != "application/json" {
		t.Errorf("Expected Content-Type application/json, got %s", ct)
	}

	// Decode and verify response body
	var resp map[string]any
	if err := json.NewDecoder(rr.Body).Decode(&resp); err != nil {
		t.Fatalf("Failed to decode response: %v", err)
	}

	if resp["USD"] != 1.0 || resp["EUR"] != 0.85 || resp["JPY"] != 155.84 || resp["GBP"] != 0.75 || resp["CHF"] != 0.80 {
		t.Errorf("Unexpected conversion rate values in response: %v", resp)
	}
}
