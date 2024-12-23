package utils

func Filter[T interface{}](arr []T, fn func(entry T) bool) []T {
	res := []T{}
	for _, v := range arr {
		if fn(v) {
			res = append(res, v)
		}
	}
	return res
}
