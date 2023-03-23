package birthday_chocolate

// todo add context
func birthday(s []int32, d int32, m int32) int32 {
	w := 0

	for i := 0; i <= len(s)-int(m); i++ {
		fp := i + int(m)
		sum := int32(0)
		for j := i; j < fp; j++ {
			sum += s[j]
		}
		if sum == d {
			w++
		}
	}

	return int32(w)
}
