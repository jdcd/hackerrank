package apple_orange

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_apple_orange(t *testing.T) {
	houseStart := int32(7)
	houseEnd := int32(15)
	appleTreePosition := int32(5)
	orangeTreePosition := int32(15)
	applePosition := []int32{-2, 2, 1}
	orangePosition := []int32{5, -6}

	expectedReturn := []int32{1, 1}
	fruitCount := countApplesAndOranges(houseStart, houseEnd, appleTreePosition, orangeTreePosition, applePosition, orangePosition)

	assert.Equal(t, expectedReturn, fruitCount)
}
