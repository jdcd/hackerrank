package min_max

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInvalidInputArray(t *testing.T) {
	//arrange
	input := []Item{}
	errorExpected := emptyArrayErr

	//act
	min, max, err := GetMinMax(input)

	//assert
	assert.Equal(t, 0, min)
	assert.Equal(t, 0, max)
	assert.Equal(t, errorExpected, err)
}

func TestMainFlow(t *testing.T) {
	//arrange
	input := getArrayExample()
	minExpected := 450
	maxExpected := 1000

	//act
	min, max, err := GetMinMax(input)

	//assert
	assert.Equal(t, minExpected, min)
	assert.Equal(t, maxExpected, max)
	assert.Equal(t, nil, err)
}

func getArrayExample() []Item {
	return []Item{
		{
			ID:    1,
			Name:  "cheems figure",
			Price: 500,
			Desc:  "a cool figure",
		},
		{
			ID:    2,
			Name:  "doge figure",
			Price: 450,
			Desc:  "doge figure",
		},
		{
			ID:    3,
			Name:  "Yuno pillow",
			Price: 1000,
			Desc:  "confortable branding",
		},
	}
}
