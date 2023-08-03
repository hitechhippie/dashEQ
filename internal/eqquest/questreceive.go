package eqquest

type QuestReceive struct {
	QuestNPCId uint32
	Receives   string
	GiveCash   []uint32
	GiveItems  []uint32
	Says       string
}
