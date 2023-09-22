package lemin

//structured records for the graph, rooms are nodes, connections are the edges linking the nodes
type Graph struct {
	Rooms      []*Room
	Start, End *Room
}

//individual room/node in the Graph
//Key is the unique name of the room/node
type Room struct {
	Key      string
	Children []*Room
}
// Children = Connections between parent Nodes

// P is slice of pointers to the Room struct
type Path struct {
	P []*Room
}

//
type Allpaths struct {
	Paths         []*Path
	Combo         [][]*Path
	OptPath       []*Path
	MinTravelTime int
}

// Input struct stores all input data for the Graph
type Input struct {
	Ants   int
	StartR string
	EndR   string
	Rooms  []string
	Coords [][]string
	Links  [][]string
}

//StartR = StartRm
//EndR = EndRm
//Rooms = GenRm
//linked to 'IsValidRoom' function
