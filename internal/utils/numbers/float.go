package numbers

import "math"

func IsInt(a float64) bool {
	epsilon := 1e-9 // Margin of error
	_, frac := math.Modf(math.Abs(a))
	return frac < epsilon || frac > 1.0-epsilon
}
