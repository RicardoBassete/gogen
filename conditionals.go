package gogen

// Generic function that works as a ternary operator.
// If the condition is true, returns vTrue, otherwise returns vFalse
func If[T any](condition bool, vTrue T, vFalse T) T {
	if condition {
		return vTrue
	} else {
		return vFalse
	}
}
