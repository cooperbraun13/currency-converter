package converter

import "testing"

// TestConvertUSDToEur calls Convert with a specified amount and checks
// for proper conversion from USD to Euro
func TestConvertUSDToEur(t *testing.T) {
	result, err := Convert(100.0, "USD", "EUR")

	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	expected := 85.0 // 100 * 0.85
	if result != expected {
		t.Errorf("got %v, want %v", result, expected)
	}
}

// TestConvertUSDToUSD validates that converting to the same currency returns the original amount
func TestConvertUSDToUSD(t *testing.T) {
	result, err := Convert(10.0, "USD", "USD")

	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	expected := 10.0 // 10 * 1 (USD is our base)
	if result != expected {
		t.Errorf("got %v, want %v", result, expected)
	}
}

// TestConvertJPYToGBP validates our conversion of non-USD currencies between each other
func TestConvertJPYToGBP(t *testing.T) {
	result, err := Convert(1250.0, "JPY", "GBP")

	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	expected := 6.02 // 1250 JPY -> 8.02 USD -> 6.02 GBP
	if result != expected {
		t.Errorf("got %v, want %v", result, expected)
	}
}

// TestInvalidFromCurrency validates that only valid currencies initialized in our map can be used in the "from" input
func TestInvalidFromCurrency(t *testing.T) {
	// Attempting to convert Chinese Renminbi (doesn't exist yet) to Swiss Franc
	_, err := Convert(100.0, "CNY", "CHF")

	// We don't not want to get an error
	if err == nil {
		t.Fatal("Expected error, got nil")
	}

	expected := "Unknown currency"
	if err.Error() != expected {
		t.Errorf("got %v, want %v", err.Error(), expected)
	}
}

// TestInvalidToCurrency validates that only valid currencies initialized in our map can be used in the "to" input
func TestInvalidToCurrency(t *testing.T) {
	// Attempting to convert U.S. dollar to Chinese Renminbi (doesn't exist yet)
	_, err := Convert(100.0, "USD", "CNY")

	if err == nil {
		t.Fatal("Expected error, got nil")
	}

	expected := "Unknown currency"
	if err.Error() != expected {
		t.Errorf("got %v, want %v", err.Error(), expected)
	}
}

// TestZeroConversion validates that if the input amount is 0, the conversion should also be 0
func TestZeroConversion(t *testing.T) {
	result, err := Convert(0.0, "EUR", "JPY")

	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	expected := 0.0
	if result != expected {
		t.Errorf("got %v, want %v", result, expected)
	}
}

// TestNegativeAmount validates the return message for negative input amounts
func TestNegativeAmount(t *testing.T) {
	_, err := Convert(-1.0, "USD", "CHF")

	if err == nil {
		t.Fatal("Expected error, got nil")
	}

	expected := "Invalid amounts entered"
	if err.Error() != expected {
		t.Errorf("got %v, want %v", err.Error(), expected)
	}
}

// TestRoundTrip will make sure that multiple conversions leading back to the same currency remains (about) the same
func TestRoundTrip(t *testing.T) {
	result1, err := Convert(100.0, "USD", "EUR")

	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	expected1 := 85.0
	if result1 != expected1 {
		t.Errorf("got %v, want %v", result1, expected1)
	}

	result2, err := Convert(result1, "EUR", "USD")

	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	expected2 := 100.0
	if result2 != expected2 {
		t.Errorf("got %v, want %v", result2, expected2)
	}
}
