package eqdbtransformation

import "dasheq/internal/eqdbobject"

func JoinSpawnentryWithSpawn2(se *[]eqdbobject.Spawnentry, s2 *[]eqdbobject.Spawn2) *[]eqdbobject.Spawnentry {
	var spawnEntryListJoined []eqdbobject.Spawnentry

	for _, spawn2 := range *s2 {
		for _, spawnEntry := range *se {
			if spawn2.SpawngroupID == spawnEntry.SpawngroupID {
				spawnEntryListJoined = append(spawnEntryListJoined, spawnEntry)
			}
		}
	}

	return &spawnEntryListJoined
}
