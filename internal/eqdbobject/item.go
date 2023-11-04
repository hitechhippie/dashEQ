package eqdbobject

import "fmt"

type Item struct {
	Id           uint32
	Name         string
	AC           int
	HP           int
	Icon         uint32
	Itemtype     ItemTypes
	Reclevel     uint8
	Reqlevel     uint8
	ScrollEffect uint32
	Size         ItemSize
	Slots        ItemSlots
	Weight       EQweight
	Magic        bool
	Mana         int
	DBnodrop     EQEmuWeirdBool
	DBnorent     EQEmuWeirdBool
	Classes      ClassList
	Damage       uint16
	Delay        uint16
	Races        RaceList
	Nodrop       bool
	Norent       bool
}

type EQEmuWeirdBool int
type EQweight float32

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
