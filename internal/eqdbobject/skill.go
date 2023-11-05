package eqdbobject

func (c ClassID) String() string {
	switch c {
	case 1:
		return "WAR"
	case 2:
		return "CLR"
	case 3:
		return "PAL"
	case 4:
		return "RNG"
	case 5:
		return "SHD"
	case 6:
		return "DRU"
	case 7:
		return "MNK"
	case 8:
		return "BRD"
	case 9:
		return "ROG"
	case 10:
		return "SHM"
	case 11:
		return "NEC"
	case 12:
		return "WIZ"
	case 13:
		return "MAG"
	case 14:
		return "ENC"
	case 15:
		return "BST"
	case 16:
		return "BER"
	default:
		return "NA"
	}
}

func (s SkillID) String() string {
	switch s {
	case 0:
		return "1H Blunt"
	case 1:
		return "1H Slashing"
	case 2:
		return "2H Blunt"
	case 3:
		return "2H Slashing"
	case 4:
		return "Abjuration"
	case 5:
		return "Alteration"
	case 6:
		return "Apply Poison"
	case 7:
		return "Archery"
	case 8:
		return "Backstab"
	case 9:
		return "Bind Wound"
	case 10:
		return "Bash"
	case 11:
		return "Block"
	case 12:
		return "Brass Instruments"
	case 13:
		return "Channeling"
	case 14:
		return "Conjuration"
	case 15:
		return "Defense"
	case 16:
		return "Disarm"
	case 17:
		return "Disarm Traps"
	case 18:
		return "Divination"
	case 19:
		return "Dodge"
	case 20:
		return "Double Attack"
	case 21:
		return "Dragon Punch"
	case 22:
		return "Dual Wield"
	case 23:
		return "Eagle Strike"
	case 24:
		return "Evocation"
	case 25:
		return "Feign Death"
	case 26:
		return "Flying Kick"
	case 27:
		return "Forage"
	case 28:
		return "Hand to Hand"
	case 29:
		return "Hide"
	case 30:
		return "Kick"
	case 31:
		return "Meditate"
	case 32:
		return "Mend"
	case 33:
		return "Offense"
	case 34:
		return "Parry"
	case 35:
		return "Pick Lock"
	case 36:
		return "1H Piercing"
	case 37:
		return "Riposte"
	case 38:
		return "Round Kick"
	case 39:
		return "Safe Fall"
	case 40:
		return "Sense Heading"
	case 41:
		return "Singing"
	case 42:
		return "Sneak"
	case 43:
		return "Specialize Abjure"
	case 44:
		return "Specialize Alteration"
	case 45:
		return "Specialize Conjuration"
	case 46:
		return "Specialize Divination"
	case 47:
		return "Specialize Evocation"
	case 48:
		return "Pick Pockets"
	case 49:
		return "Stringed Instruments"
	case 50:
		return "Swimming"
	case 51:
		return "Throwing"
	case 52:
		return "Tiger Claw"
	case 53:
		return "Tracking"
	case 54:
		return "Wind Instruments"
	case 55:
		return "Fishing"
	case 56:
		return "Make Poison"
	case 57:
		return "Tinkering"
	case 58:
		return "Research"
	case 59:
		return "Alchemy"
	case 60:
		return "Baking"
	case 61:
		return "Tailoring"
	case 62:
		return "Sense Traps"
	case 63:
		return "Blacksmithing"
	case 64:
		return "Fletching"
	case 65:
		return "Brewing"
	case 66:
		return "Alcohol Tolerance"
	case 67:
		return "Begging"
	case 68:
		return "Jewelry Making"
	case 69:
		return "Pottery"
	case 70:
		return "Percussion Instruments"
	case 71:
		return "Intimidation"
	case 72:
		return "Berserking"
	case 73:
		return "Taunt"
	case 74:
		return "Frenzy"
	case 75:
		return "Remove Trap"
	case 76:
		return "Triple Attack"
	case 77:
		return "2H Piercing"
	default:
		return "N/A"
	}
}

/*
+---------+-----------------------+------+-----+---------+-------+
| Field   | Type                  | Null | Key | Default | Extra |
+---------+-----------------------+------+-----+---------+-------+
| skillID | tinyint(3) unsigned   | NO   | PRI | 0       |       |
| class   | tinyint(3) unsigned   | NO   | PRI | 0       |       |
| level   | tinyint(3) unsigned   | NO   | PRI | 0       |       |
| cap     | mediumint(8) unsigned | NO   |     | 0       |       |
| class_  | tinyint(3) unsigned   | NO   | PRI | 0       |       |
+---------+-----------------------+------+-----+---------+-------+
*/
