package homework

func average(input [15]float32) (result float32) {
	var sum float32 = 0.0
	for _, num := range input {
		sum += num
	}
	result = sum / 15
	return result
}
