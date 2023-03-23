package template

import "testing"

func TestAny(t *testing.T) {
	tc := []struct {
		x   []int
		exp int
		n   string
	}{
		{[]int{1, 2, -3, 5, 6, 1, 2, -3, 23}, 23, "basic"},
		{[]int{1, 2, 3, 5, 6, 1, 2, 3, 25}, 25, "positives"},
		{[]int{-1, -2, -3, -5, -6, -1, -2, -3, -25}, -1, "negatives"},
	}

	for _, i := range tc {
		c := FindLargest(i.x)
		if c != i.exp {
			t.Errorf("\n Error in test: %s, got: %d, exp: %d", i.n, c, i.exp)
		}
	}
}
