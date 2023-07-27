package main

import (
	"dasheq/internal/config"
	eqdb "dasheq/internal/db"
	eqobject "dasheq/internal/eqobjects"
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime"
	"text/template"
)

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

	zoneSet, err := loadDataSetsZones(eqDBConnection)
	if err != nil {
		fmt.Println("![DB] Zone load error:", err)
		os.Exit(1)
	}
	fmt.Println("*[DB] Loaded", len(zoneSet), "zones.")

	npcSet, err := loadDataSetsNPCs(eqDBConnection)
	if err != nil {
		fmt.Println("![DB] NPC load error:", err)
		os.Exit(1)
	}
	fmt.Println("*[DB] Loaded", len(npcSet), "NPCs.")

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		log.Println("request received from", req.RemoteAddr+":", req.RequestURI)

		if req.RequestURI == "/" || req.RequestURI == "/index.html" {
			tmpl, err := template.ParseFiles("./static/html/templates/index.html")
			if err != nil {
				fmt.Println("! Error processing HTML template request:", err)
				os.Exit(1)
			}
			tmpl.Execute(w, zoneSet)
		} else if req.RequestURI == "/?reload-data-sets" {
			zoneSet, npcSet, err = loadDataSetsAll(eqDBConnection)
			fmt.Println("*[DB] Loaded", len(zoneSet), "zones.")
			fmt.Println("*[DB] Loaded", len(npcSet), "NPCs.")
			http.Redirect(w, req, "/", http.StatusMovedPermanently)
		} else {
			http.ServeFile(w, req, "./static/html"+req.URL.Path)
		}
	})

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

	// not sure the purpose of this defer
	// defer rows.Close()

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
		"level " + //Level uint8
		"FROM npc_types")
	if err != nil {
		return nil, err
	}

	// not sure the purpose of this defer
	// defer rows.Close()

	// iterate through the raw npc query data and populate the return object
	for npcRows.Next() {
		npc := new(eqobject.NPC)
		npcRows.Scan(&npc.Id, &npc.Name, &npc.Level)
		npcSet = append(npcSet, *npc)
	}
	return npcSet, nil
}

func loadDataSetsAll(c *eqdb.Connection) ([]eqobject.Zone, []eqobject.NPC, error) {
	zone, err := loadDataSetsZones(c)
	if err != nil {
		return nil, nil, err
	}

	npcs, err := loadDataSetsNPCs(c)
	if err != nil {
		return nil, nil, err
	}

	return zone, npcs, nil
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
