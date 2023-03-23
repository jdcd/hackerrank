package migratory_birds

func migratoryBirds(arr []int32) int32 {
	bc := map[int32]int32{
		arr[0]: 1,
	}
	bird := arr[0]
	for i := 1; i < len(arr); i++ {
		current := arr[i]
		bc[current]++

		if bc[current] > bc[bird] || (bc[current] > bc[bird] && current < bird) {
			bird = current
		}

	}

	return bird

}
