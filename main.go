package main

import (
	"bytes"
	"dasheq/internal/config"
	eqdb "dasheq/internal/db"
	eqobject "dasheq/internal/eqobjects"
	"dasheq/internal/eqquests"
	webtemplates "dasheq/internal/webtemplates"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"runtime"
	"strconv"
	"strings"
	"text/template"
	"time"
)

var eqDBConnection *eqdb.Connection
var dashEQConfig *config.ServerConfig
var dataSets []eqobject.DataSet
var zoneSet []eqobject.Zone
var npcSet []eqobject.NPC
var questNPCset []eqquests.QuestNPC
var questHearSet []eqquests.QuestHear

func main() {
	// configuration file location
	dashEQConfigFile := "./config.yaml"

	// load configuration from ./config.yaml
	dashEQConfig, err := loadConfiguration(dashEQConfigFile)
	if err != nil {
		fmt.Println("! Could not load configuration file:", err)
		fmt.Println("  ... please ensure you have a config.yaml file in the " +
			"server executable directory.")
		os.Exit(1)
	}
	fmt.Println("* Loaded configuration from", dashEQConfigFile)

	// initialize backend mysql database connection
	eqDBConnection, err := databaseConnect(dashEQConfig)
	if err != nil {
		fmt.Println("! Database connection error:", err)
		os.Exit(1)
	}

	// call the main data set load function and populate the data sets
	zoneSet, npcSet, questNPCset, questHearSet, dataSets, err = loadDataSetsAll(eqDBConnection, *dashEQConfig)
	if err != nil {
		fmt.Println("! Data set load error:", err)
		os.Exit(1)
	}
	fmt.Println("* [DB] Loaded", len(zoneSet), "zones.")
	fmt.Println("* [DB] Loaded", len(npcSet), "NPCs.")
	fmt.Println("* [Quest] Loaded", len(questNPCset), "quest NPCs.")
	fmt.Println("* [Quest] Loaded", len(questHearSet), "quest hear/response statements.")

	// set up the primary web handling function for web requests
	http.HandleFunc("/", webContextHandler)

	// print server status message in console
	printStatus(dashEQConfig)

	err = http.ListenAndServe(dashEQConfig.WebAddr+":"+dashEQConfig.WebPort, nil)
	if err != nil {
		fmt.Println("! Couldn't open web server socket:", err)
		os.Exit(1)
	}

	err = databaseDisconnect(eqDBConnection)
	if err != nil {
		fmt.Println("! Couldn't close the database:", err)
		os.Exit(1)
	}
}

func loadConfiguration(f string) (*config.ServerConfig, error) {
	dashEQconfig, err := config.LoadConfigFromYAML(f)
	if err != nil {
		return nil, err
	}

	return dashEQconfig, nil
}

func databaseConnect(c *config.ServerConfig) (*eqdb.Connection, error) {
	eqDBconfig := new(eqdb.ConnectionConfig)
	eqDBconnection := new(eqdb.Connection)

	eqDBconfig.Addr = c.DBaddr
	eqDBconfig.Dbname = c.DBname
	eqDBconfig.Net = c.DBnet
	eqDBconfig.Pass = c.DBpass
	eqDBconfig.User = c.DBuser

	eqDBconnection, err := eqdb.Connect(eqDBconfig)
	if err != nil {
		return nil, err
	} else {
		return eqDBconnection, err
	}
}

func databaseDisconnect(c *eqdb.Connection) error {
	err := c.Target.Close()
	if err != nil {
		return err
	} else {
		return nil
	}
}

func loadDataSetsZones(c *eqdb.Connection) ([]eqobject.Zone, error) {
	// initialize our return objects
	zoneSet := make([]eqobject.Zone, 0)

	// gather the raw query data for zones
	zoneRows, err := c.Target.Query("SELECT " +
		"id," + //Id uint16
		"short_name," + //Short_name string
		"long_name," + //Long_name  string
		"castoutdoor," + //Outdoor    uint8
		"expansion " + //Expansion  uint8
		"FROM zone")
	if err != nil {
		return nil, err
	}

	// iterate through the raw zone query data and populate the return object
	for zoneRows.Next() {
		zone := new(eqobject.Zone)
		zoneRows.Scan(&zone.Id, &zone.Short_name, &zone.Long_name, &zone.Outdoor, &zone.Expansion)
		zoneSet = append(zoneSet, *zone)
	}

	return zoneSet, nil
}

func loadDataSetsNPCs(c *eqdb.Connection) ([]eqobject.NPC, error) {
	// initialize our return objects
	npcSet := make([]eqobject.NPC, 0)

	// gather the raw query data for npcs
	npcRows, err := c.Target.Query("SELECT " +
		"id," + //Id uint32
		"name," + //Name string
		"level, " + //Level uint8
		"race, " + //uint16
		"class, " + //uint8
		"hp, " + //uint32
		"mana, " + //uint32
		"loottable_id, " + //uint32
		"npc_spells_id, " + //uint16
		"npc_faction_id, " + //uint16
		"mindmg, " + //uint16
		"maxdmg, " + //uint16
		"attack_count, " + //uint16
		"runspeed, " + //float32
		"MR, " + //uint16
		"CR, " + //uint16
		"DR, " + //uint16
		"FR, " + //uint16
		"PR, " + //uint16
		"AC " + //uint16
		"FROM npc_types")
	if err != nil {
		return nil, err
	}

	// iterate through the raw npc query data and populate the return object
	for npcRows.Next() {
		npc := new(eqobject.NPC)
		npcRows.Scan(&npc.Id, &npc.Name, &npc.Level, &npc.Race, &npc.Class, &npc.HP, &npc.Mana, &npc.LootTable, &npc.NPCSpells, &npc.NPCFaction, &npc.MinDmg, &npc.MaxDmg, &npc.AttackCount, &npc.RunSpeed, &npc.MR, &npc.CR, &npc.DR, &npc.FR, &npc.PR, &npc.AC)
		npcSet = append(npcSet, *npc)
	}
	return npcSet, nil
}

func loadDataSetsQuestNPCs(c *eqdb.Connection, p string, z []eqobject.Zone, n []eqobject.NPC) (*[]eqquests.QuestNPC, error) {
	//fmt.Fprintf(os.Stdout, "p = %v, %v, %v", p, &z, &n)
	qnpcset, err := eqquests.LoadDataQuestNPCs(p, &z, &n)
	if err != nil {
		return nil, err
	}

	return qnpcset, nil
}

func loadDataSetsQuestHears(c *eqdb.Connection, q []eqquests.QuestNPC) (*[]eqquests.QuestHear, error) {
	qhearset, err := eqquests.LoadDataQuestHears(&q)
	if err != nil {
		return nil, err
	}

	return qhearset, nil
}

func loadDataSetsAll(c *eqdb.Connection, e config.ServerConfig) ([]eqobject.Zone, []eqobject.NPC, []eqquests.QuestNPC, []eqquests.QuestHear, []eqobject.DataSet, error) {

	dataSets := make([]eqobject.DataSet, 0)

	zone, err := loadDataSetsZones(c)
	if err != nil {
		return nil, nil, nil, nil, nil, err
	}
	dataSets = append(dataSets, eqobject.DataSet{Name: "Zones", Count: uint32(len(zone)), LoadTime: time.Now().String()})

	npcs, err := loadDataSetsNPCs(c)
	if err != nil {
		return nil, nil, nil, nil, nil, err
	}
	dataSets = append(dataSets, eqobject.DataSet{Name: "NPCs", Count: uint32(len(npcs)), LoadTime: time.Now().String()})

	qnpc, err := loadDataSetsQuestNPCs(c, e.QuestDir, zone, npcs)
	if err != nil {
		return nil, nil, nil, nil, nil, err
	}
	dataSets = append(dataSets, eqobject.DataSet{Name: "Quest NPCs", Count: uint32(len(*qnpc)), LoadTime: time.Now().String()})

	qhear, err := loadDataSetsQuestHears(c, *qnpc)
	if err != nil {
		return nil, nil, nil, nil, nil, err
	}
	dataSets = append(dataSets, eqobject.DataSet{Name: "Quest Hear/Responses", Count: uint32(len(*qhear)), LoadTime: time.Now().String()})

	return zone, npcs, *qnpc, *qhear, dataSets, nil
}

func printStatus(c *config.ServerConfig) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	fmt.Println()
	fmt.Println("---*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*---")
	fmt.Println("// system memory allocated:", (m.Sys / 1000000), "MB")
	fmt.Println("// system memory consumed:", (m.HeapAlloc / 1000000), "MB")
	fmt.Println("---*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*---")
	fmt.Println()
	fmt.Println("" +
		"*** DashEQ Listening for Requests on: " + c.WebAddr + ":" + c.WebPort + " ***")
	fmt.Println()
}

func webContextHandler(w http.ResponseWriter, req *http.Request) {
	// setup logging vars to log each web request
	var logBuf bytes.Buffer
	logWeb := log.New(&logBuf, "web: ", log.Lshortfile)
	logFile, err := os.OpenFile("./logs/web.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("! Cannot open web log for writing:", err)
	}
	defer logFile.Close()
	logWeb.SetOutput(logFile)

	tmplIntermediate, err := template.ParseFiles("./static/html/templates/index.html")
	if err != nil {
		fmt.Println("! Web processing error:", err)
	}
	var buf bytes.Buffer

	logWeb.Println("request received from", req.RemoteAddr+":", req.RequestURI)

	switch req.RequestURI {
	case "/", "/index.html":
		tmplIntermediate.Execute(&buf, webtemplates.Home)
		tmpl := template.Must(template.New("").Parse(buf.String()))
		tmpl.Execute(w, dataSets)
	case "/zones.html":
		tmplIntermediate.Execute(&buf, webtemplates.Zones)
		tmpl := template.Must(template.New("").Parse(buf.String()))
		tmpl.Execute(w, zoneSet)
	case "/questnpcs.html":
		tmplIntermediate.Execute(&buf, webtemplates.QuestNPCs)
		tmpl := template.Must(template.New("").Parse(buf.String()))
		tmpl.Execute(w, questNPCset)
	case "/?reload-data-sets":
		zoneSet, npcSet, questNPCset, questHearSet, dataSets, err = loadDataSetsAll(eqDBConnection, *dashEQConfig)
		if err != nil {
			logWeb.Println("error reloading datasets!!", err)
			http.Redirect(w, req, "/", http.StatusMovedPermanently)
		} else {
			fmt.Println("* [DB] Loaded", len(zoneSet), "zones.")
			fmt.Println("* [DB] Loaded", len(npcSet), "NPCs.")
			http.Redirect(w, req, "/", http.StatusMovedPermanently)
		}
	default:
		query, _ := url.ParseQuery(req.URL.Opaque)
		if strings.Contains(req.RequestURI, "questnpcid=") {
			idUint64, _ := strconv.ParseUint(query.Get("questnpcid"), 10, 64)
			idUint32 := uint32(idUint64)
			fmt.Println(req.RequestURI)
			fmt.Println(query.Get("questnpcid"))
			fmt.Println(idUint64)
			fmt.Println(idUint32)

			var qNpcHearSubset []eqquests.QuestHear
			tmplIntermediate.Execute(&buf, webtemplates.QuestNPCdetail)
			for _, data := range questHearSet {
				if data.QuestNPCId == idUint32 {
					qNpcHearSubset = append(qNpcHearSubset, data)
				}
			}
			tmpl := template.Must(template.New("").Parse(buf.String()))
			tmpl.Execute(w, qNpcHearSubset)
		}
	}
}
