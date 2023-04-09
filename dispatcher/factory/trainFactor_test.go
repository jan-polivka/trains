package factory

import (
	"dispatcher/types"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GetTrainRegioSBahn(t *testing.T) {
	var result, _ = GetTrain("RegioSBahn")
	var expected = types.NewRegioSBahn()
	assert.IsType(t, expected, result)
}

func Test_GetTrainStrassenBahn(t *testing.T) {
	var result, _ = GetTrain("StrassenBahn")
	fmt.Println(result)
}

func Test_GetTrainError(t *testing.T) {
	var result, err = GetTrain("WrongTrain")
	assert.Nil(t, result)
	assert.Error(t, err)
}
