package divisible_sum_pairs

import (
	"fmt"
	"testing"
)

func TestSample1(t *testing.T) {
	ar := []int32{1, 3, 2, 6, 1, 2}
	n, k := int32(6), int32(3)
	exp := int32(5)

	r := divisibleSumPairs(n, k, ar)

	if r != exp {
		t.Errorf("\ntest fail: expected: %d, actual: %d ", exp, r)
	} else {
		fmt.Println("Success")
	}
}
