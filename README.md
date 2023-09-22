# LEMIN TASK

## An Ant Movement Optimization Program

### Introduction

The task is to create a program that optimizes the movement of ants in a graph-based environment. The program should read input data from a file, construct a graph, find all optimal paths for ants to move through the graph, allocate ants to paths, and output the movement of ants in a human-readable format.

### Requirements

- Read Input File:

Read input data from a specified file with extension .txt
The file requires a minimum of 6 lines of data
The input data must include:
The number of ants. Positive integer on the first line.
A start room and an end room.
All rooms must have unique coordinates. First quadrant, with positive x and y coordinates.
A list of links between rooms.
Rooms are not permitted to link to themselves

- Graph Construction:

Construct a graph representation from the provided input data.
Each room has a unique name and unique set of coordinates (X, Y).
Rooms are connected by links, indicated by a hyphen/dash symbol '-'

Optimal Path Finding:

Implement a path-finding algorithm to find all possible routes between the start and end rooms.
Determine the optimal paths based on the number of ants and path lengths.

- Ant Allocation:

Allocate ants to paths in a way that minimizes the total travel time.
Each ant must traverse only one path from the start room to the end room.

- Output:

Print the movement of all the ants along the optimal paths.
Implement two output options:
Print the movements of the ants over time steps.
Print the migration that multiple ants can move in a single go/turn.

- Code Organization:

Organize the code into modular functions with clear responsibilities.
Use meaningful variable and function names.
Implement error handling and informative error messages.

- Testing and Validation:

Implement unit tests to validate the functions.
Provide example input files and expected output for testing purposes.

### Approach

BSF - Breath First Search : Graph Data Structures

### Resources

- [lem-in task](https://learn.01founders.co/intra/london/div-01/lem-in?event=138)

- [lem-in audit questions](https://github.com/01-edu/public/tree/master/subjects/lem-in)

- [LEM-IN: by Jamie Dawson](https://medium.com/@jamierobertdawson/lem-in-finding-all-the-paths-and-deciding-which-are-worth-it-2503dffb893)

- [LEM-IN: by Maksim Mikhailov](https://mer.pw/projects/lem-in)

- [Graph Data Structures: Junmin Lee](https://www.youtube.com/watch?v=JPD1OVgoa0Q)

- [Golang Binary Search Tree](https://www.youtube.com/watch?v=D3XlKCii7L4)

- [Dijkstra Shortest Path Algorithm](https://www.youtube.com/watch?v=pVfj6mxhdMw)

- [Edmonds-Karp Algorithm](https://youtu.be/RppuJYwlcI8)

### Setup

Input Format:

```
<number_of_ants>
<initiate ##start>
<start_room_name> <x_coord> <y_coord>
<initiate ##end>
<end_room_name> <x_coord> <y_coord>
<room_name_1> <x_coord> <y_coord>
<room_name_2> <x_coord> <y_coord>
...
<start_room_name>-<room_name_1>
<room_name_1>-<room_name_2>
<room_name_2>-<end_room_name>
...
```

Output Format:

Ant movements should be printed in a human-readable format.
Each ant is preceded by a capital 'L' character.
For each time step, print the room each ant is in.
The output format should show the steps and the go/turn migration.
Example Input and Output based on the data provided:

```
5
##start
start 0 0
##end
end 2 2
room1 0 1
room2 1 0
room3 2 1
start-room1
room1-room2
room2-room3
room3-end

Example Output:
(Time steps format)

L1-room1
L1-room2 L2-room1
L1-room3 L2-room2 L3-room1
L1-end   L2-room3 L3-room2 L4-room1
         L2-end   L3-room3 L4-room2 L5-room1
                  L3-end   L4-room3 L5-room2
                           L4-end   L5-room3
                                    L5-end
```

The total number of horizontal rows indicates the amount of goes/turns needed for all of the ants to reach the end room.
The vertical columns indicates the amount of steps each ant takes to reach the end room.

### How To Run

```
go run main.go <filename>
```

## TESTS

This document provides an explanation of the test cases for all functions in the `lemin` package.

[link to Tests.md file](/errhand/Tests.md)

### Testing Go Code and Generating Reports :

Testing the Go code is crucial to ensure its correctness and reliability. In addition to running tests, generating test reports helps to understand how much of the code is covered by the tests.

### Functions used in program

[link to Functions.md file](/functions/Functions.md)
