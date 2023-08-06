package eqdbobject

import "strings"

type ItemSlots uint32

const (
	Charm ItemSlots = 1 << iota
	Ear_1
	Head
	Face
	Ear_2
	Neck
	Shoulder
	Arms
	Back
	Bracer_1
	Bracer_2
	Range
	Hands
	Primary
	Secondary
	Ring_1
	Ring_2
	Chest
	Legs
	Feet
	Waist
	Powersource
	Ammo
	slotMax
)

func (s ItemSlots) String() string {
	if s == slotMax || s == 0 {
		return "ALL"
	}

	switch s {
	case Charm:
		return "Charm"
	case Ear_1:
		return "Ear"
	case Head:
		return "Head"
	case Face:
		return "Face"
	case Ear_2:
		return "Ear"
	case Neck:
		return "Neck"
	case Shoulder:
		return "Shoulder"
	case Arms:
		return "Arms"
	case Back:
		return "Back"
	case Bracer_1:
		return "Bracer"
	case Bracer_2:
		return "Bracer"
	case Range:
		return "Range"
	case Hands:
		return "Hands"
	case Primary:
		return "Primary"
	case Secondary:
		return "Secondary"
	case Ring_1:
		return "Ring"
	case Ring_2:
		return "Ring"
	case Chest:
		return "Chest"
	case Legs:
		return "Legs"
	case Feet:
		return "Feet"
	case Waist:
		return "Waist"
	case Powersource:
		return "Powersource"
	case Ammo:
		return "Ammo"
	}

	// multiple races
	var slots []string
	for slot := Charm; slot < slotMax; slot <<= 1 {
		if s&slot != 0 {
			slots = append(slots, slot.String())
		}
	}
	return strings.Join(slots, " | ")
}
