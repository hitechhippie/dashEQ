package eqdbobject

import "dasheq/internal/eqdb"

func LoadDataZone(c *eqdb.Connection) (*[]Zone, error) {
	// initialize our return objects
	zoneSet := make([]Zone, 0)

	// gather the raw query data for zones
	zoneRows, err := c.Target.Query("SELECT " +
		"id," + //Id uint16
		"short_name," + //Short_name string
		"long_name," + //Long_name  string
		"castoutdoor," + //Outdoor    uint8
		"expansion " + //Expansion  uint8
		"FROM zone")
	if err != nil {
		return nil, err
	}

	// iterate through the raw zone query data and populate the return object
	for zoneRows.Next() {
		zone := new(Zone)
		zoneRows.Scan(&zone.Id, &zone.Short_name, &zone.Long_name, &zone.Outdoor, &zone.Expansion)
		zoneSet = append(zoneSet, *zone)
	}

	return &zoneSet, nil
}

func LoadDataNPC(c *eqdb.Connection) (*[]NPC, error) {
	// initialize our return objects
	npcSet := make([]NPC, 0)

	// gather the raw query data for npcs
	npcRows, err := c.Target.Query("SELECT " +
		"id," + //Id uint32
		"name," + //Name string
		"level, " + //Level uint8
		"race, " + //uint16
		"class, " + //uint8
		"hp, " + //uint32
		"mana, " + //uint32
		"loottable_id, " + //uint32
		"npc_spells_id, " + //uint16
		"npc_faction_id, " + //uint16
		"mindmg, " + //uint16
		"maxdmg, " + //uint16
		"attack_count, " + //uint16
		"runspeed, " + //float32
		"MR, " + //uint16
		"CR, " + //uint16
		"DR, " + //uint16
		"FR, " + //uint16
		"PR, " + //uint16
		"AC " + //uint16
		"FROM npc_types")
	if err != nil {
		return nil, err
	}

	// iterate through the raw npc query data and populate the return object
	for npcRows.Next() {
		npc := new(NPC)
		npcRows.Scan(&npc.Id, &npc.Name, &npc.Level, &npc.Race, &npc.Class, &npc.HP, &npc.Mana, &npc.LootTable, &npc.NPCSpells, &npc.NPCFaction, &npc.MinDmg, &npc.MaxDmg, &npc.AttackCount, &npc.RunSpeed, &npc.MR, &npc.CR, &npc.DR, &npc.FR, &npc.PR, &npc.AC)
		npcSet = append(npcSet, *npc)
	}
	return &npcSet, nil
}
