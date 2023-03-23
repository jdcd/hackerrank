package temperature

import "testing"

func TestNearTo0(t *testing.T) {
	tc := []struct {
		temps []int
		exp   int
		name  string
	}{
		{
			temps: []int{-4, 1, 5, 4, -9, -8, -2, 9},
			exp:   1,
			name:  "basic",
		}, {
			temps: []int{-1, -5, -4, -9, -8, -2, -9},
			exp:   -1,
			name:  "all negatives",
		}, {
			temps: []int{1, 5, 4, 9, 8, 2, 9},
			exp:   1,
			name:  "all positives",
		}, {
			temps: []int{-1, -2, -3, 1},
			exp:   1,
			name:  "complex",
		},
	}

	c := 0
	for _, i := range tc {
		c = nearTo0(i.temps)
		if c != i.exp {
			t.Errorf("\n error in near to 0 test, case :%s, exp: %d, actual: %d", i.name, i.exp, c)
		}
	}
}
