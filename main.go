package main

import (
	"dasheq/internal/config"
	"dasheq/internal/dashweb"
	"dasheq/internal/eqdb"
	"dasheq/internal/eqdbobject"
	"dasheq/internal/eqquest"
	"fmt"
	"os"
	"time"
)

// global variables
var eqdbConnection *eqdb.Connection
var dashEQconfig *config.ServerConfig
var dataSets *[]eqdbobject.DataSet
var zoneSet *[]eqdbobject.Zone
var npcSet *[]eqdbobject.NPC
var questNPCset *[]eqquest.QuestNPC
var questHearSet *[]eqquest.QuestHear

func main() {
	// initialize the main runtime err object
	var err error

	// initialize our server object to pass to dashweb
	var srv dashweb.Server

	// configuration file location
	configFilePath := "./config.yaml"

	// load configuration from configFilePath
	dashEQconfig, err = config.LoadConfigFromYAML(configFilePath)
	if err != nil {
		fmt.Println("! Could not load configuration file:", err)
		fmt.Println("  ... please ensure you have a config.yaml file in the " +
			"server executable directory.")
		os.Exit(1)
	}
	fmt.Println("* Loaded configuration from", configFilePath)

	// initialize backend mysql database connection
	eqdbConnection, err = eqdb.Connect(dashEQconfig)
	if err != nil {
		fmt.Println("! Database connection error:", err)
		os.Exit(1)
	}

	// call the main data set load function and populate the data sets
	dataSets, err = loadDataSets(eqdbConnection, dashEQconfig)
	if err != nil {
		fmt.Println("! Data set load error:", err)
		os.Exit(1)
	}

	srv.Config = dashEQconfig
	srv.DataSet = dataSets
	srv.Zone = zoneSet
	srv.NPC = npcSet
	srv.QuestNPC = questNPCset
	srv.QuestHear = questHearSet

	err = dashweb.Run(srv)
	if err != nil {
		fmt.Println("! Couldn't launch web service:", err)
		os.Exit(1)
	}
}

func loadDataSets(c *eqdb.Connection, e *config.ServerConfig) (*[]eqdbobject.DataSet, error) {

	var err error
	dataSets := make([]eqdbobject.DataSet, 0)

	zoneSet, err = eqdbobject.LoadDataZone(c)
	if err != nil {
		return nil, err
	}
	dataSets = append(dataSets, eqdbobject.DataSet{Name: "Zones", Count: uint32(len(*zoneSet)), LoadTime: time.Now().String()})

	npcSet, err = eqdbobject.LoadDataNPC(c)
	if err != nil {
		return nil, err
	}
	dataSets = append(dataSets, eqdbobject.DataSet{Name: "NPCs", Count: uint32(len(*npcSet)), LoadTime: time.Now().String()})

	questNPCset, err = eqquest.LoadDataQuestNPC(e.QuestDir, zoneSet, npcSet)
	if err != nil {
		return nil, err
	}
	dataSets = append(dataSets, eqdbobject.DataSet{Name: "Quest NPCs", Count: uint32(len(*questNPCset)), LoadTime: time.Now().String()})

	questHearSet, err = eqquest.LoadDataQuestHear(questNPCset)
	if err != nil {
		return nil, err
	}
	dataSets = append(dataSets, eqdbobject.DataSet{Name: "Quest Hear/Responses", Count: uint32(len(*questHearSet)), LoadTime: time.Now().String()})

	fmt.Println("* [DB] Loaded", len(*zoneSet), "zones.")
	fmt.Println("* [DB] Loaded", len(*npcSet), "NPCs.")
	fmt.Println("* [Quest] Loaded", len(*questNPCset), "quest NPCs.")
	fmt.Println("* [Quest] Loaded", len(*questHearSet), "quest hear/response statements.")

	return &dataSets, nil
}
