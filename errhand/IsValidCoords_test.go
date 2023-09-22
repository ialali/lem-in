package lemin

import (
	"testing"
)

func TestIsValidCoords(t *testing.T) {
    // Create a sample Input struct
    data := &Input{}

    // Test case 1: Valid coordinates
    validCoords := []string{"2", "3"}
    isValid := IsValidCoords(validCoords, data)
    if !isValid {
        t.Error("Expected coordinates to be valid, but got invalid.")
    }

    // Test case 2: Negative coordinates
    negativeCoords := []string{"-1", "-1"}
    isValid = IsValidCoords(negativeCoords, data)
    if isValid {
        t.Error("Expected coordinates to be invalid (negative), but got valid.")
    }

    // Test case 3: Duplicate coordinates
    data.Coords = [][]string{{"2", "3"}}
    duplicateCoords := []string{"2", "3"}
    isValid = IsValidCoords(duplicateCoords, data)
    if isValid {
        t.Error("Expected coordinates to be invalid (duplicate), but got valid.")
    }
}
