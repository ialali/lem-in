package lemin

// BFS performs a breadth-first search on the graph to find all paths between start and end rooms
// a pointer to the Graph struct to be searched
func BFS(graph *Graph, start, end string) []*Path {
	// Set the start and end rooms of the graph by calling the GetRoom function
	graph.Start = GetRoom(graph, start)
	graph.End = GetRoom(graph, end)

	// Initialize a queue with the starting room
	queue := [][]*Room{{graph.Start}}

	// Create a slice to store all found paths in the graph
	var allPaths []*Path

	// Continue BFS until the queue is empty
	for len(queue) != 0 {
		// Dequeue the first path from the queue
		curPath := queue[0]
		lastRoom := curPath[len(curPath)-1]
		queue = queue[1:]

		// Explore adjacent rooms
		// adj = adjRm adjacent room
		// Children = Connection
		for _, adj := range lastRoom.Children {
			// If the adjacent room is the end room, create a new path and add it to allPaths
			if adj == graph.End {
				path := &Path{P: curPath}
				path.P = append(path.P, adj)
				allPaths = append(allPaths, path)
				continue
			}

			// Create a new path by copying the current path
			newPath := make([]*Room, len(curPath))
			copy(newPath, curPath)

			// Check if the adjacent room exists in the new path to avoid revisiting the same room
			if ExistRoom(adj, newPath) {
				// Add the adjacent room to the new path and enqueue the new path
				newPath = append(newPath, adj)
				queue = append(queue, newPath)
			} else {
				// If the adjacent room already exists in the new path, skip room to avoid looping through same rooms
				continue 
			}
		}
	}
	return allPaths //this contains all the found paths between the start and end rooms.
}
