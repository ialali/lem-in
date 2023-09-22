package lemin

import "fmt"

// Adds an edge to the graph
// Children = Relationship Connection to Parent Node

func AddEdge(graph *Graph, fromRm, toRm string) error {
	// get Room
	fromRoom := GetRoom(graph, fromRm)
	toRoom := GetRoom(graph, toRm)

	// error handling: current room, next room don't exist; from-to Room are same 
	if fromRoom == nil || toRoom == nil || fromRoom == toRoom {
		err := fmt.Errorf("ERR: invalid edge(%v --> %v)", fromRm, toRm)
		return err
	} else if RoomExists(fromRoom.Children, toRm) { //does edge already exist?
		err := fmt.Errorf("ERR: existing edge(%v --> %v)", fromRm, toRm)
		return err
	} else {
		// add the edge (connection) between the rooms (nodes)
		fromRoom.Children = append(fromRoom.Children, toRoom)
		toRoom.Children = append(toRoom.Children, fromRoom)
	}
	return nil //edge (connection) added successfully
}
