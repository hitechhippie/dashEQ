package eqdbtransformation

import "dasheq/internal/eqdbobject"

func PopulateZoneListByExpansion(z *[]ZoneListByExpansion, zz *[]eqdbobject.Zone) error {
	// range over main server zone struct to populate each ZoneListByExpansion entry
	for _, zones := range *zz {
		var zone ZoneListByExpansion
		zone.Expansion = zones.Expansion
		zone.Name = zones.Long_name
		zone.Short_name = zones.Short_name

		*z = append(*z, zone)
	}

	return nil
}

// implementation for sort.Interface to sort zones by expansion
func (z ZoneListByExpansionSorted) Len() int {
	return len(z)
}

// implementation for sort.Interface to sort zones by expansion
func (z ZoneListByExpansionSorted) Less(i, j int) bool {
	return z[i].Expansion < z[j].Expansion
}

// implementation for sort.Interface to sort zones by expansion
func (z ZoneListByExpansionSorted) Swap(i, j int) {
	z[i], z[j] = z[j], z[i]
}
