package eqdbobject

import (
	"dasheq/internal/config"
	"dasheq/internal/eqdb"
	"dasheq/internal/logging"
	"errors"
)

func LoadDataZone(c *eqdb.Connection) (*[]Zone, error) {
	// initialize our return objects
	zoneSet := make([]Zone, 0)

	// gather the raw query data for zones
	zoneRows, err := c.Target.Query("SELECT " +
		"id," + //Id uint16
		"short_name," + //Short_name string
		"long_name," + //Long_name  string
		"castoutdoor," + //Outdoor  bool
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

func LoadDataSkill(c *eqdb.Connection, e *config.ServerConfig) (*[]Skill, error) {
	// initialize our return objects
	skillSet := make([]Skill, 0)

	skillRows, err := c.Target.Query("SELECT " +
		"skillID," + //Id uint32
		"class," + //Name string
		"level, " + //uint8
		"cap " + //uint8
		"FROM skill_caps")

	if err != nil {
		return nil, err
	}

	for skillRows.Next() {
		skill := new(Skill)
		skillRows.Scan(&skill.SkillID, &skill.Class, &skill.Level, &skill.Cap)
		skillSet = append(skillSet, *skill)
	}

	return &skillSet, nil
}

func LoadDataSpell(c *eqdb.Connection, e *config.ServerConfig) (*[]Spell, error) {
	// initialize our return objects
	spellSet := make([]Spell, 0)

	// gather the raw query data for npcs
	// separate routines for mainline eqemu vs eqmacemu for berserker class
	if e.EQemuEra == "eqmacemu" {
		spellRows, err := c.Target.Query("SELECT " +
			"id," + //Id uint32
			"name," + //Name string
			"classes1, " + //uint8
			"classes2, " + //uint8
			"classes3, " + //uint8
			"classes4, " + //uint8
			"classes5, " + //uint8
			"classes6, " + //uint8
			"classes7, " + //uint8
			"classes8, " + //uint8
			"classes9, " + //uint8
			"classes10, " + //uint8
			"classes11, " + //uint8
			"classes12, " + //uint8
			"classes13, " + //uint8
			"classes14, " + //uint8
			"classes15, " + //uint8
			"new_icon " + //uint8
			"FROM spells_new")

		if err != nil {
			return nil, err
		}

		for spellRows.Next() {
			spell := new(Spell)
			spellRows.Scan(&spell.Id, &spell.Name, &spell.Classes1, &spell.Classes2, &spell.Classes3, &spell.Classes4, &spell.Classes5, &spell.Classes6, &spell.Classes7, &spell.Classes8, &spell.Classes9, &spell.Classes10, &spell.Classes11, &spell.Classes12, &spell.Classes13, &spell.Classes14, &spell.Classes15, &spell.NewIcon)
			spellSet = append(spellSet, *spell)
		}

	} else if e.EQemuEra == "eqemu" {
		spellRows, err := c.Target.Query("SELECT " +
			"id," + //Id uint32
			"name," + //Name string
			"classes1, " + //uint8
			"classes2, " + //uint8
			"classes3, " + //uint8
			"classes4, " + //uint8
			"classes5, " + //uint8
			"classes6, " + //uint8
			"classes7, " + //uint8
			"classes8, " + //uint8
			"classes9, " + //uint8
			"classes10, " + //uint8
			"classes11, " + //uint8
			"classes12, " + //uint8
			"classes13, " + //uint8
			"classes14, " + //uint8
			"classes15, " + //uint8
			"classes16, " + //uint8
			"new_icon " + //uint8
			"FROM spells_new")

		if err != nil {
			return nil, err
		}

		for spellRows.Next() {
			spell := new(Spell)
			spellRows.Scan(&spell.Id, &spell.Name, &spell.Classes1, &spell.Classes2, &spell.Classes3, &spell.Classes4, &spell.Classes5, &spell.Classes6, &spell.Classes7, &spell.Classes8, &spell.Classes9, &spell.Classes10, &spell.Classes11, &spell.Classes12, &spell.Classes13, &spell.Classes14, &spell.Classes15, &spell.Classes16, &spell.NewIcon)
			spellSet = append(spellSet, *spell)
		}

	} else {
		return nil, errors.New("invalid EQemuEra configured")
	}

	return &spellSet, nil
}

func LoadDataItem(c *eqdb.Connection) (*[]Item, error) {
	logOb := logging.InitLogger("eqobject-item")

	// initialize our return objects
	itemSet := make([]Item, 0)

	// gather the raw query data for zones
	itemRows, err := c.Target.Query("SELECT " +
		"id," + //uint16
		"Name," + //string
		"ac," + //int
		"hp," + //int
		"icon," + //uint32
		"itemtype," + //uint32
		"magic," + //bool
		"mana," + //int
		"nodrop," + //int
		"norent," + //int
		"classes," + //uint32
		"damage," + //uint16
		"delay," + //uint16
		"races, " + //uint32
		"reclevel, " + //uint8
		"reqlevel, " + //uint8
		"scrolleffect, " + //uint32
		"size, " + //uint8
		"slots, " + //uint32
		"weight " + //float32
		"FROM items")
	if err != nil {
		return nil, err
	}

	// iterate through the raw zone query data and populate the return object
	for itemRows.Next() {
		item := new(Item)
		err = itemRows.Scan(
			&item.Id,
			&item.Name,
			&item.AC,
			&item.HP,
			&item.Icon,
			&item.Itemtype,
			&item.Magic,
			&item.Mana,
			&item.DBnodrop,
			&item.DBnorent,
			&item.Classes,
			&item.Damage,
			&item.Delay,
			&item.Races,
			&item.Reclevel,
			&item.Reqlevel,
			&item.ScrollEffect,
			&item.Size,
			&item.Slots,
			&item.Weight)
		if err != nil {
			logOb.Println(err)
		}
		item.Nodrop = item.DBnodrop.Bool()
		item.Norent = item.DBnorent.Bool()
		itemSet = append(itemSet, *item)
	}

	return &itemSet, nil
}
