# dashEQ

## Workflows (WIP)
* Loot per zone SQL JOIN query; should do this in memory with structs:
```
select distinct i.id, i.Name
from npc_types n
inner join spawnentry se on n.id = se.npcID
inner join spawn2 s on se.spawngroupID = s.spawngroupID
inner join loottable_entries lte on n.loottable_id = lte.loottable_id
inner join lootdrop_entries lde on lte.lootdrop_id = lde.lootdrop_id
inner join items i on lde.item_id = i.id
where s.zone = 'vexthal'
order by i.Name;
```

* Loot Page:
  * Show, by race/class/level, available loot
  * Include quest rewards and droppable items

* NPC Page:
  * Show spawngroups
  * Show loot tables

* Quest Page:
  * Parse LUA quest scripts
  * Show quest NPC, Zone, Rewards
  * Optionally show quest script
  * Select Zone ->
    * Show all Zone Quest NPCs ->
      * Select Quest NPC ->
        * Show all Quest NPC:
          * Say 'X', NPC responds 'Y'
            * Show any item(s) handed to player
          * Give 'X item', NPC responds 'Y'
            * Show any item(s) handed to player
            * Show experience reward(s)
            * Show money reward(s)