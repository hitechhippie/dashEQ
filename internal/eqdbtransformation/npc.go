package eqdbtransformation

import (
	"dasheq/internal/eqdbobject"
)

func PopulateNpcListByZone(nnew *[]eqdbobject.NPC, ncur *[]eqdbobject.NPC, s2 *[]eqdbobject.Spawn2, se *[]eqdbobject.Spawnentry, z string) {

	spawn2zoneList := PopulateSpawn2ListByZone(s2, z)
	spawnEntryzonelist := JoinSpawnentryWithSpawn2(se, spawn2zoneList)
	JoinNpcWithSpawnEntry(nnew, ncur, spawnEntryzonelist)
}

func JoinNpcWithSpawnEntry(nnew *[]eqdbobject.NPC, ncur *[]eqdbobject.NPC, se *[]eqdbobject.Spawnentry) {
	for _, spawnEntry := range *se {
		for _, npc := range *ncur {
			if spawnEntry.NpcID == npc.Id {
				*nnew = append(*nnew, npc)
			}
		}
	}
}

func DistinctNpcList(nnew *[]eqdbobject.NPC, ncur *[]eqdbobject.NPC) {
	for _, v := range *ncur {
		skip := false
		for _, u := range *nnew {
			if v == u {
				skip = true
				break
			}
		}
		if !skip {
			*nnew = append(*nnew, v)
		}
	}
}

// implementation for sort.Interface to sort npcs by level
func (n NpcListByLevelSorted) Len() int {
	return len(n)
}

// implementation for sort.Interface to sort npcs by level
func (n NpcListByLevelSorted) Less(i, j int) bool {
	return n[i].Level < n[j].Level
}

// implementation for sort.Interface to sort npcs by level
func (n NpcListByLevelSorted) Swap(i, j int) {
	n[i], n[j] = n[j], n[i]
}

/*
SELECT npc_types.id FROM npc_types,spawn2,spawnentry,spawngroup WHERE spawn2.zone = 'befallen' AND spawnentry.spawngroupID = spawn2.spawngroupID AND spawnentry.npcID = npc_types.id AND spawngroup.id = spawnentry.spawngroupID GROUP BY npc_types.id;
*/
