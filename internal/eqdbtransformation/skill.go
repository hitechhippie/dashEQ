package eqdbtransformation

// create new spell based transform datasets

import "dasheq/internal/eqdbobject"

func SkillByClassSubset(sOrig *[]eqdbobject.Skill, sNew *[]eqdbobject.Skill, c uint8) error {

	for _, skillEntry := range *sOrig {
		if skillEntry.Class == eqdbobject.ClassID(c) {
			*sNew = append(*sNew, skillEntry)
		}
	}

	return nil
}
