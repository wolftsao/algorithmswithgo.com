package module01

// Sum will sum up all of the numbers passed
// in and return the result
func Sum(numbers []int) int {
	sum := 0

	for _, v := range numbers {
		sum += v
	}

	return sum
}

func RecurSum(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}

	return numbers[0] + RecurSum(numbers[1:])
}
