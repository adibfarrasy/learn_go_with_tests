package service

func Sum(int_array []int) int {
	sum := 0
	for _, num := range int_array {
		sum += num
	}
	return sum
}
