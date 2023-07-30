package main

import (
	"dasheq/internal/config"
	eqdb "dasheq/internal/db"
	eqobject "dasheq/internal/eqobjects"
	"dasheq/internal/eqquests"
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime"
	"text/template"
	"time"
)

var eqDBConnection *eqdb.Connection
var dataSets []eqobject.DataSet
var zoneSet []eqobject.Zone
var npcSet []eqobject.NPC
var questNPCset *[]eqquests.QuestNPC
var questHearSet *[]eqquests.QuestHear

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
	zoneSet, npcSet, dataSets, err = loadDataSetsAll(eqDBConnection)
	if err != nil {
		fmt.Println("! [DB] Data set load error:", err)
		os.Exit(1)
	}
	fmt.Println("* [DB] Loaded", len(zoneSet), "zones.")
	fmt.Println("* [DB] Loaded", len(npcSet), "NPCs.")

	questNPCset, err = eqquests.LoadDataQuestNPCs(dashEQConfig.QuestDir, &zoneSet, &npcSet)
	if err != nil {
		fmt.Println("! Quest NPC load error:", err)
		os.Exit(1)
	}
	fmt.Println("* [Quest] Loaded", len(*questNPCset), "quest NPCs.")

	questHearSet, err = eqquests.LoadDataQuestHears(questNPCset)
	if err != nil {
		fmt.Println("! Quest hear set load error:", err)
		os.Exit(1)
	}
	fmt.Println("* [Quest] Loaded", len(*questHearSet), "quest hear/response statements.")

	for _, data := range *questHearSet {
		if data.QuestNPCId == 55179 {
			fmt.Println(".. player says " + data.Hears)
			fmt.Println(".. NPC says " + data.Says)
		}
	}

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

func loadDataSetsAll(c *eqdb.Connection) ([]eqobject.Zone, []eqobject.NPC, []eqobject.DataSet, error) {

	dataSets := make([]eqobject.DataSet, 0)

	zone, err := loadDataSetsZones(c)
	if err != nil {
		return nil, nil, nil, err
	}
	dataSets = append(dataSets, eqobject.DataSet{Name: "Zones", Count: uint32(len(zone)), LoadTime: time.Now().String()})

	npcs, err := loadDataSetsNPCs(c)
	if err != nil {
		return nil, nil, nil, err
	}
	dataSets = append(dataSets, eqobject.DataSet{Name: "NPCs", Count: uint32(len(npcs)), LoadTime: time.Now().String()})

	return zone, npcs, dataSets, nil
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

	log.Println("request received from", req.RemoteAddr+":", req.RequestURI)

	switch req.RequestURI {
	case "/", "/index.html":
		tmpl, err := template.ParseFiles("./static/html/templates/index.html")
		if err != nil {
			fmt.Println("! Web processing error:", err)
		}
		tmpl.Execute(w, dataSets)
	case "/zones.html":
		tmpl, err := template.ParseFiles("./static/html/templates/zones.html")
		if err != nil {
			fmt.Println("! Web processing error:", err)
		}
		tmpl.Execute(w, zoneSet)
	case "/npcs.html":
		tmpl, err := template.ParseFiles("./static/html/templates/npcs.html")
		if err != nil {
			fmt.Println("! Web processing error:", err)
		}
		tmpl.Execute(w, npcSet)
	case "/questnpcs.html":
		tmpl, err := template.ParseFiles("./static/html/templates/questnpcs.html")
		if err != nil {
			fmt.Println("! Web processing error:", err)
		}
		tmpl.Execute(w, questNPCset)
	case "/?reload-data-sets":
		zoneSet, npcSet, dataSets, _ = loadDataSetsAll(eqDBConnection)

		fmt.Println("*[DB] Loaded", len(zoneSet), "zones.")
		fmt.Println("*[DB] Loaded", len(npcSet), "NPCs.")
		http.Redirect(w, req, "/", http.StatusMovedPermanently)
	default:
		http.ServeFile(w, req, "./static/html"+req.URL.Path)
	}
}
