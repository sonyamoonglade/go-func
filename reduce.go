package go_func

// Reduce continually applies the provided function
// over the slice. Reducing the elements to a single value.
//
// Returns a zero value of T if there are no elements in the slice. It will
// panic if the reducer is nil and the slice has more than one element (required
// to invoke reduce). Otherwise returns result of applying reducer from left to
// right.
func Reduce[IterType, AccType any](f func(acc AccType, el IterType, i int) AccType, arr []IterType, acc AccType) AccType {
	for i, v := range arr {
		acc = f(acc, v, i)
	}
	return acc
}
