package types

type DriveType int

const (
	ElectroDieselDrive DriveType = 1 << iota
	ElectricDrive
)

type ITrain interface {
	setDrive(drive DriveType)
}
