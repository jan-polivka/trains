package factory

import (
	"testing"

	"github.com/jan-polivka/trains/dispatcher/types"
	"github.com/stretchr/testify/assert"
)

func Test_GetTrainRegioSBahn(t *testing.T) {
	var result, _ = GetTrain(RegioSBahnType)
	var expected = types.NewRegioSBahn()
	assert.IsType(t, expected, result)
}

func Test_GetTrainStrassenBahn(t *testing.T) {
	var result, _ = GetTrain(StrassenBahnType)
	var expected = types.NewStrassenBahn()
	assert.IsType(t, expected, result)
}

func Test_GetTrainWrongType(t *testing.T) {
	var result, err = GetTrain(-1)
	assert.Nil(t, result)
	assert.Error(t, err)
}
