package queue

// Signed is a constraint that permits any signed integer type.
type Signed interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

// Unsigned is a constraint that permits any unsigned integer type.
type Unsigned interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

// Float is a constraint that permits any floating-point type.
type Float interface {
	~float32 | ~float64
}

// Ordered is a constraint that permits any ordered type: any type that
// supports the operators < <= >= >. It is equivalent to [cmp.Ordered] and is
// used for the priority of an [Item].
//
// Reference: https://pkg.go.dev/cmp#Ordered
type Ordered interface {
	Signed | Unsigned | Float | ~string
}
