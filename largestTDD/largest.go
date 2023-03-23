package template

func FindLargest(numbers []int) int {
	l := numbers[0]
	for _, n := range numbers {
		if n > l {
			l = n
		}
	}
	return l
}
