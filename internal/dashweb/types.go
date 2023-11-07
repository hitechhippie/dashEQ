package dashweb

import (
	"dasheq/internal/config"
	"dasheq/internal/eqdbobject"
	"dasheq/internal/eqquest"
)

type Server struct {
	Config     *config.ServerConfig
	DataSet    *[]eqdbobject.DataSet
	Zone       *[]eqdbobject.Zone
	NPC        *[]eqdbobject.NPC
	Spell      *[]eqdbobject.Spell
	Skill      *[]eqdbobject.Skill
	Item       *[]eqdbobject.Item
	Spawngroup *[]eqdbobject.Spawngroup
	Spawn2     *[]eqdbobject.Spawn2
	Spawnentry *[]eqdbobject.Spawnentry
	QuestNPC   *[]eqquest.QuestNPC
	QuestHear  *[]eqquest.QuestHear
}
