package lemin

//ticks = timeTaken
//num = times

func FindOptimalPath(p *Allpaths, ticks []int, graph *Graph) {
	// Initialize the minimum travel time and optimal path with values from the first combination
	p.MinTravelTime = ticks[0]
	p.OptPath = p.Combo[0]

	// If the optimal path begins with the start room of the Graph and finishes with the end room
	// this is a special case where the optimum path only has the start & end rooms.
	if p.OptPath[0].P[0] == graph.Start && p.OptPath[0].P[1] == graph.End {
		// Modify the first combination to only contain the start and end rooms
		p.Combo[0] = p.Combo[0][0:1]
		// Update the optimal path with the modified combination
		p.OptPath = p.Combo[0]
	}

	// Iterate over the travel times and combinations to find the optimal path
	for i, num := range ticks {
		// If the current travel time is less than the previously known optimal travel time
		// find the combinations of paths with the shortest total travel time
		if num < p.MinTravelTime {
			// Update the optimal travel time
			p.MinTravelTime = num
			// Update the optimal path with the current combination's path
			p.OptPath = p.Combo[i]
		}
	}
}
