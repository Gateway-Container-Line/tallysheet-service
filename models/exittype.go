package models

type ExitType string

const (
	Coload          ExitType = "Coload"
	CargoAllOut     ExitType = "CargoAllOut"
	CargoPartialOut ExitType = "CargoPartialOut"
)
