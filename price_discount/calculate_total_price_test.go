package price_discount

import "testing"

func TestAny(t *testing.T) {
	tc := []struct {
		x   []int
		d   int
		exp int
		n   string
	}{
		{[]int{100, 20, 30, 40, 50, 6}, 50, 196, "positive"},
		{[]int{1000, 1000, 1000, 400}, 40, 3000, "repeat value"},
		{[]int{3000, 3000, 3000, 3000, 3000}, 50, 13500, "equal"},
	}

	for _, i := range tc {
		c := CalculateTotalPrice(i.x, i.d)
		if c != i.exp {
			t.Errorf("\n Error in test: %s, got: %d, exp: %d", i.n, c, i.exp)
		}
	}
}
