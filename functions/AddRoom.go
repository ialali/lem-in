package lemin

import "fmt"

// AddRoom adds a Room by its key to the Graph data struct
func AddRoom(graph *Graph, k string) error {
	if RoomExists(graph.Rooms, k) { //check if Room Exists
		err := fmt.Errorf("ERR: room %v already exists", k)
		return err
	} else {
		graph.Rooms = append(graph.Rooms, &Room{Key: k}) //create a new Room stucture with key k and add to the graph
	}
	return nil //new Room successfully added
}
