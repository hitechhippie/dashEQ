package eqdbobject

import (
	"fmt"
	"strings"
)

func (b EQEmuWeirdBool) Bool() bool {
	switch b {
	case 0:
		return true
	default:
		return false
	}
}

func (w EQweight) String() string {
	return fmt.Sprintf("%.1f", (w / 10))
}

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

func (s ItemSlotsBitmask) String() string {
	if s == slotMax || s == 0 {
		return "ALL"
	}

	switch s {
	case Charm:
		return "Charm"
	case Ear_1:
		return "Ear"
	case Head:
		return "Head"
	case Face:
		return "Face"
	case Ear_2:
		return "Ear"
	case Neck:
		return "Neck"
	case Shoulder:
		return "Shoulder"
	case Arms:
		return "Arms"
	case Back:
		return "Back"
	case Bracer_1:
		return "Bracer"
	case Bracer_2:
		return "Bracer"
	case Range:
		return "Range"
	case Hands:
		return "Hands"
	case Primary:
		return "Primary"
	case Secondary:
		return "Secondary"
	case Ring_1:
		return "Ring"
	case Ring_2:
		return "Ring"
	case Chest:
		return "Chest"
	case Legs:
		return "Legs"
	case Feet:
		return "Feet"
	case Waist:
		return "Waist"
	case Powersource:
		return "Powersource"
	case Ammo:
		return "Ammo"
	}

	// multiple races
	var slots []string
	for slot := Charm; slot < slotMax; slot <<= 1 {
		if s&slot != 0 {
			slots = append(slots, slot.String())
		}
	}
	return strings.Join(slots, " | ")
}

func (r RaceListBitmask) String() string {

	if r == raceMax || r == 0 {
		return "ALL"
	}

	switch r {
	case HUM:
		return "HUM"
	case BAR:
		return "BAR"
	case ERU:
		return "ERU"
	case ELF:
		return "ELF"
	case HIE:
		return "HIE"
	case DEF:
		return "DEF"
	case HEF:
		return "HEF"
	case DWF:
		return "DWF"
	case TRL:
		return "TRL"
	case OGR:
		return "OGR"
	case HFL:
		return "HFL"
	case GNM:
		return "GNM"
	case IKS:
		return "IKS"
	case VAH:
		return "VAH"
	case FRG:
		return "FRG"
	case DRA:
		return "DRA"
	}

	// multiple races
	var races []string
	for race := HUM; race < raceMax; race <<= 1 {
		if r&race != 0 {
			races = append(races, race.String())
		}
	}
	return strings.Join(races, " | ")
}

func (c ClassListBitmask) String() string {

	if c == classMax || c == 0 {
		return "ALL"
	}

	switch c {
	case WAR:
		return "WAR"
	case CLR:
		return "CLR"
	case PAL:
		return "PAL"
	case RNG:
		return "RNG"
	case SHD:
		return "SHD"
	case DRU:
		return "DRU"
	case MNK:
		return "MNK"
	case BRD:
		return "BRD"
	case ROG:
		return "ROG"
	case SHM:
		return "SHM"
	case NEC:
		return "NEC"
	case WIZ:
		return "WIZ"
	case MAG:
		return "MAG"
	case ENC:
		return "ENC"
	case BST:
		return "BST"
	case BER:
		return "BER"
	}

	// multiple races
	var classes []string
	for class := WAR; class < classMax; class <<= 1 {
		if c&class != 0 {
			classes = append(classes, class.String())
		}
	}
	return strings.Join(classes, " | ")
}

func (s ItemSize) String() string {
	switch s {
	case 0:
		return "TINY"
	case 1:
		return "SMALL"
	case 2:
		return "MEDIUM"
	case 3:
		return "LARGE"
	case 4:
		return "GIANT"
	case 5:
		return "GIGANTIC"
	default:
		return "UNKNOWN"
	}
}

/*
+----------------+------------------+------+-----+---------------------+-------+
| Field          | Type             | Null | Key | Default             | Extra |
+----------------+------------------+------+-----+---------------------+-------+
| id             | int(11)          | NO   | PRI | 0                   |       |
| minstatus      | smallint(5)      | NO   | MUL | 0                   |       |
| Name           | varchar(64)      | NO   | MUL |                     |       |
| aagi           | int(11)          | NO   |     | 0                   |       |
| ac             | int(11)          | NO   |     | 0                   |       |
| acha           | int(11)          | NO   |     | 0                   |       |
| adex           | int(11)          | NO   |     | 0                   |       |
| aint           | int(11)          | NO   |     | 0                   |       |
| asta           | int(11)          | NO   |     | 0                   |       |
| astr           | int(11)          | NO   |     | 0                   |       |
| awis           | int(11)          | NO   |     | 0                   |       |
| bagsize        | int(11)          | NO   |     | 0                   |       |
| bagslots       | int(11)          | NO   |     | 0                   |       |
| bagtype        | int(11)          | NO   |     | 0                   |       |
| bagwr          | int(11)          | NO   |     | 0                   |       |
| banedmgamt     | int(11)          | NO   |     | 0                   |       |
| banedmgbody    | int(11)          | NO   |     | 0                   |       |
| banedmgrace    | int(11)          | NO   |     | 0                   |       |
| bardtype       | int(11)          | NO   |     | 0                   |       |
| bardvalue      | int(11)          | NO   |     | 0                   |       |
| book           | int(11)          | NO   |     | 0                   |       |
| casttime       | int(11)          | NO   |     | 0                   |       |
| casttime_      | int(11)          | NO   |     | 0                   |       |
| classes        | int(11)          | NO   |     | 0                   |       |
| color          | int(10) unsigned | NO   |     | 0                   |       |
| price          | int(11)          | NO   |     | 0                   |       |
| cr             | int(11)          | NO   |     | 0                   |       |
| damage         | int(11)          | NO   |     | 0                   |       |
| deity          | int(11)          | NO   |     | 0                   |       |
| delay          | int(11)          | NO   |     | 0                   |       |
| dr             | int(11)          | NO   |     | 0                   |       |
| clicktype      | int(11)          | NO   |     | 0                   |       |
| clicklevel2    | int(11)          | NO   |     | 0                   |       |
| elemdmgtype    | int(11)          | NO   |     | 0                   |       |
| elemdmgamt     | int(11)          | NO   |     | 0                   |       |
| factionamt1    | int(11)          | NO   |     | 0                   |       |
| factionamt2    | int(11)          | NO   |     | 0                   |       |
| factionamt3    | int(11)          | NO   |     | 0                   |       |
| factionamt4    | int(11)          | NO   |     | 0                   |       |
| factionmod1    | int(11)          | NO   |     | 0                   |       |
| factionmod2    | int(11)          | NO   |     | 0                   |       |
| factionmod3    | int(11)          | NO   |     | 0                   |       |
| factionmod4    | int(11)          | NO   |     | 0                   |       |
| filename       | varchar(32)      | NO   |     |                     |       |
| focuseffect    | int(11)          | NO   |     | 0                   |       |
| fr             | int(11)          | NO   |     | 0                   |       |
| fvnodrop       | int(11)          | NO   |     | 0                   |       |
| clicklevel     | int(11)          | NO   |     | 0                   |       |
| hp             | int(11)          | NO   |     | 0                   |       |
| icon           | int(11)          | NO   |     | 0                   |       |
| idfile         | varchar(30)      | NO   |     |                     |       |
| itemclass      | int(11)          | NO   |     | 0                   |       |
| itemtype       | int(11)          | NO   |     | 0                   |       |
| light          | int(11)          | NO   |     | 0                   |       |
| lore           | varchar(80)      | NO   | MUL |                     |       |
| magic          | int(11)          | NO   |     | 0                   |       |
| mana           | int(11)          | NO   |     | 0                   |       |
| material       | int(11)          | NO   |     | 0                   |       |
| maxcharges     | int(11)          | NO   |     | 0                   |       |
| mr             | int(11)          | NO   |     | 0                   |       |
| nodrop         | int(11)          | NO   |     | 0                   |       |
| norent         | int(11)          | NO   |     | 0                   |       |
| pr             | int(11)          | NO   |     | 0                   |       |
| procrate       | int(11)          | NO   |     | 0                   |       |
| races          | int(11)          | NO   |     | 0                   |       |
| range          | int(11)          | NO   |     | 0                   |       |
| reclevel       | int(11)          | NO   |     | 0                   |       |
| recskill       | int(11)          | NO   |     | 0                   |       |
| reqlevel       | int(11)          | NO   |     | 0                   |       |
| sellrate       | float            | NO   |     | 0                   |       |
| size           | int(11)          | NO   |     | 0                   |       |
| skillmodtype   | int(11)          | NO   |     | 0                   |       |
| skillmodvalue  | int(11)          | NO   |     | 0                   |       |
| slots          | int(11)          | NO   |     | 0                   |       |
| clickeffect    | int(11)          | NO   |     | 0                   |       |
| tradeskills    | int(11)          | NO   |     | 0                   |       |
| weight         | int(11)          | NO   |     | 0                   |       |
| booktype       | int(11)          | NO   |     | 0                   |       |
| recastdelay    | int(11)          | NO   |     | 0                   |       |
| recasttype     | int(11)          | NO   |     | 0                   |       |
| updated        | datetime         | NO   |     | 0000-00-00 00:00:00 |       |
| comment        | varchar(255)     | NO   |     |                     |       |
| stacksize      | int(11)          | NO   |     | 0                   |       |
| stackable      | int(11)          | NO   |     | 0                   |       |
| proceffect     | int(11)          | NO   |     | 0                   |       |
| proctype       | int(11)          | NO   |     | 0                   |       |
| proclevel2     | int(11)          | NO   |     | 0                   |       |
| proclevel      | int(11)          | NO   |     | 0                   |       |
| worneffect     | int(11)          | NO   |     | 0                   |       |
| worntype       | int(11)          | NO   |     | 0                   |       |
| wornlevel2     | int(11)          | NO   |     | 0                   |       |
| wornlevel      | int(11)          | NO   |     | 0                   |       |
| focustype      | int(11)          | NO   |     | 0                   |       |
| focuslevel2    | int(11)          | NO   |     | 0                   |       |
| focuslevel     | int(11)          | NO   |     | 0                   |       |
| scrolleffect   | int(11)          | NO   |     | 0                   |       |
| scrolltype     | int(11)          | NO   |     | 0                   |       |
| scrolllevel2   | int(11)          | NO   |     | 0                   |       |
| scrolllevel    | int(11)          | NO   |     | 0                   |       |
| serialized     | datetime         | YES  |     | NULL                |       |
| verified       | datetime         | YES  |     | NULL                |       |
| serialization  | text             | YES  |     | NULL                |       |
| source         | varchar(20)      | NO   |     |                     |       |
| lorefile       | varchar(32)      | NO   |     |                     |       |
| questitemflag  | int(11)          | NO   |     | 0                   |       |
| clickunk5      | int(11)          | NO   |     | 0                   |       |
| clickunk6      | varchar(32)      | NO   |     |                     |       |
| clickunk7      | int(11)          | NO   |     | 0                   |       |
| procunk1       | int(11)          | NO   |     | 0                   |       |
| procunk2       | int(11)          | NO   |     | 0                   |       |
| procunk3       | int(11)          | NO   |     | 0                   |       |
| procunk4       | int(11)          | NO   |     | 0                   |       |
| procunk6       | varchar(32)      | NO   |     |                     |       |
| procunk7       | int(11)          | NO   |     | 0                   |       |
| wornunk1       | int(11)          | NO   |     | 0                   |       |
| wornunk2       | int(11)          | NO   |     | 0                   |       |
| wornunk3       | int(11)          | NO   |     | 0                   |       |
| wornunk4       | int(11)          | NO   |     | 0                   |       |
| wornunk5       | int(11)          | NO   |     | 0                   |       |
| wornunk6       | varchar(32)      | NO   |     |                     |       |
| wornunk7       | int(11)          | NO   |     | 0                   |       |
| focusunk1      | int(11)          | NO   |     | 0                   |       |
| focusunk2      | int(11)          | NO   |     | 0                   |       |
| focusunk3      | int(11)          | NO   |     | 0                   |       |
| focusunk4      | int(11)          | NO   |     | 0                   |       |
| focusunk5      | int(11)          | NO   |     | 0                   |       |
| focusunk6      | varchar(32)      | NO   |     |                     |       |
| focusunk7      | int(11)          | NO   |     | 0                   |       |
| scrollunk1     | int(11)          | NO   |     | 0                   |       |
| scrollunk2     | int(11)          | NO   |     | 0                   |       |
| scrollunk3     | int(11)          | NO   |     | 0                   |       |
| scrollunk4     | int(11)          | NO   |     | 0                   |       |
| scrollunk5     | int(11)          | NO   |     | 0                   |       |
| scrollunk6     | varchar(32)      | NO   |     |                     |       |
| scrollunk7     | int(11)          | NO   |     | 0                   |       |
| clickname      | varchar(64)      | NO   |     |                     |       |
| procname       | varchar(64)      | NO   |     |                     |       |
| wornname       | varchar(64)      | NO   |     |                     |       |
| focusname      | varchar(64)      | NO   |     |                     |       |
| scrollname     | varchar(64)      | NO   |     |                     |       |
| created        | varchar(64)      | NO   |     |                     |       |
| bardeffect     | smallint(6)      | NO   |     | 0                   |       |
| bardeffecttype | smallint(6)      | NO   |     | 0                   |       |
| bardlevel2     | smallint(6)      | NO   |     | 0                   |       |
| bardlevel      | smallint(6)      | NO   |     | 0                   |       |
| bardunk1       | smallint(6)      | NO   |     | 0                   |       |
| bardunk2       | smallint(6)      | NO   |     | 0                   |       |
| bardunk3       | smallint(6)      | NO   |     | 0                   |       |
| bardunk4       | smallint(6)      | NO   |     | 0                   |       |
| bardunk5       | smallint(6)      | NO   |     | 0                   |       |
| bardname       | varchar(64)      | NO   |     |                     |       |
| bardunk7       | smallint(6)      | NO   |     | 0                   |       |
| gmflag         | tinyint(4)       | NO   |     | 0                   |       |
| soulbound      | tinyint(4)       | NO   |     | 0                   |       |
+----------------+------------------+------+-----+---------------------+-------+
*/

/* Races bitmasking
ID	Bitmask	Name
1	1		Human
2	2		Barbarian
3	4		Erudite
4	8		Wood Elf
5	16		High Elf
6	32		Dark Elf
7	64		Half Elf
8	128		Dwarf
9	256		Troll
10	512		Ogre
11	1024	Halfling
12	2048	Gnome
128	4096	Iksar
130	8192	Vah Shir
330	16384	Froglok
522	32768	Drakkin */
