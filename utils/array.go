package utils

func InArray[E comparable](v E, s []E) bool {
	for _, vs := range s {
		if v == vs {
			return true
		}
	}
	return false
}
