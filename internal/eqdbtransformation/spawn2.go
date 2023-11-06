package eqdbtransformation

import "dasheq/internal/eqdbobject"

func PopulateSpawn2ListByZone(s2 *[]eqdbobject.Spawn2, z string) *[]eqdbobject.Spawn2 {
	var spawn2listByZone []eqdbobject.Spawn2

	// range over spawn2 object, identifying matching zones
	for _, spawn2 := range *s2 {
		if spawn2.Zone == z {
			spawn2listByZone = append(spawn2listByZone, spawn2)
		}
	}
	return &spawn2listByZone
}
