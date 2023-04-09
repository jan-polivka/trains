package factory

import (
	"dispatcher/types"
	"fmt"
)

type TrainType int

const (
	RegioSBahnType = 1 << iota
	StrassenBahnType
)

func GetTrain(trainType TrainType) (types.ITrain, error) {
	if trainType == RegioSBahnType {
		return types.NewRegioSBahn(), nil
	}
	return nil, fmt.Errorf("Wrong type passed")
}
