package slice

func Contains[T comparable](elems []T, v T) bool {
	for _, s := range elems {
		if v == s {
			return true
		}
	}
	return false
}

func Remove[T comparable](elems []T, v T) []T {
	var res []T
	for _, s := range elems {
		if v != s {
			res = append(res, s)
		}
	}
	return res
}
