package lemin

import (
	"testing"
)

func TestIsValidRoom(t *testing.T) {
    // Create a sample Input struct
    data := &Input{}

    // Test case 1: Valid room with coordinates
    err := IsValidRoom("Room1 2 3", "", data)
    if err != nil {
        t.Errorf("Expected no error, but got an error: %v", err)
    }

    // Test case 2: Room with invalid format
    err = IsValidRoom("InvalidRoomFormat", "", data)
    if err == nil {
        t.Error("Expected an error for invalid room format, but got no error.")
    }

    // Test case 3: Room with negative coordinates
    err = IsValidRoom("Room2 -1 -1", "", data)
    if err == nil {
        t.Error("Expected an error for negative coordinates, but got no error.")
    }

    // Test case 4: Room with duplicate coordinates
    data.Coords = [][]string{{"2", "3"}}
    err = IsValidRoom("Room3 2 3", "", data)
    if err == nil {
        t.Error("Expected an error for duplicate coordinates, but got no error.")
    }

    // Test case 5: Room with reserved starting character
    err = IsValidRoom("#InvalidRoom 1 2", "", data)
    if err == nil {
        t.Error("Expected an error for reserved starting character, but got no error.")
    }
}
