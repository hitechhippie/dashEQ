package eqdbtransformation

// create new spell based transform datasets

import "dasheq/internal/eqdbobject"

func PopulateSpellListByClass(s *[]SpellListByClass, c int, ss *[]eqdbobject.Spell, i *[]eqdbobject.Item) {
	switch c {
	case 1:
		for _, data := range *ss {
			if data.Classes1 < 255 {
				var spell SpellListByClass
				spell.Class = "Warrior"
				spell.Id = data.Id
				spell.Name = data.Name
				spell.Level = data.Classes1
				spell.NewIcon = data.NewIcon

				spell.Scroll = getSpellScrollID(data.Id, i)

				*s = append(*s, spell)
			}
		}
	case 2:
		for _, data := range *ss {
			if data.Classes2 < 255 {
				var spell SpellListByClass
				spell.Class = "Cleric"
				spell.Id = data.Id
				spell.Name = data.Name
				spell.Level = data.Classes2
				spell.NewIcon = data.NewIcon

				spell.Scroll = getSpellScrollID(data.Id, i)

				*s = append(*s, spell)
			}
		}
	case 3:
		for _, data := range *ss {
			if data.Classes3 < 255 {
				var spell SpellListByClass
				spell.Class = "Paladin"
				spell.Id = data.Id
				spell.Name = data.Name
				spell.Level = data.Classes3
				spell.NewIcon = data.NewIcon

				spell.Scroll = getSpellScrollID(data.Id, i)

				*s = append(*s, spell)
			}
		}
	case 4:
		for _, data := range *ss {
			if data.Classes4 < 255 {
				var spell SpellListByClass
				spell.Class = "Ranger"
				spell.Id = data.Id
				spell.Name = data.Name
				spell.Level = data.Classes4
				spell.NewIcon = data.NewIcon

				spell.Scroll = getSpellScrollID(data.Id, i)

				*s = append(*s, spell)
			}
		}
	case 5:
		for _, data := range *ss {
			if data.Classes5 < 255 {
				var spell SpellListByClass
				spell.Class = "Shadow Knight"
				spell.Id = data.Id
				spell.Name = data.Name
				spell.Level = data.Classes5
				spell.NewIcon = data.NewIcon

				spell.Scroll = getSpellScrollID(data.Id, i)

				*s = append(*s, spell)
			}
		}
	case 6:
		for _, data := range *ss {
			if data.Classes6 < 255 {
				var spell SpellListByClass
				spell.Class = "Druid"
				spell.Id = data.Id
				spell.Name = data.Name
				spell.Level = data.Classes6
				spell.NewIcon = data.NewIcon

				spell.Scroll = getSpellScrollID(data.Id, i)

				*s = append(*s, spell)
			}
		}
	case 7:
		for _, data := range *ss {
			if data.Classes7 < 255 {
				var spell SpellListByClass
				spell.Class = "Monk"
				spell.Id = data.Id
				spell.Name = data.Name
				spell.Level = data.Classes7
				spell.NewIcon = data.NewIcon

				spell.Scroll = getSpellScrollID(data.Id, i)

				*s = append(*s, spell)
			}
		}
	case 8:
		for _, data := range *ss {
			if data.Classes8 < 255 {
				var spell SpellListByClass
				spell.Class = "Bard"
				spell.Id = data.Id
				spell.Name = data.Name
				spell.Level = data.Classes8
				spell.NewIcon = data.NewIcon

				spell.Scroll = getSpellScrollID(data.Id, i)

				*s = append(*s, spell)
			}
		}
	case 9:
		for _, data := range *ss {
			if data.Classes9 < 255 {
				var spell SpellListByClass
				spell.Class = "Rogue"
				spell.Id = data.Id
				spell.Name = data.Name
				spell.Level = data.Classes9
				spell.NewIcon = data.NewIcon

				spell.Scroll = getSpellScrollID(data.Id, i)

				*s = append(*s, spell)
			}
		}
	case 10:
		for _, data := range *ss {
			if data.Classes10 < 255 {
				var spell SpellListByClass
				spell.Class = "Shaman"
				spell.Id = data.Id
				spell.Name = data.Name
				spell.Level = data.Classes10
				spell.NewIcon = data.NewIcon

				spell.Scroll = getSpellScrollID(data.Id, i)

				*s = append(*s, spell)
			}
		}
	case 11:
		for _, data := range *ss {
			if data.Classes11 < 255 {
				var spell SpellListByClass
				spell.Class = "Necromancer"
				spell.Id = data.Id
				spell.Name = data.Name
				spell.Level = data.Classes11
				spell.NewIcon = data.NewIcon

				spell.Scroll = getSpellScrollID(data.Id, i)

				*s = append(*s, spell)
			}
		}
	case 12:
		for _, data := range *ss {
			if data.Classes12 < 255 {
				var spell SpellListByClass
				spell.Class = "Wizard"
				spell.Id = data.Id
				spell.Name = data.Name
				spell.Level = data.Classes12
				spell.NewIcon = data.NewIcon

				spell.Scroll = getSpellScrollID(data.Id, i)

				*s = append(*s, spell)
			}
		}
	case 13:
		for _, data := range *ss {
			if data.Classes13 < 255 {
				var spell SpellListByClass
				spell.Class = "Magician"
				spell.Id = data.Id
				spell.Name = data.Name
				spell.Level = data.Classes13
				spell.NewIcon = data.NewIcon

				spell.Scroll = getSpellScrollID(data.Id, i)

				*s = append(*s, spell)
			}
		}
	case 14:
		for _, data := range *ss {
			if data.Classes14 < 255 {
				var spell SpellListByClass
				spell.Class = "Enchanter"
				spell.Id = data.Id
				spell.Name = data.Name
				spell.Level = data.Classes14
				spell.NewIcon = data.NewIcon

				spell.Scroll = getSpellScrollID(data.Id, i)

				*s = append(*s, spell)
			}
		}
	case 15:
		for _, data := range *ss {
			if data.Classes15 < 255 {
				var spell SpellListByClass
				spell.Class = "Beastlord"
				spell.Id = data.Id
				spell.Name = data.Name
				spell.Level = data.Classes15
				spell.NewIcon = data.NewIcon

				spell.Scroll = getSpellScrollID(data.Id, i)

				*s = append(*s, spell)
			}
		}
	case 16:
		for _, data := range *ss {
			if data.Classes16 < 255 {
				var spell SpellListByClass
				spell.Class = "Berserker"
				spell.Id = data.Id
				spell.Name = data.Name
				spell.Level = data.Classes16
				spell.NewIcon = data.NewIcon

				spell.Scroll = getSpellScrollID(data.Id, i)

				*s = append(*s, spell)
			}
		}
	}
}

func getSpellScrollID(si uint32, i *[]eqdbobject.Item) uint32 {
	for _, itemdata := range *i {
		if itemdata.ScrollEffect == si {
			return itemdata.Id
		}
	}
	return 0
}

// implementation for sort.Interface to sort spells by level
func (s SpellListByClassSorted) Len() int {
	return len(s)
}

// implementation for sort.Interface to sort spells by level
func (s SpellListByClassSorted) Less(i, j int) bool {
	return s[i].Level < s[j].Level
}

// implementation for sort.Interface to sort spells by level
func (s SpellListByClassSorted) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
