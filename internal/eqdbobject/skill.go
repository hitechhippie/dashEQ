package eqdbobject

type Skill struct {
	SkillID uint8
	Class   uint8
	Level   uint8
	Cap     uint8
}

/*
+---------+-----------------------+------+-----+---------+-------+
| Field   | Type                  | Null | Key | Default | Extra |
+---------+-----------------------+------+-----+---------+-------+
| skillID | tinyint(3) unsigned   | NO   | PRI | 0       |       |
| class   | tinyint(3) unsigned   | NO   | PRI | 0       |       |
| level   | tinyint(3) unsigned   | NO   | PRI | 0       |       |
| cap     | mediumint(8) unsigned | NO   |     | 0       |       |
| class_  | tinyint(3) unsigned   | NO   | PRI | 0       |       |
+---------+-----------------------+------+-----+---------+-------+
*/
