package factory

import (
	"dispatcher/types"
	"fmt"
)

const (
	RegioSBahnType = 1 << iota
	StrassenBahnType
)

func GetTrain(trainType string) (types.ITrain, error) {
	if trainType == "RegioSBahn" {
		return types.NewRegioSBahn(), nil
	}
	return nil, fmt.Errorf("Wrong type passed")
}
