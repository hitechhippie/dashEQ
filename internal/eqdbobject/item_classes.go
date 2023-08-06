package eqdbobject

import "strings"

type ClassList uint32

const (
	WAR ClassList = 1 << iota
	CLR
	PAL
	RNG
	SHD
	DRU
	MNK
	BRD
	ROG
	SHM
	NEC
	WIZ
	MAG
	ENC
	BST
	BER
	classMax
)

func (c ClassList) String() string {

	if c == classMax || c == 0 {
		return "ALL"
	}

	switch c {
	case WAR:
		return "WAR"
	case CLR:
		return "CLR"
	case PAL:
		return "PAL"
	case RNG:
		return "RNG"
	case SHD:
		return "SHD"
	case DRU:
		return "DRU"
	case MNK:
		return "MNK"
	case BRD:
		return "BRD"
	case ROG:
		return "ROG"
	case SHM:
		return "SHM"
	case NEC:
		return "NEC"
	case WIZ:
		return "WIZ"
	case MAG:
		return "MAG"
	case ENC:
		return "ENC"
	case BST:
		return "BST"
	case BER:
		return "BER"
	}

	// multiple races
	var classes []string
	for class := WAR; class < classMax; class <<= 1 {
		if c&class != 0 {
			classes = append(classes, class.String())
		}
	}
	return strings.Join(classes, " | ")
}
