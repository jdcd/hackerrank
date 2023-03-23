package temperature

func nearTo0(temps []int) int {
	nn, np := 0, 0
	pf, nf := false, false

	// if len of temps is 0, return 0 rule
	if len(temps) == 0 {
		return 0
	}

	for _, tmp := range temps {
		if tmp > 0 {
			if !pf {
				np = tmp
				pf = true
				continue
			}
			if tmp < np {
				np = tmp
			}
		} else if tmp < 0 {
			if !nf {
				nn = tmp
				nf = true
				continue
			}

			if tmp > nn {
				nn = tmp
			}
		}
	}

	if !pf {
		return nn
	} else if !nf {
		return np
	}

	// if closest negative and closest positive are equals, return positive value rule
	min := np
	if -nn < min {
		min = nn
	}
	return min

}
