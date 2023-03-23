package birthday_chocolate

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBirthDay0(t *testing.T) {
	s := []int32{1, 2, 1, 3, 2}
	d, m := int32(3), int32(2)
	exp := int32(2)

	r := birthday(s, d, m)

	assert.Equal(t, exp, r)
}

func TestBirthDay1(t *testing.T) {
	s := []int32{1, 1, 1, 1, 1, 1}
	d, m := int32(3), int32(2)
	exp := int32(0)

	r := birthday(s, d, m)

	assert.Equal(t, exp, r)
}

func TestBirthDay2(t *testing.T) {
	s := []int32{4}
	d, m := int32(4), int32(1)
	exp := int32(1)

	r := birthday(s, d, m)

	assert.Equal(t, exp, r)
}
