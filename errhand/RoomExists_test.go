package lemin

import (
	lemin "lemin/functions"
	"testing"
)

func TestRoomExists(t *testing.T) {
	// Create a sample slice of rooms
	rooms := []*lemin.Room{
		&lemin.Room{Key: "Room1"},
		&lemin.Room{Key: "Room2"},
		&lemin.Room{Key: "Room3"},
	}

	// Test case 1: Check for an existing room
	exists := lemin.RoomExists(rooms, "Room2")
	if !exists {
		t.Error("Expected to find an existing room, but got false.")
	}

	// Test case 2: Check for a non-existing room
	exists = lemin.RoomExists(rooms, "Room4")
	if exists {
		t.Error("Expected not to find a non-existing room, but got true.")
	}
}
