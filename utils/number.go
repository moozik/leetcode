package utils

func Abs[T int8 | int32 | int64 | int | float32 | float64](x T) T {
	if x < 0 {
		return -x
	}
	return x
}

func Min[T int8 | int32 | int64 | int | float32 | float64](x, y T) T {
	if x < y {
		return x
	}
	return y
}

func Max[T int8 | int32 | int64 | int | float32 | float64](x, y T) T {
	if x > y {
		return x
	}
	return y
}

func InArray[T comparable](nums []T, target T) bool {
	for _, item := range nums {
		if item == target {
			return true
		}
	}
	return false
}
