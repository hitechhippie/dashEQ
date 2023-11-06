package eqdbtransformation

import "dasheq/internal/eqdbobject"

type SpellListByClass struct {
	Class   string
	Id      uint32
	Name    string
	Level   uint8
	NewIcon uint16
	Scroll  uint32
}

// implementation for sort.Interface to sort spells by level
type SpellListByClassSorted []SpellListByClass

type ZoneListByExpansion struct {
	Expansion  eqdbobject.ZoneExpansion
	Name       string
	Short_name string
}

// implementation for sort.Interface to sort zones by expansion
type ZoneListByExpansionSorted []ZoneListByExpansion

// implementation for sort.Interface to sort npcs by level
type NpcListByLevelSorted []eqdbobject.NPC
