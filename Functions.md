## Functions used in lem-in task

#### main

The main function reads the input file, parses the data, creates a graph, finds paths using BFS, generates path combinations, calculates travel times, identifies the optimal path, and prints ant movements.

The code flows from parsing the input data to creating a graph structure, finding paths, generating combinations, calculating travel times, determining the optimal path, allocating ants, and finally printing the ant movements. Each function contributes to this overall flow by handling specific tasks and interactions between data structures.

#### GetData

Checks the inital requirements: int for ants; start & end rooms; a connection between start and end rooms; valid room name and coordinates.
This function extracts relevant data from the input file and populates the Input struct.

#### IsValidCoords

This function checks if the coordinates for each room are valid and unique from the input data given.

#### RoomExists

This function checks if a room with a specific name already exists in a slice of rooms.

#### IsValidRoom

This function validates and processes a room entry from the input data. It uses IsValidCoords and updates the Input struct with room information.

#### AddRoom

This function adds a room to the Graph structure. It ensures that the room being added doesn't already exist.

#### GetRoom

This function retrieves a pointer to a room with a given name from the Graph.

#### existRoom

This function checks if a given room exists in a specific path.

#### AddEdge

This function adds an edge (link) between two rooms in the Graph. It ensures the validity of the edge and prevents duplicate edges.

#### CheckPathIntersection

This function checks if a given path intersects with rooms that are already included in the combination.

#### AddPathToUniqueRooms

This function adds the rooms in a path to the uniqueRooms map, which helps track rooms in a combination.

#### BFS - Breadth First Search

This function performs a breadth-first search (BFS) on the graph to find all paths between the start and end rooms. It uses existRoom to determine if a room exists in a path.

#### GenerateCombinations

This function generates combinations of paths that can be taken by ants. It uses CheckPathIntersection and AddPathToUniqueRooms to create combinations.

#### CalculateTravelTimesForEachCombination

This function calculates the travel times for each combination of paths. It takes into account the number of ants and the length of paths.

#### RemoveStart

This function removes the start room from the paths in the optimal path combination, as this is room-zero where all ants start from, and never travel to.

#### FindOptimalPath

This function determines the optimal path for ants to take based on travel times and paths. It updates p.MinTravelTime and p.OptPath.

#### AllocateAnts

This function allocates ants to different paths based on the optimal path and travel times.

#### PrintAntMovement

This function prints the movement of all the ants along the optimal paths used. It uses the ant allocations and optimal path to generate the ant movements.
