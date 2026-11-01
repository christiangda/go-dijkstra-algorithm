package dijkstra

// Signed represent all go signed integers type
type Signed interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

// Unsigned represent all go signed unsigned integers type
type Unsigned interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

// Float represent all go float types
type Float interface {
	~float32 | ~float64
}

// Ordered is a constraint that permits any ordered type: any type that supports the operators < <= >= >.
// Reference: https://pkg.go.dev/cmp#Ordered
type Ordered interface {
	Signed | Unsigned | Float | ~string
}
