package eqdbobject

// future npc type implementations placeholder

/*
+------------------------+-----------------------+------+-----+---------+----------------+
| Field                  | Type                  | Null | Key | Default | Extra          |
+------------------------+-----------------------+------+-----+---------+----------------+
| id                     | int(11)               | NO   | PRI | NULL    | auto_increment |
| name                   | text                  | NO   |     | NULL    |                |
| lastname               | varchar(32)           | YES  |     | NULL    |                |
| level                  | tinyint(2) unsigned   | NO   |     | 0       |                |
| race                   | smallint(5) unsigned  | NO   |     | 0       |                |
| class                  | tinyint(2) unsigned   | NO   |     | 0       |                |
| bodytype               | int(11)               | NO   |     | 1       |                |
| hp                     | int(11)               | NO   |     | 0       |                |
| mana                   | int(11)               | NO   |     | 0       |                |
| gender                 | tinyint(2) unsigned   | NO   |     | 0       |                |
| texture                | tinyint(2) unsigned   | NO   |     | 0       |                |
| helmtexture            | tinyint(2) unsigned   | NO   |     | 0       |                |
| herosforgemodel        | int(11)               | NO   |     | 0       |                |
| size                   | float                 | NO   |     | 0       |                |
| hp_regen_rate          | int(11) unsigned      | NO   |     | 0       |                |
| mana_regen_rate        | int(11) unsigned      | NO   |     | 0       |                |
| loottable_id           | int(11) unsigned      | NO   |     | 0       |                |
| merchant_id            | int(11) unsigned      | NO   |     | 0       |                |
| alt_currency_id        | int(11) unsigned      | NO   |     | 0       |                |
| npc_spells_id          | int(11) unsigned      | NO   |     | 0       |                |
| npc_spells_effects_id  | int(11) unsigned      | NO   |     | 0       |                |
| npc_faction_id         | int(11)               | NO   |     | 0       |                |
| adventure_template_id  | int(10) unsigned      | NO   |     | 0       |                |
| trap_template          | int(10) unsigned      | YES  |     | 0       |                |
| mindmg                 | int(10) unsigned      | NO   |     | 0       |                |
| maxdmg                 | int(10) unsigned      | NO   |     | 0       |                |
| attack_count           | smallint(6)           | NO   |     | -1      |                |
| npcspecialattks        | varchar(36)           | NO   |     |         |                |
| special_abilities      | text                  | YES  |     | NULL    |                |
| aggroradius            | int(10) unsigned      | NO   |     | 0       |                |
| assistradius           | int(10) unsigned      | NO   |     | 0       |                |
| face                   | int(10) unsigned      | NO   |     | 1       |                |
| luclin_hairstyle       | int(10) unsigned      | NO   |     | 1       |                |
| luclin_haircolor       | int(10) unsigned      | NO   |     | 1       |                |
| luclin_eyecolor        | int(10) unsigned      | NO   |     | 1       |                |
| luclin_eyecolor2       | int(10) unsigned      | NO   |     | 1       |                |
| luclin_beardcolor      | int(10) unsigned      | NO   |     | 1       |                |
| luclin_beard           | int(10) unsigned      | NO   |     | 0       |                |
| drakkin_heritage       | int(10)               | NO   |     | 0       |                |
| drakkin_tattoo         | int(10)               | NO   |     | 0       |                |
| drakkin_details        | int(10)               | NO   |     | 0       |                |
| armortint_id           | int(10) unsigned      | NO   |     | 0       |                |
| armortint_red          | tinyint(3) unsigned   | NO   |     | 0       |                |
| armortint_green        | tinyint(3) unsigned   | NO   |     | 0       |                |
| armortint_blue         | tinyint(3) unsigned   | NO   |     | 0       |                |
| d_melee_texture1       | int(11)               | NO   |     | 0       |                |
| d_melee_texture2       | int(11)               | NO   |     | 0       |                |
| ammo_idfile            | varchar(30)           | NO   |     | IT10    |                |
| prim_melee_type        | tinyint(4) unsigned   | NO   |     | 28      |                |
| sec_melee_type         | tinyint(4) unsigned   | NO   |     | 28      |                |
| ranged_type            | tinyint(4) unsigned   | NO   |     | 7       |                |
| runspeed               | float                 | NO   |     | 0       |                |
| MR                     | smallint(5)           | NO   |     | 0       |                |
| CR                     | smallint(5)           | NO   |     | 0       |                |
| DR                     | smallint(5)           | NO   |     | 0       |                |
| FR                     | smallint(5)           | NO   |     | 0       |                |
| PR                     | smallint(5)           | NO   |     | 0       |                |
| Corrup                 | smallint(5)           | NO   |     | 0       |                |
| PhR                    | smallint(5) unsigned  | NO   |     | 0       |                |
| see_invis              | smallint(4)           | NO   |     | 0       |                |
| see_invis_undead       | smallint(4)           | NO   |     | 0       |                |
| qglobal                | int(2) unsigned       | NO   |     | 0       |                |
| AC                     | smallint(5)           | NO   |     | 0       |                |
| npc_aggro              | tinyint(4)            | NO   |     | 0       |                |
| spawn_limit            | tinyint(4)            | NO   |     | 0       |                |
| attack_speed           | float                 | NO   |     | 0       |                |
| attack_delay           | tinyint(3) unsigned   | NO   |     | 30      |                |
| findable               | tinyint(4)            | NO   |     | 0       |                |
| STR                    | mediumint(8) unsigned | NO   |     | 75      |                |
| STA                    | mediumint(8) unsigned | NO   |     | 75      |                |
| DEX                    | mediumint(8) unsigned | NO   |     | 75      |                |
| AGI                    | mediumint(8) unsigned | NO   |     | 75      |                |
| _INT                   | mediumint(8) unsigned | NO   |     | 80      |                |
| WIS                    | mediumint(8) unsigned | NO   |     | 75      |                |
| CHA                    | mediumint(8) unsigned | NO   |     | 75      |                |
| see_hide               | tinyint(4)            | NO   |     | 0       |                |
| see_improved_hide      | tinyint(4)            | NO   |     | 0       |                |
| trackable              | tinyint(4)            | NO   |     | 1       |                |
| isbot                  | tinyint(4)            | NO   |     | 0       |                |
| exclude                | tinyint(4)            | NO   |     | 1       |                |
| ATK                    | mediumint(9)          | NO   |     | 0       |                |
| Accuracy               | mediumint(9)          | NO   |     | 0       |                |
| Avoidance              | mediumint(9) unsigned | NO   |     | 0       |                |
| slow_mitigation        | smallint(4)           | NO   |     | 0       |                |
| version                | smallint(5) unsigned  | NO   |     | 0       |                |
| maxlevel               | tinyint(3)            | NO   |     | 0       |                |
| scalerate              | int(11)               | NO   |     | 100     |                |
| private_corpse         | tinyint(3) unsigned   | NO   |     | 0       |                |
| unique_spawn_by_name   | tinyint(3) unsigned   | NO   |     | 0       |                |
| underwater             | tinyint(3) unsigned   | NO   |     | 0       |                |
| isquest                | tinyint(3)            | NO   |     | 0       |                |
| emoteid                | int(10) unsigned      | NO   |     | 0       |                |
| spellscale             | float                 | NO   |     | 100     |                |
| healscale              | float                 | NO   |     | 100     |                |
| no_target_hotkey       | tinyint(1) unsigned   | NO   |     | 0       |                |
| raid_target            | tinyint(1) unsigned   | NO   |     | 0       |                |
| armtexture             | tinyint(2)            | NO   |     | 0       |                |
| bracertexture          | tinyint(2)            | NO   |     | 0       |                |
| handtexture            | tinyint(2)            | NO   |     | 0       |                |
| legtexture             | tinyint(2)            | NO   |     | 0       |                |
| feettexture            | tinyint(2)            | NO   |     | 0       |                |
| light                  | tinyint(2)            | NO   |     | 0       |                |
| walkspeed              | tinyint(2)            | NO   |     | 0       |                |
| peqid                  | int(11)               | NO   |     | 0       |                |
| unique_                | tinyint(2)            | NO   |     | 0       |                |
| fixed                  | tinyint(2)            | NO   |     | 0       |                |
| ignore_despawn         | tinyint(2)            | NO   |     | 0       |                |
| show_name              | tinyint(2)            | NO   |     | 1       |                |
| untargetable           | tinyint(2)            | NO   |     | 0       |                |
| charm_ac               | smallint(5)           | YES  |     | 0       |                |
| charm_min_dmg          | int(10)               | YES  |     | 0       |                |
| charm_max_dmg          | int(10)               | YES  |     | 0       |                |
| charm_attack_delay     | tinyint(3)            | YES  |     | 0       |                |
| charm_accuracy_rating  | mediumint(9)          | YES  |     | 0       |                |
| charm_avoidance_rating | mediumint(9)          | YES  |     | 0       |                |
| charm_atk              | mediumint(9)          | YES  |     | 0       |                |
| skip_global_loot       | tinyint(4)            | YES  |     | 0       |                |
| rare_spawn             | tinyint(4)            | YES  |     | 0       |                |
| stuck_behavior         | tinyint(4)            | NO   |     | 0       |                |
| model                  | smallint(5)           | NO   |     | 0       |                |
| flymode                | tinyint(4)            | NO   |     | -1      |                |
| always_aggro           | tinyint(1)            | NO   |     | 0       |                |
| exp_mod                | int(11)               | NO   |     | 100     |                |
+------------------------+-----------------------+------+-----+---------+----------------+
*/
