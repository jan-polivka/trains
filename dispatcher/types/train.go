package types

type Train struct {
	drive DriveType
}

func (train *Train) setDrive(drive DriveType) {
	train.drive = drive
}

type RegioSBahn struct {
	Train
}

func NewRegioSBahn() ITrain {
	return &RegioSBahn{
		Train: Train{
			drive: ElectricDrive,
		},
	}
}

type StrassenBahn struct {
	Train
}

func NewStrassenBahn() ITrain {
	return &StrassenBahn{
		Train: Train{
			drive: ElectricDrive,
		},
	}
}
