package lemin

import (
	"testing"
)

func TestGetData(t *testing.T) {
    // Create a sample dataFile
    dataFile := []string{
        "4",
        "4 5 4",
        "##start",
        "0 1 4",
        "1 3 6",
        "##end",
        "5 6 4",
        "2 3 4",
        "3 3 1",
        "0-1",
        "2-4",
        "1-4",
        "0-2",
        "4-5",
        "3-0",
        "4-3",
    }

    // Create an empty Input struct
    data := &Input{}

    // Test the GetData function
    err := GetData(data, dataFile)
    if err != nil {
        t.Errorf("Expected no error, but got an error: %v", err)
    }

    // Verify that data is populated correctly
    if data.Ants != 4 {
        t.Errorf("Expected Ants to be 4, but got %d", data.Ants)
    }
    if data.StartR != "0" {
        t.Errorf("Expected StartR to be '0', but got '%s'", data.StartR)
    }
    if data.EndR != "5" {
        t.Errorf("Expected EndR to be '5', but got '%s'", data.EndR)
    }
    if len(data.Rooms) != 7 {
        t.Errorf("Expected 7 rooms, but got %d", len(data.Rooms))
    }
    if len(data.Links) != 7 {
        t.Errorf("Expected 7 links, but got %d", len(data.Links))
    }
}
