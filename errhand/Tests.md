# Tests used for Error Handling

To run the Go tests for the project, use the go test command followed by the package path of the tests you want to run.
Here's how to run the tests for the project:

- Open the terminal or command prompt.
- Navigate to the directory containing the Go project files.
- Run the following command to run all tests the your project:

```
go test ./...
```

This command tells Go to run all tests in the current directory and its subdirectories (./...)
If there are multiple packages within the project, this command will run tests for all of them.

Make sure you are in the root directory of the Go project where the Go files and test files are located.

After running the go test command, Go will execute the test functions and display the test results in the terminal.

## Explanations of tests used within the project

### Test for the AddEdge function:

Creates a sample graph with two rooms ("Room1" and "Room2") and then test various cases:

- Adding a valid edge between "Room1" and "Room2."
- Attempting to add the same edge again, which should result in an error.
- Attempting to add an edge between a room ("Room3") that doesn't exist in the graph, which should also result in an error.

---

### Test for the AddRoom function:

Create a sample graph with two existing rooms ("Room1" and "Room2").
Test two cases:

- Adding a valid room ("Room3") to the graph.
- Attempting to add an existing room ("Room1") to the graph again, which should result in an error.

---

### Test for the GetData function:

Create a sample dataFile containing data similar to what you provided.

- Create an empty Input struct.
- Test the GetData function with the sample data and check if it populates the data struct correctly.
- Verify that various fields in the data struct (e.g., Ants, StartR, EndR, Rooms, and Links) have the expected values.

---

### Test for the IsValidRoom function:

Create a sample Input struct.
Test several cases:

- Valid room with coordinates.
- Room with an invalid format.
- Room with negative coordinates.
- Room with coordinates that are already in use (duplicate).
- Room with a reserved starting character ('#').

Check if the function returns the expected errors or no errors based on the test cases.

---

### Test for the IsValidCoords function:

Create a sample Input struct.
Test several cases:

- Valid room with coordinates.
- Room with an invalid format.
- Room with negative coordinates.
- Room with coordinates that are already in use (duplicate).
- Room with a reserved starting character ('#').

Check if the function returns the expected errors or no errors based on the test cases.

---

### Test for the AllocateAnts function:

Create a sample Allpaths struct with three paths and a total of 10 ants.

- Call the AllocateAnts function with the sample data.
- Check if the allocations match the expected allocations for each path.
- Test the case with no ants to allocate to ensure it returns an empty allocation slice.

---

### Test for the GetRoom function:

Create a sample Graph struct with three rooms.
Test two cases:

- Getting an existing room ("Room1").
- Getting a non-existing room ("Room4").

The first case, check if the returned room is not nil and has the expected key ("Room1").
The second case, check if the function returns nil for a non-existing room.

---

### Test for the RoomExists function:

Create a sample slice of rooms containing three rooms.

Test two cases:

- Checking for an existing room ("Room2").
- Checking for a non-existing room ("Room4").

The first case, expect the function to return true as the room exists in the slice.
The second case, expect the function to return false as the room does not exist in the slice.
