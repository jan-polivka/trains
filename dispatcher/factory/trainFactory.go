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
	bahnTypes := map[int]func(text string){1: func(text string) { fmt.Println(text) }}
	res, exists := bahnTypes[int(trainType)]
	if exists {
		res("Hello")
	}
	if trainType == RegioSBahnType {
		return types.NewRegioSBahn(), nil
	}
	return nil, fmt.Errorf("Wrong type passed")
}
