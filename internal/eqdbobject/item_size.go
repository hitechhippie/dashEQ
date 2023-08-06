package eqdbobject

type ItemSize uint8

func (s ItemSize) String() string {
	switch s {
	case 0:
		return "TINY"
	case 1:
		return "SMALL"
	case 2:
		return "MEDIUM"
	case 3:
		return "LARGE"
	case 4:
		return "GIANT"
	case 5:
		return "GIGANTIC"
	default:
		return "UNKNOWN"
	}
}
