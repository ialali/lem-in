package lemin

import "fmt"

// Adds an edge to the graph
// from = fromRm
// to = toRm
// Children = Connection

func AddEdge(graph *Graph, from, to string) error {
	// get Room
	fromRoom := GetRoom(graph, from)
	toRoom := GetRoom(graph, to)

	// error handling: current room, next room don't exist; from-to Room are same 
	if fromRoom == nil || toRoom == nil || fromRoom == toRoom {
		err := fmt.Errorf("ERR: invalid edge(%v --> %v)", from, to)
		return err
	} else if RoomExists(fromRoom.Children, to) { //does edge already exist?
		err := fmt.Errorf("ERR: existing edge(%v --> %v)", from, to)
		return err
	} else {
		// add the edge (connection) between the rooms (nodes)
		fromRoom.Children = append(fromRoom.Children, toRoom)
		toRoom.Children = append(toRoom.Children, fromRoom)
	}
	return nil //edge (connection) added successfully
}
