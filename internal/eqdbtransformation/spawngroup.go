package eqdbtransformation

import "dasheq/internal/eqdbobject"

func JoinSpawngroupWithSpawnentry(sg *[]eqdbobject.Spawngroup, se *[]eqdbobject.Spawnentry) *[]eqdbobject.Spawngroup {
	var spawnGroupListJoined []eqdbobject.Spawngroup

	for _, spawnEntry := range *se {
		for _, spawnGroup := range *sg {
			if spawnEntry.SpawngroupID == spawnGroup.Id {
				spawnGroupListJoined = append(spawnGroupListJoined, spawnGroup)
			}
		}
	}

	return &spawnGroupListJoined
}
