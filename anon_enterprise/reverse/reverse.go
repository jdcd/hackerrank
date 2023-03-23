package reverse

func reverseSlice(mySlice []int) []int {

	/* todo only work with half array*/
	reverse := make([]int, len(mySlice))
	j := 0
	for i := len(mySlice) - 1; i >= 0; i-- {
		reverse[j] = mySlice[i]
		j++
	}
	return reverse
}
