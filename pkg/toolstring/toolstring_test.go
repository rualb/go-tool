package toolstring

import (
	"testing"
)

func TestEllipses(t *testing.T) {
	tests := []struct {
		str    string
		maxLen int
		want   string
	}{
		{"123456789", 5, "123.."},
		{"1234", 5, "1234"},
		{"123456", 5, "123.."},
		{"12", 5, "12"},
		{"", 5, ""},
		{"123", 2, ".."},
	}

	for _, tt := range tests {
		t.Run(tt.str, func(t *testing.T) {
			got := Ellipses(tt.str, tt.maxLen)
			if got != tt.want {
				t.Errorf("Ellipses(%q, %d) = %q; want %q", tt.str, tt.maxLen, got, tt.want)
			}
		})
	}
}

func TestLeft(t *testing.T) {
	tests := []struct {
		str  string
		n    int
		want string
	}{
		{"123456789", 5, "12345"},
		{"1234", 5, "1234"},
		{"123456", 5, "12345"},
		{"12", 5, "12"},
		{"", 5, ""},
		{"abcdef", 0, ""},
	}

	for _, tt := range tests {
		t.Run(tt.str, func(t *testing.T) {
			got := Left(tt.str, tt.n)
			if got != tt.want {
				t.Errorf("Left(%q, %d) = %q; want %q", tt.str, tt.n, got, tt.want)
			}
		})
	}
}

func TestRight(t *testing.T) {
	tests := []struct {
		str  string
		n    int
		want string
	}{
		{"123456789", 5, "56789"},
		{"1234", 5, "1234"},
		{"123456", 5, "23456"},
		{"12", 5, "12"},
		{"", 5, ""},
		{"abcdef", 0, ""},
	}

	for _, tt := range tests {
		t.Run(tt.str, func(t *testing.T) {
			got := Right(tt.str, tt.n)
			if got != tt.want {
				t.Errorf("Right(%q, %d) = %q; want %q", tt.str, tt.n, got, tt.want)
			}
		})
	}
}
