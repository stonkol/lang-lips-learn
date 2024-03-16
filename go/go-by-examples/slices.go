// SLICES
// they are an important data type in Go, giving a more powerful interface to sequences than arrays.

package main

import (
	"fmt"
)

// Unlike arrays, slices are typed only by the elements they contain (not the number of elements).
// An uninitialized slice equals to nil and has length 0.
func main() {
	var s []string
	fmt.Println("uninit:", s, s == nil, len(s) == 0)

	// To create an empty slice with non-zero length, use the builtin make.
	// Here we make a slice of strings of length 3 (initially zero-valued).
	// By default a new slice’s capacity is equal to its length;
	// if we know the slice is going to grow ahead of time, it’s possible
	// to pass a capacity explicitly as an additional parameter to make.
	s = make([]string, 3)
	fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))

	// We can set and get just like with arrays.
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	println("set: ", s)
	println("get: ", s[2])

	// len returns the length of the slice as expected.
	println("len: ", len(s))
}
