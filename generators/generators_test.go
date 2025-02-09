package generators_test

import (
	"testing"

	"github.com/farmer032/golang-generators/generators"
	"github.com/stretchr/testify/assert"
)

func TestSliceGenerator(t *testing.T) {

	sampleSlice := make([]int, 10)

	asrt := assert.New(t)

	for idx, _ := range sampleSlice {
		sampleSlice[idx] = idx + 1
	}

	iter := generators.NewSliceGenerator(&sampleSlice)
	currentNumber := 1
	for iter.HasNext() {
		asrt.Equal(currentNumber, iter.Next())
		currentNumber++
	}
}
