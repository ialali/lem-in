package lemin

import (
	lemin "lemin/functions"
	"testing"
)

func TestIsValidRoom(t *testing.T) {
    // Create a sample Input struct
    data := &lemin.Input{}

    // Test case 1: Valid room with coordinates
    err := lemin.IsValidRoom("Room1 2 3", "", data)
    if err != nil {
        t.Errorf("Expected no error, but got an error: %v", err)
    }

    // Test case 2: Room with invalid format
    err = lemin.IsValidRoom("InvalidRoomFormat", "", data)
    if err == nil {
        t.Error("Expected an error for invalid room format, but got no error.")
    }

    // Test case 3: Room with negative coordinates
    err = lemin.IsValidRoom("Room2 -1 -1", "", data)
    if err == nil {
        t.Error("Expected an error for negative coordinates, but got no error.")
    }

    // Test case 4: Room with duplicate coordinates
    data.Coords = [][]string{{"2", "3"}}
    err = lemin.IsValidRoom("Room3 2 3", "", data)
    if err == nil {
        t.Error("Expected an error for duplicate coordinates, but got no error.")
    }

    // Test case 5: Room with reserved starting character
    err = lemin.IsValidRoom("#InvalidRoom 1 2", "", data)
    if err == nil {
        t.Error("Expected an error for reserved starting character, but got no error.")
    }
}
