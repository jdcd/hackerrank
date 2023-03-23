package repeatWorlds

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWhenWeHaveOpenSquareKeyShouldReturnFalse(t *testing.T) {
	//arrange
	input := "hello julian Hello HELLO! JULIAN."
	expected := WordsCount{
		"hello":  3,
		"julian": 2,
	}

	//act
	response := repeatWorlds(input)

	assert.Equal(t, expected, response)
}
