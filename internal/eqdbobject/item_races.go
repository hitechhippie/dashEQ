package eqdbobject

import "strings"

type RaceList uint32

const (
	HUM RaceList = 1 << iota
	BAR
	ERU
	ELF
	HIE
	DEF
	HEF
	DWF
	TRL
	OGR
	HFL
	GNM
	IKS
	VAH
	FRG
	DRA
	raceMax
)

func (r RaceList) String() string {

	if r == raceMax || r == 0 {
		return "ALL"
	}

	switch r {
	case HUM:
		return "HUM"
	case BAR:
		return "BAR"
	case ERU:
		return "ERU"
	case ELF:
		return "ELF"
	case HIE:
		return "HIE"
	case DEF:
		return "DEF"
	case HEF:
		return "HEF"
	case DWF:
		return "DWF"
	case TRL:
		return "TRL"
	case OGR:
		return "OGR"
	case HFL:
		return "HFL"
	case GNM:
		return "GNM"
	case IKS:
		return "IKS"
	case VAH:
		return "VAH"
	case FRG:
		return "FRG"
	case DRA:
		return "DRA"
	}

	// multiple races
	var races []string
	for race := HUM; race < raceMax; race <<= 1 {
		if r&race != 0 {
			races = append(races, race.String())
		}
	}
	return strings.Join(races, " | ")
}
