package eqquests

import (
	"bufio"
	"bytes"
	"os"
	"regexp"
	"strings"
)

type QuestHear struct {
	QuestNPCId uint32
	Hears      string
	GiveItems  []uint32
	Says       string
}

func LoadDataQuestHears(q *[]QuestNPC) (*[]QuestHear, error) {
	var questHearSet []QuestHear
	regExHearPre := regexp.MustCompile(`.*findi\(\"`)
	regExSayPre := regexp.MustCompile(`.*Say\(\"`)
	regExEmotePre := regexp.MustCompile(`.*Emote\(\"`)
	regExPost := regexp.MustCompile(`\"\).*`)

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
				questHear.Hears = regExHearPre.ReplaceAllString(currentLine, "")
				questHear.Hears = regExPost.ReplaceAllString(questHear.Hears, "")
			} else if strings.Contains(currentLine, "e.self:Say") {
				questHear.Says = regExSayPre.ReplaceAllString(currentLine, "")
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
