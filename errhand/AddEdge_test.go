package lemin

import (
	lemin "lemin/functions"
	"testing"
)

func TestAddEdge(t *testing.T) {
    // Create a sample graph
    graph := &lemin.Graph{
        Rooms: []*lemin.Room{
            {Key: "Room1"},
            {Key: "Room2"},
        },
    }

    // Test case 1: Add a valid edge
    err := lemin.AddEdge(graph, "Room1", "Room2")
    if err != nil {
        t.Errorf("Expected no error, but got an error: %v", err)
    }

    // Verify that the edge was added correctly
    if !lemin.RoomExists(graph.Rooms[0].Children, "Room2") {
        t.Error("Expected Room2 to be in Room1's Children, but it's not.")
    }

    // Test case 2: Attempt to add an existing edge
    err = lemin.AddEdge(graph, "Room1", "Room2")
    if err == nil {
        t.Error("Expected an error for existing edge, but got no error.")
    }

    // Test case 3: Attempt to add an edge between non-existent rooms
    err = lemin.AddEdge(graph, "Room1", "Room3")
    if err == nil {
        t.Error("Expected an error for non-existent rooms, but got no error.")
    }
}

