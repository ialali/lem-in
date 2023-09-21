package lemin

// tick = timeInterval
// RemoveStart = RemoveStartRm

import (
	"fmt"
	"strconv"
)

// PrintAntMovement prints the ant migration along the optimal path using ant allocations
func PrintAntMovement(p *Allpaths, ants int) {
	// Initialize a 2D slice to store ant movements for each tick (timeInterval)
	movementOutput := make([][]string, p.MinTravelTime-1)
	antName := 1 //unique integer identifier for ants

	// Allocate ants to paths using the AllocateAnts function
	antAllocations := AllocateAnts(p, ants)

	// Remove the start room from the optimal path
	RemoveStart(p)

	// Iterate over allocated ants for each path, [i] is index of path
	for i, antCount := range antAllocations {
		// Iterate over each ant's movements along the path, [j] is index of each ant
		for j := 0; j < antCount; j++ {
			// Iterate over rooms in the current path
			for k, room := range p.OptPath[i].P { //[k] is index of each room
				// Append ant movement to the corresponding tick(timeTaken) in the movementOutput
				// construct a string in format of 'L(int)<antName>-RoomName<room.Key>' 
				movementOutput[k+j] = append(movementOutput[k+j], "L"+strconv.Itoa(antName)+"-"+room.Key)
			}
			// Increment the ant identifier
			antName++ //next unique ant
		}
	}

	// Print the ant movements
	// level = moveOfAnts 
	// movement = moveIndAnt individual movement of 1 ant at that specific time interval
	fmt.Println("----------- Ants Movements ------------- ")
	for _, level := range movementOutput {
		for _, movement := range level {
			fmt.Print(movement + " ")
		}
		fmt.Println()
	}
	fmt.Println("Number of turns:", len(movementOutput))
}
//the nested loop iterates over all the ant movements within a specific time interval and prints them together on the same line
