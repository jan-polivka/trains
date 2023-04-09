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
	bahnTypes := map[TrainType]func() types.ITrain{1: func() types.ITrain { return types.NewRegioSBahn() }}
	res, exists := bahnTypes[trainType]
	if exists {
		fmt.Println(res())
		fmt.Println("test")
	}
	if trainType == RegioSBahnType {
		return types.NewRegioSBahn(), nil
	}
	return nil, fmt.Errorf("Wrong type passed")
}
