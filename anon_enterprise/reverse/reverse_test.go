package reverse

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMianFlow(t *testing.T) {
	//arragne
	input := []int{1, 2, 3, 4, 5}
	expected := []int{5, 4, 3, 2, 1}

	//act
	respose := reverseSlice(input)

	//assert
	assert.Equal(t, expected, respose)

}

func TestEmptyArrayFlow(t *testing.T) {
	//arragne
	input := []int{}
	expected := []int{}

	//act
	respose := reverseSlice(input)

	//assert
	assert.Equal(t, expected, respose)

}
