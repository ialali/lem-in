package lemin

// CheckPathIntersection checks whether the given path intersects with rooms in the uniqueRooms map
// a pointer to the Path struct, to check for intersections
func CheckPathIntersection(uniqueRooms map[string]*Room, path *Path) bool {
	// Iterate through the rooms in the path, excluding the start and end rooms as intersections between these rooms are allowed.
	for i := 1; i < len(path.P)-1; i++ {
		// If the room is already present in uniqueRooms, there is an intersection
		// retrieve the room key from the map, and check that it exists
		if _, ok := uniqueRooms[path.P[i].Key]; ok {
			return false // There is already an intersection between the path and room, so return false
		}
	}
	return true // No intersections found, so return true
}
