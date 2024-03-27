package others

func sumSlice[T int | float32 | float64](slice []T) T {
	var sum T
	for _, n := range slice {
		sum += n
	}
	return sum
}
