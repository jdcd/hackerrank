package migratory_birds

import "testing"

func TestSampleCase1(t *testing.T) {
	arr := []int32{1, 2, 3, 4, 5, 4, 3, 2, 1, 3, 4}
	exp := int32(3)

	r := migratoryBirds(arr)

	if r != exp {
		t.Errorf("\nError, expected: %d actual: %d", exp, r)
	}
}

func TestSampleCase5(t *testing.T) {
	arr := []int32{1, 2, 3, 4, 5, 4, 3, 2, 1, 3, 4}
	exp := int32(3)

	r := migratoryBirds(arr)

	if r != exp {
		t.Errorf("\nError, expected: %d actual: %d", exp, r)
	}
}
