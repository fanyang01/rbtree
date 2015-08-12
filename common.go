/*
Under current language spec, there are following patterns to implement generic containers:

1) the pattern used by sort package

	type Interface interface {
		// Compare compares value of reciever with another, and returns an integer:
		// 0 if reciever is equal to another,
		// 1 if reciever is greater than another, and
		// -1 if reciever is less than another
		Compare(another Interface) int
	}

2) using some callbacks:

	func Compare(x, y interface{}) int

3) using `go generate` to generate code for specific type.

This package uses callbacks.
*/
package rbtree

// CompareFunc compares x and y, and returns an integer
// = 0 if x is equal to y,
// > 0 if x is greater than y, and
// < 0 if x is less than y
type CompareFunc func(x, y interface{}) int

// These functions are provided for convinence
var (
	CompareInt    CompareFunc = compareInt
	CompareString             = compareString
)

func compareInt(x, y interface{}) int {
	a, b := x.(int), y.(int)
	if a > b {
		return 1
	} else if a < b {
		return -1
	} else {
		return 0
	}
}

func compareString(x, y interface{}) int {
	a, b := x.(string), y.(string)
	if a > b {
		return 1
	} else if a < b {
		return -1
	} else {
		return 0
	}
}
