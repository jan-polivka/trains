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
	bahnTypes := map[TrainType]func() types.ITrain{RegioSBahnType: func() types.ITrain { return types.NewRegioSBahn() }}
	res, exists := bahnTypes[trainType]
	if exists {
		return res(), nil
	}
	return nil, fmt.Errorf("Wrong type passed")
}
