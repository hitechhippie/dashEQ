package eqdbobject

type ItemTypes uint32

func (t ItemTypes) String() string {
	switch t {
	case 0:
		return "1H Slashing"
	case 1:
		return "2H Slashing"
	case 2:
		return "1H Piercing"
	case 3:
		return "1H Blunt"
	case 4:
		return "2H Blunt"
	case 5:
		return "Archery"
	case 7:
		return "Throwing"
	case 8:
		return "Shield"
	case 10:
		return "Armor"
	case 11:
		return "Tradeskill Items"
	case 12:
		return "Lockpicking"
	case 14:
		return "Food"
	case 15:
		return "Drink"
	case 16:
		return "Light Source"
	case 17:
		return "Common Inventory Item"
	case 18:
		return "Bind Wound"
	case 19:
		return "Thrown Casting Items"
	case 20:
		return "Spells / Song Sheets"
	case 21:
		return "Potions"
	case 22:
		return "Fletched Arrows"
	case 23:
		return "Wind Instruments"
	case 24:
		return "Stringed Instruments"
	case 25:
		return "Brass Instruments"
	case 26:
		return "Percussion Instruments"
	case 27:
		return "Ammo"
	case 29:
		return "Jewelry"
	case 31:
		return "Readable Notes and Scrolls"
	case 32:
		return "Readable Books"
	case 33:
		return "Keys"
	case 34:
		return "Odd Items"
	case 35:
		return "2H Piercing"
	case 36:
		return "Fishing Poles"
	case 37:
		return "Fishing Bait"
	case 38:
		return "Alcoholic Beverages"
	case 39:
		return "More Keys"
	case 40:
		return "Compasses"
	case 42:
		return "Poisons"
	case 45:
		return "Hand to Hand"
	case 52:
		return "Charms"
	case 53:
		return "Dyes"
	case 54:
		return "Augments"
	case 55:
		return "Augment Solvents"
	case 56:
		return "Augment Distillers"
	case 58:
		return "Fellowship Banner Materials"
	case 60:
		return "Cultural Armor Manuals"
	case 63:
		return "New Curencies like Orum"
	default:
		return "Unknown"
	}
}
