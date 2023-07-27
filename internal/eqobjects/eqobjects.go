package eqobjects

type DataSet struct {
	Name     string
	Count    uint32
	LoadTime string
}

type NPC struct {
	Id    uint32
	Name  string
	Level uint8
}

type Zone struct {
	Id         uint16
	Short_name string
	Long_name  string
	Outdoor    uint8
	Expansion  uint8
}
