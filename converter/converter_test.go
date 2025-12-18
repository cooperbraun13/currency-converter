package converter

import "testing"

// TestConvertUSDToEur calls Convert with a specified amount and checks
// for proper conversion from USD to Euro
func TestConvertUSDToEur(t *testing.T) {
	result, err := Convert(100, "USD", "EUR")

	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	expected := 85.0 // 100 * 0.85
	if result != expected {
		t.Errorf("got %v, want %v", result, expected)
	}
}
