package divisible_sum_pairs

func divisibleSumPairs(n int32, k int32, ar []int32) int32 {
	var count int32
	for i := 0; i < len(ar); i++ {
		for j := 0; j < len(ar); j++ {
			if i < j {
				if (ar[i]+ar[j])%k == 0 {
					count++
				}

			}
		}

	}

	return count

}
