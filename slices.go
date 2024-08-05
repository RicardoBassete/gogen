package gogen

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
