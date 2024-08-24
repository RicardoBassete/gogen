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

// Returns the last element of an slice/array.
// If the slice has more than one element, returns the last element of the slice.
// If the slice is empty, returns the zero value of T
func GetLast[T any](array []T) T {
	if len(array) == 0 {
		var result T
		return result // Returns the zero value of T
	}
	return array[len(array)-1]
}

// Compares two slices and returns true if both slices are equal
func CompArr[T comparable](arr1 []T, arr2 []T) bool {
	if len(arr1) != len(arr2) {
		return false
	}
	if len(arr1) == 0 && len(arr2) == 0 {
		return true
	}
	for i, v := range arr1 {
		if v != arr2[i] {
			return false
		}
	}
	return true
}

// Splits the input slice in half and returns both halves.
// If the slice is empty returns two empty slices.
func SplitSlice[T any](slice []T) ([]T, []T) {
	l := len(slice)
	if l == 0 {
		return []T{}, []T{}
	}
	mid := l / 2
	return slice[:mid], slice[mid:]
}

// Creates a copy of a slice, filtered down to just the elements from the
// given slice that pass the test implemented by the provided function.
func Filter[T any](slice []T, test func(T) bool) []T {
	result := make([]T, 0)
	for _, element := range slice {
		if test(element) {
			result = append(result, element)
		}
	}
	return result
}
