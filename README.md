Package rbtree implements red-black tree introduced in "Introduction to Algorithms".

Under current language spec, there are following patterns to implement generic containers:

1) the pattern used by sort package

```go
	type Interface interface {
		Compare(another Interface) int
	}
```

2) using callbacks:

```go
	func Compare(x, y interface{}) int
```

3) using `go generate` to generate code for specific type.

This package uses callbacks. Using tricks to get pointer of empty interface values can avoid data copying and greatly improve performance.
It's your responsibility to assure type safe.

```go
	// ValuePtr is a helper function to get the pointer to value stored in
	// empty interface.
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
```

See test code for more usage.
