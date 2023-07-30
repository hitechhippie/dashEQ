package eqquests

import (
	"bytes"
	"dasheq/internal/eqobjects"
	"errors"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

/*
QuestNPC
  Id (NPC->Id) uint32
  Name string
  Zone (Zone->Id) uint32
  QuestHear
    Id (QuestNPC->Id) uint32
	Hears string
	GiveItems {uint32...} (Item->Id)
	Says string
  QuestReceive
    Id (QuestNPC->Id) uint32
	Receives {uint32...} (Item->Id)
	GiveCash
	GiveItems {uint32...} (Item->Id)
*/

/*
QuestReward(mob, uint32 copper);

QuestReward(mob, uint32 copper, uint32 silver, uint32 gold, uint32 platinum, uint32 item_id, uint32 exp); ** this is most common found?

QuestReward(mob, uint32 copper, uint32 silver, uint32 gold);
QuestReward(mob, uint32 copper, uint32 silver);
QuestReward(mob, uint32 copper, uint32 silver, uint32 gold, uint32 platinum, uint32 item_id, uint32 exp, bool faction);
QuestReward(mob);
QuestReward(mob, uint32 copper, uint32 silver, uint32 gold, uint32 platinum);
QuestReward(mob, uint32 copper, uint32 silver, uint32 gold, uint32 platinum, uint32 item_id);
*/

type QuestNPC struct {
	Id   uint32
	Name string
	Zone uint32
	File string
}

func LoadDataQuestNPCs(d string, zoneSet *[]eqobjects.Zone, npcSet *[]eqobjects.NPC) (*[]QuestNPC, error) {
	var questNPCset []QuestNPC
	curDir := filepath.Base(d)

	// setup logging vars to log this load run
	var logBuf bytes.Buffer
	logQuest := log.New(&logBuf, "quest loader: ", log.Lshortfile)
	logFile, err := os.OpenFile("./logs/dashEQ-questloader.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return nil, err
	}
	defer logFile.Close()
	logQuest.SetOutput(logFile)

	err = filepath.WalkDir(d, func(path string, de fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if de.IsDir() {
			curDir = de.Name()
			return nil
		} else if !de.IsDir() && strings.HasSuffix(de.Name(), ".lua") {
			questNPCset, err = processQuestFile(path, curDir, questNPCset, zoneSet, npcSet, logQuest)
			if err != nil {
				return err
			} else {
				return nil
			}
		} else {
			return nil
		}
	})
	if err != nil {
		return nil, err
	}

	return &questNPCset, nil
}

// f = filename, d = current working directory
func processQuestFile(f string, d string, questNPCset []QuestNPC, zoneSet *[]eqobjects.Zone, npcSet *[]eqobjects.NPC, logQuest *log.Logger) ([]QuestNPC, error) {
	data, err := os.ReadFile(f)
	if err != nil {
		return nil, err
	}
	// build the case-insensitive regex object to strip LUA file extensions
	regExLUA := regexp.MustCompile(`(?i)\.lua`)

	if strings.Contains(string(data), "e.message:findi") ||
		strings.Contains(string(data), "item_lib.check_turn_in") {
		questNPC := new(QuestNPC)
		questNPC.File = f
		questNPC.Name = filepath.Base(f)
		questNPC.Name = regExLUA.ReplaceAllString(questNPC.Name, "")

		questNPC.Id, err = npcIdLookup(questNPC.Name, npcSet)
		if err != nil {
			logQuest.SetPrefix(time.Now().Format("2006-01-02 15:04:05") + " - ")
			logQuest.Println(err)
			return questNPCset, nil
		}

		questNPC.Zone, err = zoneIdLookup(d, zoneSet)
		if err != nil {
			logQuest.SetPrefix(time.Now().Format("2006-01-02 15:04:05") + " - ")
			logQuest.Println(err)
			return questNPCset, nil
		}

		questNPCset = append(questNPCset, *questNPC)
		return questNPCset, nil
	}
	return questNPCset, nil
}

// s = NPC name string determined from quest filename
// d = the populated NPC slice with all the NPC data
func npcIdLookup(s string, npcSet *[]eqobjects.NPC) (uint32, error) {
	for _, d := range *npcSet {
		if strings.Compare(s, d.Name) == 0 {
			return d.Id, nil
		}
	}
	err := errors.New("NPC ID not found in NPC data set: " + s)
	return 0, err
}

// s = Zone name string determined from current directory
// d = the populated NPC slice with all the NPC data
func zoneIdLookup(s string, zoneSet *[]eqobjects.Zone) (uint32, error) {
	for _, d := range *zoneSet {
		if strings.Compare(s, d.Short_name) == 0 {
			return d.Id, nil
		}
	}
	err := errors.New("Zone ID not found in Zone data set: " + s)
	return 0, err
}
