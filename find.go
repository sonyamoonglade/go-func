package go_func

// Find will return the first element when the callback
// returns true or -1 if no element is found.
// It follows the same logic as the find() function in Javascript.
//
// If the list is empty then -1 is always returned.
func Find[T any](ss []T, fn func(T, int) bool) T {
	for i, value := range ss {
		if fn(value, i) {
			return value
		}
	}

	return *new(T)
}
