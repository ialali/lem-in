package lemin

import (
	lemin "lemin/functions"
	"testing"
)

func TestAddRoom(t *testing.T) {
    // Create a sample graph
    graph := &lemin.Graph{
        Rooms: []*lemin.Room{
            {Key: "Room1"},
            {Key: "Room2"},
        },
    }

    // Test case 1: Add a valid room
    err := lemin.AddRoom(graph, "Room3")
    if err != nil {
        t.Errorf("Expected no error, but got an error: %v", err)
    }

    // Verify that the room was added correctly
    if !lemin.RoomExists(graph.Rooms, "Room3") {
        t.Error("Expected Room3 to be in the graph's Rooms, but it's not.")
    }

    // Test case 2: Attempt to add an existing room
    err = lemin.AddRoom(graph, "Room1")
    if err == nil {
        t.Error("Expected an error for an existing room, but got no error.")
    }
}