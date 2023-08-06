package dashweb

import (
	"bytes"
	"dasheq/internal/config"
	"dasheq/internal/eqdbobject"
	"dasheq/internal/eqquest"
	"dasheq/internal/logging"
	"fmt"
	"log"
	"net/http"
	"runtime"
	"strconv"
	"strings"
	"text/template"
)

type Server struct {
	Config    *config.ServerConfig
	DataSet   *[]eqdbobject.DataSet
	Zone      *[]eqdbobject.Zone
	NPC       *[]eqdbobject.NPC
	Item      *[]eqdbobject.Item
	QuestNPC  *[]eqquest.QuestNPC
	QuestHear *[]eqquest.QuestHear
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

	printStatus(srv.Config)
	return nil
}

func contextHandler(w http.ResponseWriter, req *http.Request) {

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
		tmplIntermediate.Execute(&buf, Home)
		tmpl := template.Must(template.New("").Parse(buf.String()))
		tmpl.Execute(w, srv.DataSet)
	case "/zones.html":
		tmplIntermediate.Execute(&buf, Zones)
		tmpl := template.Must(template.New("").Parse(buf.String()))
		tmpl.Execute(w, srv.Zone)
	case "/items.html":
		tmplIntermediate.Execute(&buf, Items)
		tmpl := template.Must(template.New("").Parse(buf.String()))
		tmpl.Execute(w, srv.Item)
	case "/questnpcs.html":
		tmplIntermediate.Execute(&buf, QuestNPCs)
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
			tmplIntermediate.Execute(&buf, QuestNPCdetail)
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
			tmplIntermediate.Execute(&buf, ItemDetail)
			tmpl := template.Must(template.New("").Parse(buf.String()))
			tmpl.Execute(w, qItem)
		} else {
			http.ServeFile(w, req, "./static/html"+req.URL.Path)
		}
	}
}

func printStatus(c *config.ServerConfig) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	fmt.Println()
	fmt.Println("// system memory allocated:", (m.Sys / 1000000), "MB")
	fmt.Println("// system memory consumed:", (m.HeapAlloc / 1000000), "MB")
	fmt.Println()
	fmt.Println("*** DashEQ Listening for Requests on: " + c.WebAddr + ":" + c.WebPort + " ***")
	fmt.Println()
}
