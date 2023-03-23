package template

import "testing"

func TestAny(t *testing.T) {
	tc := []struct {
		x   int
		exp int
		n   string
	}{
		{2, 3, "test"},
	}

	for _, i := range tc {
		c := testMe(i.x)
		if c != i.exp {
			t.Errorf("\n Error in test: %s, got: %d, exp: %d", i.n, c, i.exp)
		}
	}
}
