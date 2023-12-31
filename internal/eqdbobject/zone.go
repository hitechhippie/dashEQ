package eqdbobject

import (
	"errors"
	"strconv"
	"strings"
)

func (e ZoneExpansion) String() string {
	switch e {
	case 1:
		return "Classic"
	case 2:
		return "The Ruins of Kunark"
	case 3:
		return "The Scars of Velious"
	case 4:
		return "The Shadows of Luclin"
	case 5:
		return "The Planes of Power"
	case 6:
		return "The Legacy of Ykesha"
	case 7:
		return "Lost Dungeons of Norrath"
	case 8:
		return "Gates of Discord"
	case 9:
		return "Omens of War"
	case 10:
		return "Dragons of Norrath"
	case 11:
		return "Depths of Darkhollow"
	case 12:
		return "Prophecy of Ro"
	case 13:
		return "The Serpent's Spine"
	case 14:
		return "The Buried Sea"
	case 15:
		return "Secrets of Faydwer"
	case 16:
		return "Seeds of Destruction"
	case 17:
		return "Underfoot"
	case 18:
		return "House of Thule"
	case 19:
		return "Veil of Alaris"
	case 20:
		return "Rain of Fear"
	case 21:
		return "Call of the Forsaken"
	case 22:
		return "The Darkened Sea"
	case 23:
		return "The Broken Mirror"
	case 24:
		return "Empires of Kunark"
	case 25:
		return "Ring of Scale"
	case 26:
		return "The Burning Lands"
	case 27:
		return "Torment of Velious"
	default:
		return "N/A"
	}
}

// s = Zone name string determined from current directory
// d = the populated NPC slice with all the NPC data
func ZoneIdLookup(s string, z *[]Zone) (uint32, error) {
	for _, d := range *z {
		if strings.Compare(s, d.Short_name) == 0 {
			return d.Id, nil
		}
	}
	err := errors.New("Zone ID not found in Zone data set: " + s)
	return 0, err
}

// i = given zone id for lookup
// z = populated Zone struct from requestor
func ZoneNameLookup(i uint32, z *[]Zone) (string, error) {
	for _, d := range *z {
		if d.Id == i {
			return d.Long_name, nil
		}
	}
	err := errors.New("Zone name not found in Zone data set: " + strconv.FormatUint(uint64(i), 10))
	return "", err
}

/*
+---------------------------+------------------------------+------+-----+---------+----------------+
| Field                     | Type                         | Null | Key | Default | Extra          |
+---------------------------+------------------------------+------+-----+---------+----------------+
| short_name                | varchar(32)                  | YES  | MUL | NULL    |                |
| id                        | int(10)                      | NO   | PRI | NULL    | auto_increment |
| file_name                 | varchar(16)                  | YES  |     | NULL    |                |
| long_name                 | text                         | NO   |     | NULL    |                |
| map_file_name             | varchar(100)                 | YES  |     | NULL    |                |
| safe_x                    | float                        | NO   |     | 0       |                |
| safe_y                    | float                        | NO   |     | 0       |                |
| safe_z                    | float                        | NO   |     | 0       |                |
| safe_heading              | float                        | NO   |     | 0       |                |
| graveyard_id              | float                        | NO   |     | 0       |                |
| min_level                 | tinyint(3) unsigned          | NO   |     | 0       |                |
| min_status                | tinyint(3) unsigned          | NO   |     | 0       |                |
| zoneidnumber              | int(4)                       | NO   | MUL | 0       |                |
| version                   | tinyint(3) unsigned          | NO   |     | 0       |                |
| timezone                  | int(5)                       | NO   |     | 0       |                |
| maxclients                | int(5)                       | NO   |     | 0       |                |
| ruleset                   | int(10) unsigned             | NO   |     | 0       |                |
| note                      | varchar(80)                  | YES  |     | NULL    |                |
| underworld                | float                        | NO   |     | 0       |                |
| minclip                   | float                        | NO   |     | 450     |                |
| maxclip                   | float                        | NO   |     | 450     |                |
| fog_minclip               | float                        | NO   |     | 450     |                |
| fog_maxclip               | float                        | NO   |     | 450     |                |
| fog_blue                  | tinyint(3) unsigned          | NO   |     | 0       |                |
| fog_red                   | tinyint(3) unsigned          | NO   |     | 0       |                |
| fog_green                 | tinyint(3) unsigned          | NO   |     | 0       |                |
| sky                       | tinyint(3) unsigned          | NO   |     | 1       |                |
| ztype                     | tinyint(3) unsigned          | NO   |     | 1       |                |
| zone_exp_multiplier       | decimal(6,2)                 | NO   |     | 0.00    |                |
| walkspeed                 | float                        | NO   |     | 0.4     |                |
| time_type                 | tinyint(3) unsigned          | NO   |     | 2       |                |
| fog_red1                  | tinyint(3) unsigned          | NO   |     | 0       |                |
| fog_green1                | tinyint(3) unsigned          | NO   |     | 0       |                |
| fog_blue1                 | tinyint(3) unsigned          | NO   |     | 0       |                |
| fog_minclip1              | float                        | NO   |     | 450     |                |
| fog_maxclip1              | float                        | NO   |     | 450     |                |
| fog_red2                  | tinyint(3) unsigned          | NO   |     | 0       |                |
| fog_green2                | tinyint(3) unsigned          | NO   |     | 0       |                |
| fog_blue2                 | tinyint(3) unsigned          | NO   |     | 0       |                |
| fog_minclip2              | float                        | NO   |     | 450     |                |
| fog_maxclip2              | float                        | NO   |     | 450     |                |
| fog_red3                  | tinyint(3) unsigned          | NO   |     | 0       |                |
| fog_green3                | tinyint(3) unsigned          | NO   |     | 0       |                |
| fog_blue3                 | tinyint(3) unsigned          | NO   |     | 0       |                |
| fog_minclip3              | float                        | NO   |     | 450     |                |
| fog_maxclip3              | float                        | NO   |     | 450     |                |
| fog_red4                  | tinyint(3) unsigned          | NO   |     | 0       |                |
| fog_green4                | tinyint(3) unsigned          | NO   |     | 0       |                |
| fog_blue4                 | tinyint(3) unsigned          | NO   |     | 0       |                |
| fog_minclip4              | float                        | NO   |     | 450     |                |
| fog_maxclip4              | float                        | NO   |     | 450     |                |
| fog_density               | float                        | NO   |     | 0       |                |
| flag_needed               | varchar(128)                 | NO   |     |         |                |
| canbind                   | tinyint(4)                   | NO   |     | 1       |                |
| cancombat                 | tinyint(4)                   | NO   |     | 1       |                |
| canlevitate               | tinyint(4)                   | NO   |     | 1       |                |
| castoutdoor               | tinyint(4)                   | NO   |     | 1       |                |
| hotzone                   | tinyint(3) unsigned          | NO   |     | 0       |                |
| insttype                  | tinyint(1) unsigned zerofill | NO   |     | 0       |                |
| shutdowndelay             | bigint(16) unsigned          | NO   |     | 5000    |                |
| peqzone                   | tinyint(4)                   | NO   |     | 1       |                |
| expansion                 | tinyint(3)                   | NO   |     | 0       |                |
| suspendbuffs              | tinyint(1) unsigned          | NO   |     | 0       |                |
| rain_chance1              | int(4)                       | NO   |     | 0       |                |
| rain_chance2              | int(4)                       | NO   |     | 0       |                |
| rain_chance3              | int(4)                       | NO   |     | 0       |                |
| rain_chance4              | int(4)                       | NO   |     | 0       |                |
| rain_duration1            | int(4)                       | NO   |     | 0       |                |
| rain_duration2            | int(4)                       | NO   |     | 0       |                |
| rain_duration3            | int(4)                       | NO   |     | 0       |                |
| rain_duration4            | int(4)                       | NO   |     | 0       |                |
| snow_chance1              | int(4)                       | NO   |     | 0       |                |
| snow_chance2              | int(4)                       | NO   |     | 0       |                |
| snow_chance3              | int(4)                       | NO   |     | 0       |                |
| snow_chance4              | int(4)                       | NO   |     | 0       |                |
| snow_duration1            | int(4)                       | NO   |     | 0       |                |
| snow_duration2            | int(4)                       | NO   |     | 0       |                |
| snow_duration3            | int(4)                       | NO   |     | 0       |                |
| snow_duration4            | int(4)                       | NO   |     | 0       |                |
| gravity                   | float                        | NO   |     | 0.4     |                |
| type                      | int(3)                       | NO   |     | 0       |                |
| skylock                   | tinyint(4)                   | NO   |     | 0       |                |
| fast_regen_hp             | int(11)                      | NO   |     | 180     |                |
| fast_regen_mana           | int(11)                      | NO   |     | 180     |                |
| fast_regen_endurance      | int(11)                      | NO   |     | 180     |                |
| npc_max_aggro_dist        | int(11)                      | NO   |     | 600     |                |
| max_movement_update_range | int(11) unsigned             | NO   |     | 600     |                |
| min_expansion             | tinyint(4)                   | NO   |     | -1      |                |
| max_expansion             | tinyint(4)                   | NO   |     | -1      |                |
| content_flags             | varchar(100)                 | YES  |     | NULL    |                |
| content_flags_disabled    | varchar(100)                 | YES  |     | NULL    |                |
| underworld_teleport_index | int(4)                       | NO   |     | 0       |                |
| lava_damage               | int(11)                      | YES  |     | 50      |                |
| min_lava_damage           | int(11)                      | NO   |     | 10      |                |
+---------------------------+------------------------------+------+-----+---------+----------------+
*/
