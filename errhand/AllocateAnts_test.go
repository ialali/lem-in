package lemin

import (
	"testing"
)

func TestAllocateAnts(t *testing.T) {
    // Create a sample Allpaths struct with paths and ants
    pathsInfo := &Allpaths{
        Paths: []*Path{
            &Path{P: []*Room{}},
            &Path{P: []*Room{}},
            &Path{P: []*Room{}},
        },
    }
    totalAnts := 10

    // Test the AllocateAnts function
    antAllocations := AllocateAnts(pathsInfo, totalAnts)

    // Check if the allocations are correct
    expectedAllocations := []int{3, 3, 4} // Expected allocations for the given paths and ants
    for i, allocation := range antAllocations {
        if allocation != expectedAllocations[i] {
            t.Errorf("Expected allocation of %d ants for path %d, but got %d", expectedAllocations[i], i, allocation)
        }
    }

    // Test case with no ants to allocate
    emptyPathsInfo := &Allpaths{}
    emptyAnts := 0
    emptyAllocations := AllocateAnts(emptyPathsInfo, emptyAnts)
    if len(emptyAllocations) != 0 {
        t.Errorf("Expected empty allocations, but got %v", emptyAllocations)
    }
}
