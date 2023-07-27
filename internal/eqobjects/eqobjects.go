package eqobjects

type DataSet struct {
	Name     string
	Count    uint32
	LoadTime string
}

type NPC struct {
	Id          uint32
	Name        string
	Level       uint8
	Race        uint16
	Class       uint8
	HP          uint32
	Mana        uint32
	LootTable   uint32
	NPCSpells   uint16
	NPCFaction  uint16
	MinDmg      uint16
	MaxDmg      uint16
	AttackCount uint16
	RunSpeed    float32
	MR          uint16
	CR          uint16
	DR          uint16
	FR          uint16
	PR          uint16
	AC          uint16
}

type Zone struct {
	Id         uint16
	Short_name string
	Long_name  string
	Outdoor    uint8
	Expansion  uint8
}
