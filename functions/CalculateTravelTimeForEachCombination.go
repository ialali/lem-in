package lemin

func CalculateTravelTimesForEachCombination(pathsInfo *Allpaths, totalAnts int) []int {
	travelTimes := []int{} // A slice to store calculated travel times for each combination

	for _, combination := range pathsInfo.Combo { // Iterate through each combination of paths
		combinedPathLength := 0 // To store the total length of all paths in the combination
		travelTime := 0         // To store the calculated travel time for this combination
		//the combinations are stored in the Combo field of the Allpath struct

		for _, paths := range combination { // Iterate through each individual path in the combination
			combinedPathLength += len(paths.P) // Add the length of the current path to the total length
			//calculates the total number of rooms visited across all paths in the combination
		}

		// Calculate the travel time using the formula: (totalAnts + (combinedPathLength - 1)) / len(combination)
		travelTime = (totalAnts + (combinedPathLength - 1)) / len(combination)

		// Append the calculated travel time for this combination to the travelTimes slice
		travelTimes = append(travelTimes, travelTime)
	}

	return travelTimes // Return the slice containing travel times for each combination
}
