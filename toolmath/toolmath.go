package toolmath

import "cmp"

// Max returns the maximum of two values that satisfy the cmp.Ordered constraint.
func Max[T cmp.Ordered](values ...T) T {
	if len(values) == 0 {
		var zero T  // In Go, for all numeric types, the default value is 0, for boolean it's false, and for strings, it's an empty string (“”)
		return zero // Return zero value for the type if no values are provided
	}

	max := values[0]
	for _, value := range values[1:] {
		if value > max {
			max = value
		}
	}
	return max
}
