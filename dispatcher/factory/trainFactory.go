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
	bahnTypes := map[int]func(){}
	res, exists := bahnTypes[1]
	if exists {
		res()
	}
	if trainType == RegioSBahnType {
		return types.NewRegioSBahn(), nil
	}
	return nil, fmt.Errorf("Wrong type passed")
}
