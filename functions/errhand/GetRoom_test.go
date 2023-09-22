package lemin

import (
	"testing"
)

func TestGetRoom(t *testing.T) {
    // Create a sample Graph with rooms
    graph := &Graph{
        Rooms: []*Room{
            &Room{Key: "Room1"},
            &Room{Key: "Room2"},
            &Room{Key: "Room3"},
        },
    }

    // Test case 1: Get an existing room
    room1 := GetRoom(graph, "Room1")
    if room1 == nil {
        t.Error("Expected to get an existing room, but got nil.")
    }
    if room1.Key != "Room1" {
        t.Errorf("Expected room with key 'Room1', but got '%s'", room1.Key)
    }

    // Test case 2: Get a non-existing room
    room4 := GetRoom(graph, "Room4")
    if room4 != nil {
        t.Error("Expected to get nil for a non-existing room, but got a room.")
    }
}
