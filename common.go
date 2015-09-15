// Package rbtree implements red-black tree introduced in "Introduction to Algorithms".
/*
Under current language spec, there are following patterns to implement generic containers:

1) the pattern used by sort package

	type Interface interface {
		Compare(another Interface) int
	}

2) using callbacks:

	func Compare(x, y interface{}) int

3) using `go generate` to generate code for specific type.

This package uses callbacks. Using tricks to get pointer of empty interface
values can avoid data copying and runtime assertions, therefore greatly improve
performance.  It's your responsibility to assure type safe.
*/
package rbtree

import (
	"strings"
	"unsafe"
)

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

type iface struct {
	typ  unsafe.Pointer
	data unsafe.Pointer
}

// ValuePtr is a helper function to get the pointer to value stored in empty interface.
func ValuePtr(v interface{}) unsafe.Pointer {
	return ((*iface)(unsafe.Pointer(&v))).data
}

func compareInt(x, y interface{}) int {
	a := *(*int)(ValuePtr(x))
	b := *(*int)(ValuePtr(y))
	if a > b {
		return 1
	} else if a < b {
		return -1
	} else {
		return 0
	}
}

func compareString(x, y interface{}) int {
	a := *(*string)(ValuePtr(x))
	b := *(*string)(ValuePtr(y))
	return strings.Compare(a, b)
}
