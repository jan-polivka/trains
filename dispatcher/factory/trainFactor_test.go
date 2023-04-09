package factory

import (
	"dispatcher/types"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GetTrainRegioSBahn(t *testing.T) {
	var result, _ = GetTrain(RegioSBahnType)
	var expected = types.NewRegioSBahn()
	assert.IsType(t, expected, result)
}

func Test_GetTrainStrassenBahn(t *testing.T) {
	// var result, _ = GetTrain("StrassenBahn")
	// fmt.Println(result)
}
