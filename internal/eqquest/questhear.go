package eqquest

type QuestHear struct {
	QuestNPCId   uint32
	QuestNPCName string
	Hears        string
	GiveItems    []uint32
	Says         string
}
