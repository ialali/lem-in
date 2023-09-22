package lemin

import (
	"testing"
	lemin "lemin/functions"
)


	// Test case 1: Check for an existing roomfunc TestRoomExists(t *testing.T) {
	// Create a sample slice of rooms
	rooms := []*Room{
		&Room{Key: "Room1"},
		&Room{Key: "Room2"},
		&Room{Key: "Room3"},
	}

	// Test case 1: Check for an existing room
	exists := RoomExists(rooms, "Room2")
	if !exists {
		t.Error("Expected to find an existing room, but got false.")
	}

	// Test case 2: Check for a non-existing room
	exists = RoomExists(rooms, "Room4")
	if exists {
		t.Error("Expected not to find a non-existing room, but got true.")
	}
}
