package eqquest

import (
	"bufio"
	"bytes"
	"dasheq/internal/eqdbobject"
	"dasheq/internal/logging"
	"errors"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func LoadDataQuestNPC(d string, zoneSet *[]eqdbobject.Zone, npcSet *[]eqdbobject.NPC) (*[]QuestNPC, error) {
	var questNPCset []QuestNPC
	var logQuest *log.Logger
	var err error
	var curDir string

	// initialize current directory to filepath base
	curDir = filepath.Base(d)

	// log each data load run
	logQuest = logging.InitLogger("loaddataquestnpc")

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

func LoadDataQuestHear(q *[]QuestNPC) (*[]QuestHear, error) {
	var questHearSet []QuestHear
	regExHearPre := regexp.MustCompile(`.*findi\(\"`)
	regExSayPre := regexp.MustCompile(`.*Say\(\"`)
	regExEmotePre := regexp.MustCompile(`.*Emote\(\"`)
	regExPost := regexp.MustCompile(`\"\).*`)
	regExPlayerName := regexp.MustCompile(`e\.other:GetCleanName\(\)`)

	for _, q := range *q {
		data, err := os.ReadFile(q.File)
		if err != nil {
			return nil, err
		}

		questHear := new(QuestHear)
		scanner := bufio.NewScanner(bytes.NewReader(data))
		for scanner.Scan() {
			currentLine := scanner.Text()

			if strings.Contains(currentLine, "e.message:findi") {
				if questHear.QuestNPCId != uint32(0) {
					questHearSet = append(questHearSet, *questHear)
					questHear = new(QuestHear)
				}
				questHear.QuestNPCId = q.Id
				questHear.QuestNPCName = q.Name
				questHear.Hears = regExHearPre.ReplaceAllString(currentLine, "")
				questHear.Hears = regExPost.ReplaceAllString(questHear.Hears, "")
			} else if strings.Contains(currentLine, "e.self:Say") {
				questHear.Says = regExSayPre.ReplaceAllString(currentLine, "")
				questHear.Says = regExPlayerName.ReplaceAllString(questHear.Says, "(player)")
				questHear.Says = regExPost.ReplaceAllString(questHear.Says, "")
			} else if strings.Contains(currentLine, "e.self.Emote") {
				questHear.Says = regExEmotePre.ReplaceAllString(currentLine, "")
				questHear.Says = regExPost.ReplaceAllString(questHear.Says, "")
			}
		}
		questHearSet = append(questHearSet, *questHear)
	}
	return &questHearSet, nil
}

// f = filename, d = current working directory
func processQuestFile(f string, d string, questNPCset []QuestNPC, zoneSet *[]eqdbobject.Zone, npcSet *[]eqdbobject.NPC, l *log.Logger) ([]QuestNPC, error) {
	data, err := os.ReadFile(f)
	if err != nil {
		return nil, err
	}
	// build the case-insensitive regex object to strip LUA file extensions
	regExLUA := regexp.MustCompile(`(?i)\.lua`)

	if strings.Contains(string(data), "e.message:findi") ||
		strings.Contains(string(data), "item_lib.check_turn_in") {
		qnpc := new(QuestNPC)
		qnpc.File = f
		qnpc.Name = filepath.Base(f)
		qnpc.Name = regExLUA.ReplaceAllString(qnpc.Name, "")

		qnpc.Id, err = npcIdLookup(qnpc.Name, npcSet)
		if err != nil {
			l.Println(err)
			return questNPCset, nil
		}

		qnpc.Zone, err = zoneIdLookup(d, zoneSet)
		if err != nil {
			l.Println(err)
			return questNPCset, nil
		}

		questNPCset = append(questNPCset, *qnpc)
		return questNPCset, nil
	}
	return questNPCset, nil
}

// s = NPC name string determined from quest filename
// d = the populated NPC slice with all the NPC data
func npcIdLookup(s string, n *[]eqdbobject.NPC) (uint32, error) {
	for _, d := range *n {
		if strings.Compare(s, d.Name) == 0 {
			return d.Id, nil
		}
	}
	err := errors.New("NPC ID not found in NPC data set: " + s)
	return 0, err
}

// s = Zone name string determined from current directory
// d = the populated NPC slice with all the NPC data
func zoneIdLookup(s string, z *[]eqdbobject.Zone) (uint32, error) {
	for _, d := range *z {
		if strings.Compare(s, d.Short_name) == 0 {
			return d.Id, nil
		}
	}
	err := errors.New("Zone ID not found in Zone data set: " + s)
	return 0, err
}
