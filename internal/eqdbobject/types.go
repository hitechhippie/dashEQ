package eqdbobject

type DataSet struct {
	Name     string
	Count    uint32
	LoadTime string
}

type Zone struct {
	Id         uint32
	Short_name string
	Long_name  string
	Outdoor    bool
	Expansion  ZoneExpansion
}

type ZoneExpansion uint8

type NPC struct {
	Id          uint32
	Name        string
	Lastname    string
	Level       uint8
	MaxLevel    uint8
	Race        NPCrace
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

type NPCrace uint16

type Spell struct {
	Id        uint32
	Name      string
	Classes1  uint8
	Classes2  uint8
	Classes3  uint8
	Classes4  uint8
	Classes5  uint8
	Classes6  uint8
	Classes7  uint8
	Classes8  uint8
	Classes9  uint8
	Classes10 uint8
	Classes11 uint8
	Classes12 uint8
	Classes13 uint8
	Classes14 uint8
	Classes15 uint8
	Classes16 uint8
	NewIcon   uint16
}

type Skill struct {
	Skill SkillID
	Class ClassID
	Level uint8
	Cap   uint8
}

type Spawngroup struct {
	Id            uint32
	Name          string
	Spawn_limit   uint8
	Max_x         float32
	Min_x         float32
	Max_y         float32
	Min_y         float32
	Delay         uint32
	Mindelay      uint32
	Despawn       uint8
	Despawn_timer uint32
	Wp_spawns     uint8
}

type Spawn2 struct {
	Id           uint32
	SpawngroupID uint32
	Zone         string
	X            float32
	Y            float32
	Z            float32
	RespawnTime  uint32
}

type Spawnentry struct {
	SpawngroupID uint32
	NpcID        uint32
	Chance       uint8
}

type ClassID uint8

type SkillID uint8

type ClassListBitmask uint32

// Populate classlist bitmask with iota
const (
	WAR ClassListBitmask = 1 << iota
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

type RaceListBitmask uint32

// Populate races bitmask with iota
const (
	HUM RaceListBitmask = 1 << iota
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

type EQEmuWeirdBool int

type EQweight float32

type Item struct {
	Id           uint32
	Name         string
	AC           int
	HP           int
	Icon         uint32
	Itemtype     ItemTypes
	Reclevel     uint8
	Reqlevel     uint8
	ScrollEffect uint32
	Size         ItemSize
	Slots        ItemSlotsBitmask
	Weight       EQweight
	Magic        bool
	Mana         int
	DBnodrop     EQEmuWeirdBool
	DBnorent     EQEmuWeirdBool
	Classes      ClassListBitmask
	Damage       uint16
	Delay        uint16
	Races        RaceListBitmask
	Nodrop       bool
	Norent       bool
}

type ItemSize uint8

type ItemTypes uint32

type ItemSlotsBitmask uint32

// Populate item slots bitmask with iota
const (
	Charm ItemSlotsBitmask = 1 << iota
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
