package maria_baseBall

func breakingRecords(scores []int) []int {
	minRecord, maxRecord := scores[0], scores[0]
	breakMin, breakMax := 0, 0
	for _, score := range scores {
		if score < minRecord {
			breakMin++
			minRecord = score
		}

		if score > maxRecord {
			breakMax++
			maxRecord = score
		}
	}
	return []int{breakMax, breakMin}

}
