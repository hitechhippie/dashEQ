package eqquest

type QuestNPC struct {
	Id       uint32
	Name     string
	Zone     uint32
	ZoneName string
	File     string
}

type QuestReceive struct {
	QuestNPCId uint32
	Receives   string
	GiveCash   []uint32
	GiveItems  []uint32
	Says       string
}

type QuestHear struct {
	QuestNPCId   uint32
	QuestNPCName string
	Hears        string
	GiveItems    []uint32
	Says         string
}
