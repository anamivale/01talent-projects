package structs

type AntColony struct {
	Ants  int
	Rooms []Room
	Links map[string][]string
	Start string
	End   string
}
type Room struct {
	Name      string
	IsVisited bool
	X         int
	Y         int
}

type Path struct {
	Rooms []string
}

var FileContents string

var Existinglink = make(map[string]bool)
