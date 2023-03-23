package maria_baseBall

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWhenMariaBreaksMinimumAndMaxRecords(t *testing.T) {
	//arrange
	inputArray := []int{10, 5, 20, 20, 4, 5, 2, 25, 1}
	expectedResponse := []int{2, 4}

	//act
	result := breakingRecords(inputArray)

	//assert
	assert.Equal(t, expectedResponse, result)

}
