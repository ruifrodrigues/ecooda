package utils

type Number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~float32 | ~float64
}

func Pointer[T Number](val T) *T {
	a := new(T)
	*a = val

	return a
}
