package lemin

// existRoom checks if a room exists in a given path
func ExistRoom(rm *Room, newPath []*Room) bool {
	for _, v := range newPath {
		if v == rm {
			return false // Room already exists in the path
		}
	}
	return true // Room doesn't exist in the path
}
