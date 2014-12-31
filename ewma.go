package ewma

// Average returns the average of float64 numbers.
func Average(numbers ...float64) float64 {
	sum := 0.0
	if len(numbers) == 0 {
		return 0.0
	}
	for _, elem := range numbers {
		sum += elem
	}
	return sum / float64(len(numbers))
}
