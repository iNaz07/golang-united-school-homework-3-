package homework

func average(input [15]float32) (result float32) {
	//Place your code here
	var count float32
	for _, num := range input {
		if num != 0 {
			result += num
			count++
		}
	}
	return result / count

}
