package eqquest

type QuestNPC struct {
	Id   uint32
	Name string
	Zone uint32
	File string
}

/*
QuestNPC
  Id (NPC->Id) uint32
  Name string
  Zone (Zone->Id) uint32
  QuestHear
    Id (QuestNPC->Id) uint32
	Hears string
	GiveItems {uint32...} (Item->Id)
	Says string
  QuestReceive
    Id (QuestNPC->Id) uint32
	Receives {uint32...} (Item->Id)
	GiveCash
	GiveItems {uint32...} (Item->Id)
*/

/*
QuestReward(mob, uint32 copper);

QuestReward(mob, uint32 copper, uint32 silver, uint32 gold, uint32 platinum, uint32 item_id, uint32 exp); ** this is most common found?

QuestReward(mob, uint32 copper, uint32 silver, uint32 gold);
QuestReward(mob, uint32 copper, uint32 silver);
QuestReward(mob, uint32 copper, uint32 silver, uint32 gold, uint32 platinum, uint32 item_id, uint32 exp, bool faction);
QuestReward(mob);
QuestReward(mob, uint32 copper, uint32 silver, uint32 gold, uint32 platinum);
QuestReward(mob, uint32 copper, uint32 silver, uint32 gold, uint32 platinum, uint32 item_id);
*/
