package base62

import (
	"testing"
)

func TestEncode(t *testing.T) {
	tests := []struct {
		input    uint64
		expected string
	}{
		{0, "0"},
		{61, "z"},
		{62, "10"},
		{4815162342, "5FrG2"},
		{18446744073709551615, "LygHa16AHZG"},
	}

	for _, tt := range tests {
		if result := Encode(tt.input); result != tt.expected {
			t.Errorf("Encode(%d) = %s; want %s", tt.input, result, tt.expected)
		}
	}
}

func TestDecode(t *testing.T) {
	tests := []struct {
		input       string
		expected    uint64
		expectError bool
	}{
		{"0", 0, false},
		{"5FrG2", 4815162342, false},
		{"LygHa16AHZG", 18446744073709551615, false},
		{"", 0, true},
		{"abc-123", 0, true},
	}

	for _, tt := range tests {
		result, err := Decode(tt.input)
		if (err != nil) != tt.expectError {
			t.Errorf("Decode(%s) unexpected error state", tt.input)
		}
		if !tt.expectError && result != tt.expected {
			t.Errorf("Decode(%s) = %d; want %d", tt.input, result, tt.expected)
		}
	}
}
func BenchmarkEncodeStack(b *testing.B) {
	var val uint64 = 18446744073709551615 // Worst case maximum length
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Encode(val)
	}
}