package minmax

type Ordered interface {
	string | uint | int | int8 | int16 | int32 | int64 | float32 | float64
}

func Min[T Ordered](a, b T) T {
	if a < b {
		return a
	}
	return b
}

func Max[T Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}
