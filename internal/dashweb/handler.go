package dashweb

import (
	"bytes"
	"dasheq/internal/config"
	"dasheq/internal/eqdbobject"
	"dasheq/internal/eqdbtransformation"
	"dasheq/internal/eqquest"
	"dasheq/internal/logging"
	"fmt"
	"log"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"text/template"
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

// global server variable
var srv Server

func Run(s Server) error {
	// initialize error object for web request error handling
	var err error

	// copy the provided server instance object to our global instance
	srv = s

	// set up the primary web handling function for web requests
	http.HandleFunc("/", contextHandler)

	// open the web service
	err = http.ListenAndServe(srv.Config.WebAddr+":"+srv.Config.WebPort, nil)
	if err != nil {
		return err
	}

	return nil
}

func contextHandler(w http.ResponseWriter, req *http.Request) {

	// setup logging vars to log each web request
	var logWeb *log.Logger
	var err error

	// log each web request
	logWeb = logging.InitLogger("web", false)

	tmplIntermediate, err := template.ParseFiles("./static/html/templates/index.html")
	if err != nil {
		fmt.Println("! Web processing error:", err)
	}
	var buf bytes.Buffer

	logWeb.Println("request received from", req.RemoteAddr+":", req.RequestURI)

	switch req.RequestURI {
	case "/", "/index.html":
		tmplIntermediate.Execute(&buf, TmplHome)
		tmpl := template.Must(template.New("").Parse(buf.String()))
		tmpl.Execute(w, nil)
	case "/datasets.html":
		tmplIntermediate.Execute(&buf, TmplDataSets)
		tmpl := template.Must(template.New("").Parse(buf.String()))
		tmpl.Execute(w, srv.DataSet)
	case "/zones.html":
		tmplIntermediate.Execute(&buf, TmplZones)
		tmpl := template.Must(template.New("").Parse(buf.String()))
		tmpl.Execute(w, srv.Zone)
	case "/spells.html":
		tmplIntermediate.Execute(&buf, TmplSpells)
		tmpl := template.Must(template.New("").Parse(buf.String()))
		tmpl.Execute(w, srv.Spell)
	case "/skills.html":
		tmplIntermediate.Execute(&buf, TmplSkills)
		tmpl := template.Must(template.New("").Parse(buf.String()))
		tmpl.Execute(w, srv.Skill)
	case "/items.html":
		tmplIntermediate.Execute(&buf, TmplItems)
		tmpl := template.Must(template.New("").Parse(buf.String()))
		tmpl.Execute(w, srv.Item)
	case "/questnpcs.html":
		tmplIntermediate.Execute(&buf, TmplQuestNPCs)
		tmpl := template.Must(template.New("").Parse(buf.String()))
		tmpl.Execute(w, srv.QuestNPC)
	default:
		if strings.Contains(req.RequestURI, "questnpcid=") {
			query := req.URL.Query().Get("questnpcid")
			idUint64, _ := strconv.ParseUint(query, 10, 64)
			idUint32 := uint32(idUint64)

			var qNpcHearSubset []eqquest.QuestHear
			for _, data := range *srv.QuestHear {
				if data.QuestNPCId == idUint32 {
					qNpcHearSubset = append(qNpcHearSubset, data)
				}
			}
			tmplIntermediate.Execute(&buf, TmplQuestNPCdetail)
			tmpl := template.Must(template.New("").Parse(buf.String()))
			tmpl.Execute(w, qNpcHearSubset)

		} else if strings.Contains(req.RequestURI, "itemid=") {
			query := req.URL.Query().Get("itemid")
			idUint64, _ := strconv.ParseUint(query, 10, 64)
			idUint32 := uint32(idUint64)

			var qItem eqdbobject.Item
			for _, data := range *srv.Item {
				if data.Id == idUint32 {
					qItem = data
				}
			}
			tmplIntermediate.Execute(&buf, TmplItemDetail)
			tmpl := template.Must(template.New("").Parse(buf.String()))
			tmpl.Execute(w, qItem)

		} else if strings.Contains(req.RequestURI, "spellsbyclassid=") {
			query := req.URL.Query().Get("spellsbyclassid")
			idUint8, _ := strconv.ParseUint(query, 10, 8)

			var qSpellByClass []eqdbtransformation.SpellListByClass
			eqdbtransformation.PopulateSpellListByClass(&qSpellByClass, int(idUint8), srv.Spell, srv.Item)

			// Sort by level
			sort.Sort(eqdbtransformation.SpellListByClassSorted(qSpellByClass))

			tmplIntermediate.Execute(&buf, TmplSpellsByClass)
			tmpl := template.Must(template.New("").Parse(buf.String()))
			tmpl.Execute(w, qSpellByClass)

		} else if strings.Contains(req.RequestURI, "skillsbyclassid=") {
			var qSkillsByClassID []eqdbobject.Skill

			query := req.URL.Query().Get("skillsbyclassid")
			qSkillQueryID, _ := strconv.ParseUint(query, 10, 8)

			// Populate the SkillByClassSubset struct for query
			eqdbtransformation.SkillByClassSubset(srv.Skill, &qSkillsByClassID, uint8(qSkillQueryID))

			tmplIntermediate.Execute(&buf, TmplSkills)
			tmpl := template.Must(template.New("").Parse(buf.String()))
			tmpl.Execute(w, qSkillsByClassID)

		} else if strings.Contains(req.RequestURI, "itemsbyzonetop") {
			var zonesByExpansion []eqdbtransformation.ZoneListByExpansion

			// Populate the ZoneListByExpansion struct for query
			eqdbtransformation.PopulateZoneListByExpansion(&zonesByExpansion, srv.Zone)

			// Sort by expansion
			sort.Sort(eqdbtransformation.ZoneListByExpansionSorted(zonesByExpansion))

			tmplIntermediate.Execute(&buf, TmplItemsByZoneTop)
			tmpl := template.Must(template.New("").Parse(buf.String()))
			tmpl.Execute(w, zonesByExpansion)

		} else if strings.Contains(req.RequestURI, "itemsbyzone=") {
			var itemsByZone []eqdbobject.NPC

			query := req.URL.Query().Get("itemsbyzone")

			// Populate
			eqdbtransformation.PopulateNpcListByZone(&itemsByZone, srv.NPC, srv.Spawn2, srv.Spawnentry, query)

			tmplIntermediate.Execute(&buf, TmplItemsByZone)
			tmpl := template.Must(template.New("").Parse(buf.String()))
			tmpl.Execute(w, itemsByZone)

		} else if strings.Contains(req.RequestURI, "npcsbyzonetop") {
			var zonesByExpansion []eqdbtransformation.ZoneListByExpansion

			// Populate the ZoneListByExpansion struct for query
			eqdbtransformation.PopulateZoneListByExpansion(&zonesByExpansion, srv.Zone)

			// Sort by expansion
			sort.Sort(eqdbtransformation.ZoneListByExpansionSorted(zonesByExpansion))

			tmplIntermediate.Execute(&buf, TmplNpcsByZoneTop)
			tmpl := template.Must(template.New("").Parse(buf.String()))
			tmpl.Execute(w, zonesByExpansion)

		} else if strings.Contains(req.RequestURI, "npcsbyzone=") {
			var npcsByZone []eqdbobject.NPC
			var npcsByZoneDistinct []eqdbobject.NPC

			query := req.URL.Query().Get("npcsbyzone")

			// Populate
			eqdbtransformation.PopulateNpcListByZone(&npcsByZone, srv.NPC, srv.Spawn2, srv.Spawnentry, query)
			eqdbtransformation.DistinctNpcList(&npcsByZoneDistinct, &npcsByZone)

			sort.Sort(eqdbtransformation.NpcListByLevelSorted(npcsByZoneDistinct))

			tmplIntermediate.Execute(&buf, TmplNpcsByZone)
			tmpl := template.Must(template.New("").Parse(buf.String()))
			tmpl.Execute(w, npcsByZoneDistinct)

		} else {
			http.ServeFile(w, req, "./static/html"+req.URL.Path)
		}
	}
}
