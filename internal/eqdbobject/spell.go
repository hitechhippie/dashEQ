package eqdbobject

// future spell type implementations placeholder

/*
+--------------------------+--------------+------+-----+------------+-------+
| Field                    | Type         | Null | Key | Default    | Extra |
+--------------------------+--------------+------+-----+------------+-------+
| id                       | int(11)      | NO   | PRI | 0          |       |
| name                     | varchar(64)  | YES  |     | NULL       |       |
| player_1                 | varchar(64)  | YES  |     | BLUE_TRAIL |       |
| teleport_zone            | varchar(64)  | YES  |     | NULL       |       |
| you_cast                 | varchar(120) | YES  |     | NULL       |       |
| other_casts              | varchar(120) | YES  |     | NULL       |       |
| cast_on_you              | varchar(120) | YES  |     | NULL       |       |
| cast_on_other            | varchar(120) | YES  |     | NULL       |       |
| spell_fades              | varchar(120) | YES  |     | NULL       |       |
| range                    | int(11)      | NO   |     | 100        |       |
| aoerange                 | int(11)      | NO   |     | 0          |       |
| pushback                 | float(14,2)  | NO   |     | 0.00       |       |
| pushup                   | float(14,2)  | NO   |     | 0.00       |       |
| cast_time                | int(11)      | NO   |     | 0          |       |
| recovery_time            | int(11)      | NO   |     | 0          |       |
| recast_time              | int(11)      | NO   |     | 0          |       |
| buffdurationformula      | int(11)      | NO   |     | 7          |       |
| buffduration             | int(11)      | NO   |     | 65         |       |
| AEDuration               | int(11)      | NO   |     | 0          |       |
| mana                     | int(11)      | NO   |     | 0          |       |
| effect_base_value1       | int(11)      | NO   |     | 100        |       |
| effect_base_value2       | int(11)      | NO   |     | 0          |       |
| effect_base_value3       | int(11)      | NO   |     | 0          |       |
| effect_base_value4       | int(11)      | NO   |     | 0          |       |
| effect_base_value5       | int(11)      | NO   |     | 0          |       |
| effect_base_value6       | int(11)      | NO   |     | 0          |       |
| effect_base_value7       | int(11)      | NO   |     | 0          |       |
| effect_base_value8       | int(11)      | NO   |     | 0          |       |
| effect_base_value9       | int(11)      | NO   |     | 0          |       |
| effect_base_value10      | int(11)      | NO   |     | 0          |       |
| effect_base_value11      | int(11)      | NO   |     | 0          |       |
| effect_base_value12      | int(11)      | NO   |     | 0          |       |
| effect_limit_value1      | int(11)      | NO   |     | 0          |       |
| effect_limit_value2      | int(11)      | NO   |     | 0          |       |
| effect_limit_value3      | int(11)      | NO   |     | 0          |       |
| effect_limit_value4      | int(11)      | NO   |     | 0          |       |
| effect_limit_value5      | int(11)      | NO   |     | 0          |       |
| effect_limit_value6      | int(11)      | NO   |     | 0          |       |
| effect_limit_value7      | int(11)      | NO   |     | 0          |       |
| effect_limit_value8      | int(11)      | NO   |     | 0          |       |
| effect_limit_value9      | int(11)      | NO   |     | 0          |       |
| effect_limit_value10     | int(11)      | NO   |     | 0          |       |
| effect_limit_value11     | int(11)      | NO   |     | 0          |       |
| effect_limit_value12     | int(11)      | NO   |     | 0          |       |
| max1                     | int(11)      | NO   |     | 0          |       |
| max2                     | int(11)      | NO   |     | 0          |       |
| max3                     | int(11)      | NO   |     | 0          |       |
| max4                     | int(11)      | NO   |     | 0          |       |
| max5                     | int(11)      | NO   |     | 0          |       |
| max6                     | int(11)      | NO   |     | 0          |       |
| max7                     | int(11)      | NO   |     | 0          |       |
| max8                     | int(11)      | NO   |     | 0          |       |
| max9                     | int(11)      | NO   |     | 0          |       |
| max10                    | int(11)      | NO   |     | 0          |       |
| max11                    | int(11)      | NO   |     | 0          |       |
| max12                    | int(11)      | NO   |     | 0          |       |
| icon                     | int(11)      | NO   |     | 0          |       |
| memicon                  | int(11)      | NO   |     | 0          |       |
| components1              | int(11)      | NO   |     | -1         |       |
| components2              | int(11)      | NO   |     | -1         |       |
| components3              | int(11)      | NO   |     | -1         |       |
| components4              | int(11)      | NO   |     | -1         |       |
| component_counts1        | int(11)      | NO   |     | 1          |       |
| component_counts2        | int(11)      | NO   |     | 1          |       |
| component_counts3        | int(11)      | NO   |     | 1          |       |
| component_counts4        | int(11)      | NO   |     | 1          |       |
| NoexpendReagent1         | int(11)      | NO   |     | -1         |       |
| NoexpendReagent2         | int(11)      | NO   |     | -1         |       |
| NoexpendReagent3         | int(11)      | NO   |     | -1         |       |
| NoexpendReagent4         | int(11)      | NO   |     | -1         |       |
| formula1                 | int(11)      | NO   |     | 100        |       |
| formula2                 | int(11)      | NO   |     | 100        |       |
| formula3                 | int(11)      | NO   |     | 100        |       |
| formula4                 | int(11)      | NO   |     | 100        |       |
| formula5                 | int(11)      | NO   |     | 100        |       |
| formula6                 | int(11)      | NO   |     | 100        |       |
| formula7                 | int(11)      | NO   |     | 100        |       |
| formula8                 | int(11)      | NO   |     | 100        |       |
| formula9                 | int(11)      | NO   |     | 100        |       |
| formula10                | int(11)      | NO   |     | 100        |       |
| formula11                | int(11)      | NO   |     | 100        |       |
| formula12                | int(11)      | NO   |     | 100        |       |
| LightType                | int(11)      | NO   |     | 0          |       |
| goodEffect               | int(11)      | NO   |     | 0          |       |
| Activated                | int(11)      | NO   |     | 0          |       |
| resisttype               | int(11)      | NO   |     | 0          |       |
| effectid1                | int(11)      | NO   |     | 254        |       |
| effectid2                | int(11)      | NO   |     | 254        |       |
| effectid3                | int(11)      | NO   |     | 254        |       |
| effectid4                | int(11)      | NO   |     | 254        |       |
| effectid5                | int(11)      | NO   |     | 254        |       |
| effectid6                | int(11)      | NO   |     | 254        |       |
| effectid7                | int(11)      | NO   |     | 254        |       |
| effectid8                | int(11)      | NO   |     | 254        |       |
| effectid9                | int(11)      | NO   |     | 254        |       |
| effectid10               | int(11)      | NO   |     | 254        |       |
| effectid11               | int(11)      | NO   |     | 254        |       |
| effectid12               | int(11)      | NO   |     | 254        |       |
| targettype               | int(11)      | NO   |     | 2          |       |
| basediff                 | int(11)      | NO   |     | 0          |       |
| skill                    | int(11)      | NO   |     | 98         |       |
| zonetype                 | int(11)      | NO   |     | -1         |       |
| EnvironmentType          | int(11)      | NO   |     | 0          |       |
| TimeOfDay                | int(11)      | NO   |     | 0          |       |
| classes1                 | int(11)      | NO   |     | 255        |       |
| classes2                 | int(11)      | NO   |     | 255        |       |
| classes3                 | int(11)      | NO   |     | 255        |       |
| classes4                 | int(11)      | NO   |     | 255        |       |
| classes5                 | int(11)      | NO   |     | 255        |       |
| classes6                 | int(11)      | NO   |     | 255        |       |
| classes7                 | int(11)      | NO   |     | 255        |       |
| classes8                 | int(11)      | NO   |     | 255        |       |
| classes9                 | int(11)      | NO   |     | 255        |       |
| classes10                | int(11)      | NO   |     | 255        |       |
| classes11                | int(11)      | NO   |     | 255        |       |
| classes12                | int(11)      | NO   |     | 255        |       |
| classes13                | int(11)      | NO   |     | 255        |       |
| classes14                | int(11)      | NO   |     | 255        |       |
| classes15                | int(11)      | NO   |     | 255        |       |
| CastingAnim              | int(11)      | NO   |     | 44         |       |
| TargetAnim               | int(11)      | NO   |     | 13         |       |
| TravelType               | int(11)      | NO   |     | 0          |       |
| SpellAffectIndex         | int(11)      | NO   |     | -1         |       |
| disallow_sit             | int(11)      | NO   |     | 0          |       |
| deities0                 | int(11)      | NO   |     | 0          |       |
| deities1                 | int(11)      | NO   |     | 0          |       |
| deities2                 | int(11)      | NO   |     | 0          |       |
| deities3                 | int(11)      | NO   |     | 0          |       |
| deities4                 | int(11)      | NO   |     | 0          |       |
| deities5                 | int(11)      | NO   |     | 0          |       |
| deities6                 | int(11)      | NO   |     | 0          |       |
| deities7                 | int(11)      | NO   |     | 0          |       |
| deities8                 | int(11)      | NO   |     | 0          |       |
| deities9                 | int(11)      | NO   |     | 0          |       |
| deities10                | int(11)      | NO   |     | 0          |       |
| deities11                | int(11)      | NO   |     | 0          |       |
| deities12                | int(12)      | NO   |     | 0          |       |
| deities13                | int(11)      | NO   |     | 0          |       |
| deities14                | int(11)      | NO   |     | 0          |       |
| deities15                | int(11)      | NO   |     | 0          |       |
| deities16                | int(11)      | NO   |     | 0          |       |
| npc_no_cast              | int(11)      | NO   |     | 0          |       |
| ai_pt_bonus              | int(11)      | NO   |     | 0          |       |
| new_icon                 | int(11)      | NO   |     | 161        |       |
| spellanim                | int(11)      | NO   |     | 0          |       |
| uninterruptable          | int(11)      | NO   |     | 0          |       |
| ResistDiff               | int(11)      | NO   |     | -150       |       |
| dot_stacking_exempt      | int(11)      | NO   |     | 0          |       |
| deleteable               | int(11)      | NO   |     | 0          |       |
| RecourseLink             | int(11)      | NO   |     | 0          |       |
| no_partial_resist        | int(11)      | NO   |     | 0          |       |
| small_targets_only       | int(11)      | NO   |     | 0          |       |
| use_persistent_particles | int(11)      | NO   |     | 0          |       |
| short_buff_box           | int(11)      | NO   |     | -1         |       |
| descnum                  | int(11)      | NO   |     | 0          |       |
| typedescnum              | int(11)      | YES  |     | NULL       |       |
| effectdescnum            | int(11)      | YES  |     | NULL       |       |
| effectdescnum2           | int(11)      | NO   |     | 0          |       |
| npc_no_los               | int(11)      | NO   |     | 0          |       |
| reflectable              | int(11)      | NO   |     | 0          |       |
| resist_per_level         | int(11)      | NO   |     | 0          |       |
| resist_cap               | int(11)      | NO   |     | 0          |       |
| EndurCost                | int(11)      | NO   |     | 0          |       |
| EndurTimerIndex          | int(11)      | NO   |     | 0          |       |
| IsDiscipline             | int(11)      | NO   |     | 0          |       |
| HateAdded                | int(11)      | NO   |     | 0          |       |
| EndurUpkeep              | int(11)      | NO   |     | 0          |       |
| pvpresistbase            | int(11)      | NO   |     | -150       |       |
| pvpresistcalc            | int(11)      | NO   |     | 100        |       |
| pvpresistcap             | int(11)      | NO   |     | -150       |       |
| spell_category           | int(11)      | NO   |     | -99        |       |
| pvp_duration             | int(11)      | NO   |     | 0          |       |
| pvp_duration_cap         | int(11)      | NO   |     | 0          |       |
| cast_not_standing        | int(11)      | NO   |     | 0          |       |
| can_mgb                  | int(11)      | NO   |     | 0          |       |
| nodispell                | int(11)      | NO   |     | -1         |       |
| npc_category             | int(11)      | NO   |     | 0          |       |
| npc_usefulness           | int(11)      | NO   |     | 0          |       |
| wear_off_message         | int(11)      | NO   |     | 0          |       |
| suspendable              | int(11)      | YES  |     | 0          |       |
| spellgroup               | int(11)      | YES  |     | 0          |       |
| allow_spellscribe        | int(11)      | NO   |     | 0          |       |
| allowrest                | int(11)      | YES  |     | 0          |       |
| custom_icon              | int(11)      | YES  |     | 0          |       |
| not_player_spell         | int(11)      | YES  |     | 0          |       |
| disabled                 | tinyint(3)   | NO   |     | 0          |       |
+--------------------------+--------------+------+-----+------------+-------+
*/

/* Classes
Number   Name
1	     Warrior
2        Cleric
3        Paladin
4        Ranger
5        Shadow Knight
6        Druid
7        Monk
8        Bard
9        Rogue
10       Shaman
11       Necromancer
12       Wizard
13       Magician
14       Enchanter
15       Beastlord
16       Berserker
*/
