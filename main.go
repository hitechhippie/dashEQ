package main

import (
	"bytes"
	"dasheq/internal/config"
	"dasheq/internal/eqdb"
	"dasheq/internal/eqdbobject"
	"dasheq/internal/eqquest"
	"dasheq/internal/logging"
	"dasheq/internal/webtemplates"
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime"
	"strconv"
	"strings"
	"text/template"
	"time"
)

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
	zoneSet, npcSet, questNPCset, questHearSet, dataSets, err = loadDataSets(eqdbConnection, dashEQconfig)
	if err != nil {
		fmt.Println("! Data set load error:", err)
		os.Exit(1)
	}

	// set up the primary web handling function for web requests
	http.HandleFunc("/", webContextHandler)

	// print server status message in console
	printStatus(dashEQconfig)

	err = http.ListenAndServe(dashEQconfig.WebAddr+":"+dashEQconfig.WebPort, nil)
	if err != nil {
		fmt.Println("! Couldn't open web server socket:", err)
		os.Exit(1)
	}

	// close the databbase
	//err = eqdb.Close(eqdbConnection)
}

func loadDataSets(c *eqdb.Connection, e *config.ServerConfig) (
	*[]eqdbobject.Zone,
	*[]eqdbobject.NPC,
	*[]eqquest.QuestNPC,
	*[]eqquest.QuestHear,
	*[]eqdbobject.DataSet,
	error) {

	dataSets := make([]eqdbobject.DataSet, 0)

	zone, err := eqdbobject.LoadDataZone(c)
	if err != nil {
		return nil, nil, nil, nil, nil, err
	}
	dataSets = append(dataSets, eqdbobject.DataSet{Name: "Zones", Count: uint32(len(*zone)), LoadTime: time.Now().String()})

	npc, err := eqdbobject.LoadDataNPC(c)
	if err != nil {
		return nil, nil, nil, nil, nil, err
	}
	dataSets = append(dataSets, eqdbobject.DataSet{Name: "NPCs", Count: uint32(len(*npc)), LoadTime: time.Now().String()})

	qnpc, err := eqquest.LoadDataQuestNPC(e.QuestDir, zone, npc)
	if err != nil {
		return nil, nil, nil, nil, nil, err
	}
	dataSets = append(dataSets, eqdbobject.DataSet{Name: "Quest NPCs", Count: uint32(len(*qnpc)), LoadTime: time.Now().String()})

	qhear, err := eqquest.LoadDataQuestHear(qnpc)
	if err != nil {
		return nil, nil, nil, nil, nil, err
	}
	dataSets = append(dataSets, eqdbobject.DataSet{Name: "Quest Hear/Responses", Count: uint32(len(*qhear)), LoadTime: time.Now().String()})

	fmt.Println("* [DB] Loaded", len(*zone), "zones.")
	fmt.Println("* [DB] Loaded", len(*npc), "NPCs.")
	fmt.Println("* [Quest] Loaded", len(*qnpc), "quest NPCs.")
	fmt.Println("* [Quest] Loaded", len(*qhear), "quest hear/response statements.")

	return zone, npc, qnpc, qhear, &dataSets, nil
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
	var logWeb *log.Logger
	var err error

	// log each web request
	logWeb = logging.InitLogger("web")

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
		zoneSet, npcSet, questNPCset, questHearSet, dataSets, err = loadDataSets(eqdbConnection, dashEQconfig)
		if err != nil {
			logWeb.Println("error reloading datasets!!", err)
			http.Redirect(w, req, "/", http.StatusMovedPermanently)
		}
		http.Redirect(w, req, "/", http.StatusMovedPermanently)
	default:
		if strings.Contains(req.RequestURI, "questnpcid=") {
			query := req.URL.Query().Get("questnpcid")
			idUint64, _ := strconv.ParseUint(query, 10, 64)
			idUint32 := uint32(idUint64)

			var qNpcHearSubset []eqquest.QuestHear
			for _, data := range *questHearSet {
				if data.QuestNPCId == idUint32 {
					qNpcHearSubset = append(qNpcHearSubset, data)
				}
			}
			tmplIntermediate.Execute(&buf, webtemplates.QuestNPCdetail)
			tmpl := template.Must(template.New("").Parse(buf.String()))
			tmpl.Execute(w, qNpcHearSubset)
		} else {
			http.ServeFile(w, req, "./static/html"+req.URL.Path)
		}
	}
}
