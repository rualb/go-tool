package toolmath

import (
	"testing"
	// toolmath2 "github.com/rualb/go-tool/pkg/toolmath"
)

func TestMaxInt(t *testing.T) {
	tests := []struct {
		x, y, expected int
	}{
		{3, 5, 5},
		{10, 7, 10},
		{-1, -2, -1},
		{0, 0, 0},
	}

	for _, tt := range tests {
		result := Max(tt.x, tt.y)
		if result != tt.expected {
			t.Errorf("Max(%d, %d) = %d; want %d", tt.x, tt.y, result, tt.expected)
		}
	}
}

func TestMaxFloat64(t *testing.T) {
	tests := []struct {
		x, y, expected float64
	}{
		{2.71, 3.14, 3.14},
		{1.1, 1.9, 1.9},
		{-1.5, -1.4, -1.4},
		{0.0, 0.0, 0.0},
	}

	for _, tt := range tests {
		result := Max(tt.x, tt.y)
		if result != tt.expected {
			t.Errorf("Max(%.2f, %.2f) = %.2f; want %.2f", tt.x, tt.y, result, tt.expected)
		}
	}
}

func TestMaxString(t *testing.T) {
	tests := []struct {
		x, y, expected string
	}{
		{"apple", "banana", "banana"},
		{"cherry", "apple", "cherry"},
		{"equal", "equal", "equal"},
		{"", "non-empty", "non-empty"},
	}

	for _, tt := range tests {
		result := Max(tt.x, tt.y)
		if result != tt.expected {
			t.Errorf("Max(%s, %s) = %s; want %s", tt.x, tt.y, result, tt.expected)
		}
	}
}

func TestMaxVariadic(t *testing.T) {
	tests := []struct {
		values []int
		expect int
	}{
		{values: []int{10, 20, 30, 40}, expect: 40},
		{values: []int{5, 5, 5, 5}, expect: 5},
		{values: []int{-1, -2, -3, -4}, expect: -1},
		{values: []int{}, expect: 0}, // Expect zero value for the type
	}

	for _, test := range tests {
		result := Max(test.values...)
		if result != test.expect {
			t.Errorf("Max(%v) = %d; want %d", test.values, result, test.expect)
		}
	}
}

func TestMaxVariadicFloat(t *testing.T) {
	tests := []struct {
		values []float64
		expect float64
	}{
		{values: []float64{3.5, 2.1, 7.8, 1.2}, expect: 7.8},
		{values: []float64{1.1, 4.4, 3.3}, expect: 4.4},
		{values: []float64{5.5, 5.5, 5.5}, expect: 5.5},
		{values: []float64{}, expect: 0}, // Expect zero value for the type
	}

	for _, test := range tests {
		result := Max(test.values...)
		if result != test.expect {
			t.Errorf("Max(%v) = %f; want %f", test.values, result, test.expect)
		}
	}
}

func TestMaxVariadicString(t *testing.T) {
	tests := []struct {
		values []string
		expect string
	}{
		{values: []string{"apple", "banana", "cherry"}, expect: "cherry"},
		{values: []string{"grape", "banana", "apple"}, expect: "grape"},
		{values: []string{"apple", "apple"}, expect: "apple"},
		{values: []string{}, expect: ""}, // Expect empty string for the type
	}

	for _, test := range tests {
		result := Max(test.values...)
		if result != test.expect {
			t.Errorf("Max(%v) = %s; want %s", test.values, result, test.expect)
		}
	}
}
