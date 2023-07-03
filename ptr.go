package ptr

// To returns a pointer to v.
func To[T any](v T) *T {
	return &v
}

// Val returns the value pointed by p or (zero, false) if p is nil.
func Val[T any](p *T) (T, bool) {
	if p != nil {
		return *p, true
	}

	var zero T
	return zero, false
}

// ValOr returns the value pointed by p or v if nil.
func ValOr[T any](p *T, v T) T {
	if p != nil {
		return *p
	}

	return v
}

// ValOrZero returns the value pointed by p or zero if nil.
func ValOrZero[T any](p *T) T {
	v, _ := Val(p)
	return v
}
