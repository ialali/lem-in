// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"strconv"
// 	"strings"
// )

// // initialize a Node struct for the numerr := graph.Render(chart.PNG, os.Stdout)ber of ants, all given rooms(nodes), that have a Name, X-Y coordinates, and connections to other nodes
// type AntFarm struct {
// 	Ants int
// 	RmS	string
// 	RmE	string
// 	RmGen	[]string //name of Node
// 	X, Y int
// 	Coords [][]string //coordinates of each node
// 	Connections [][]string //connections between nodes
// }

// // 1. CHECK FILENAME HAS .txt EXTENSION
// func hasTXTExtension(filename string) bool {
// 	return strings.HasSuffix(filename, ".txt")
// }
// // 2. CHECK FILE CONTAINS A MINIMUM OF 6 LINES
// func fileContainsMinLines(filePath string, minLines int) (bool, error) {
// 	file, err := os.Open(filePath)
// 	if err != nil {
// 			return false, err
// 	}
// 	defer file.Close()

// 	scanner := bufio.NewScanner(file)
// 	lineCount := 0

// 	for scanner.Scan() {
// 			lineCount++
// 			// Break out of the loop early if the minimum line count is reached.
// 			if lineCount >= minLines {
// 					return true, nil
// 			}
// 	}
// 	if err := scanner.Err(); err != nil {
// 			return false, err
// 	}
// 	// If the minimum line count is not reached, return false.
// 	return false, nil
// }
// // 3. CHECK FIRST LINE IS A POSITIVE INTEGER
// func firstLineContainsPositiveInteger(filePath string) (bool, error) {
// 	file, err := os.Open(filePath)
// 	if err != nil {
// 			return false, err
// 	}
// 	defer file.Close()

// 	scanner := bufio.NewScanner(file)
// 	if scanner.Scan() {
// 			firstLine := scanner.Text()
// 			// Trim leading and trailing whitespace from the line.
// 			firstLine = strings.TrimSpace(firstLine)

// 			// Attempt to convert the first line to an integer.
// 			number, err := strconv.Atoi(firstLine)
// 			if err != nil {
// 					return false, nil // The first line is not a valid integer.
// 			}
// 			// Check if the number is positive.
// 			if number > 0 {
// 					return true, nil // The first line contains a positive integer.
// 			}
// 	}

// 	if err := scanner.Err(); err != nil {
// 			return false, err
// 	}

// 	// If the first line doesn't contain a positive integer, return false.
// 	return false, nil
// }
// // 4. CHECK FILE CONTAINS ##start AND ##end
// func fileContainsStartAndEnd(filePath string) (bool, error) {
// 	file, err := os.Open(filePath)
// 	if err != nil {
// 			return false, err
// 	}
// 	defer file.Close()

// 	scanner := bufio.NewScanner(file)
// 	foundStart := false
// 	foundEnd := false

// 	for scanner.Scan() {
// 			line := scanner.Text()
// 			// Check if the line contains "##start" and "##end" using the strings.Contains function.
// 			if strings.Contains(line, "##start") {
// 					foundStart = true
// 			}
// 			if strings.Contains(line, "##end") {
// 					foundEnd = true
// 			}

// 			// If both "##start" and "##end" are found, break out of the loop.
// 			if foundStart && foundEnd {
// 					break
// 			}
// 	}

// 	if err := scanner.Err(); err != nil {
// 			return false, err
// 	}

// 	// Check if both "##start" and "##end" were found in the file.
// 	return foundStart && foundEnd, nil
// }
// // 5. CHECK FOR DUPLICATE ##start AND ##end
// func noDuplicateStartAndEnd(filePath string) (bool, error) {
// 	file, err := os.Open(filePath)
// 	if err != nil {
// 			return false, err
// 	}
// 	defer file.Close()

// 	scanner := bufio.NewScanner(file)
// 	startCount := 0
// 	endCount := 0

// 	for scanner.Scan() {
// 			line := scanner.Text()
// 			// Check if the line contains "##start" and "##end" using the strings.Contains function.
// 			if strings.Contains(line, "##start") {
// 					startCount++
// 			}
// 			if strings.Contains(line, "##end") {
// 					endCount++
// 			}
// 			// If there are duplicates of "##start" or "##end," return false.
// 			if startCount > 1 || endCount > 1 {
// 					return false, nil
// 			}
// 	}

// 	if err := scanner.Err(); err != nil {
// 			return false, err
// 	}
// 	// If there are no duplicates of "##start" and "##end," return true.
// 	return true, nil
// }

// // 6. CHECK IF EACH LINE BEGINS WITH L
// func checkLinesForL(filePath string) (bool, error) {
// 	file, err := os.Open(filePath)
// 	if err != nil {
// 			return false, err
// 	}
// 	defer file.Close()

// 	scanner := bufio.NewScanner(file)
// 	lineCount := 0

// 	for scanner.Scan() {
// 			line := scanner.Text()
// 			lineCount++
// 			// Check if the line begins with the letter "L".
// 			if line != "" && line[0] == 'L' {
// 					fmt.Printf("Error: Line %d starts with 'L': %s\n", lineCount, line)
// 					// return false, nil
// 					// exit program
// 					os.Exit(1)
// 			}
// 	}
// 	if err := scanner.Err(); err != nil {
// 			return false, err
// 	}
// 	// If all lines pass the check, return true.
// 	return true, nil
// }
// // 7. CHECK IF EACH LINE BEGINS WITH HASH #
// func checkLinesForHash(filePath string) (bool, error) {
// 	file, err := os.Open(filePath)
// 	if err != nil {
// 			return false, err
// 	}
// 	defer file.Close()

// 	scanner := bufio.NewScanner(file)
// 	// Initialize a flag to ignore lines starting with '##start' or '##end'.
// 	ignoreNext := false

// 	// Check each line for the '#' character at the beginning.
// 	lineCount := 0
// 	for scanner.Scan() {
// 			line := scanner.Text()
// 			lineCount++

// 			// Check if the line begins with '##start' or '##end'.
// 			if strings.HasPrefix(line, "##start") || strings.HasPrefix(line, "##end") || strings.HasPrefix(line, "#comment"){
// 				ignoreNext = true
// 				continue
// 			}

// 			// Check if the line starts with '#' when not ignored.
// 			if !ignoreNext && strings.HasPrefix(line, "#") {
// 				fmt.Printf("Error: Line %d starts with '#': %s\n", lineCount, line)
// 				// exit program
// 				os.Exit(1)
// 			}
// 			// Reset the flag to false.
// 			ignoreNext = false

// 			// // Check if the line begins with the hash "#".
// 			// if line != "" && line[0] == '#' {
// 			// 		fmt.Printf("Error: Line %d starts with '#': %s\n", lineCount, line)
// 			// 		return false, nil
// 			// }
// 	}
// 	if err := scanner.Err(); err != nil {
// 		fmt.Println("Error: hash# at beginning", err)
// 		return false, err
// 	}
// 	// If all lines pass the check, return true.
// 	return true, nil
// }
// //8. valid Room data
// func validRoom(fiePath string) {
// 	// Open the input file.
// 	file, err := os.Open(filePath)
// 	if err != nil {
// 			fmt.Println("Error:", err)
// 			return
// 	}
// 	defer file.Close()
// 	// Create a scanner to read lines from the file.
// 	scanner := bufio.NewScanner(file)

// 	// Iterate through each line of the file.
// 	lineNum := 0
// 	for scanner.Scan() {
// 			line := scanner.Text()
// 			lineNum++

// 			//Check if the line contains '#comment' if true, skip
// 			if strings.TrimSpace(line) == "#comment" {
// 					continue
// 			}
// 			// Split the line into fields using space as the separator.
// 			fields := strings.Fields(line)
// 			// Check if the line has exactly 3 fields.
// 			if len(fields) != 3 {
// 					fmt.Printf("Line %d does not contain 3 fields: %s\n", lineNum, line)
// 					continue
// 			}
// 			//Parse fields 2 and 3 as integers
// 			field2, _ := strconv.Atoi(fields[1])
// 			field3, _ := strconv.Atoi(fields[2])

// 			//Check if fields 2 & 3 are positive integers
// 			if field2 < 0 || field3 < 0 {
// 				fmt.Printf("Line %d, Fields 2 and 3 must be positive integers: %s\n", lineNum, line)
// 		} else {
// 				fmt.Printf("Line %d does contain 3 valid fields with 2 positive integers\n", lineNum)
// 		}
// 	}
// 	// Check for any errors during scanning.
// 	if err := scanner.Err(); err != nil {
// 			fmt.Println("Error: 3 arguments needed", err)
// 	}
// }

// func main() {
// 	// 1. Check if there are at least two command line arguments (including the program name).
// 	if len(os.Args) < 2 {
// 			fmt.Println("Include an external file with extension .txt")
// 			return //exit program
// 	}
// 	// Get the filename from the command line arguments.
// 	filename := os.Args[1]

// 	if hasTXTExtension(filename) {
// 			fmt.Println(filename, "has a .txt extension")
// 	} else {
// 			fmt.Println(filename, "does not have a .txt extension")
// 			return //exit program
// 	}

// 	// 2. Check file has 6 lines of data
// 	filePath := filename // Replace with the path to your file 'your_file.txt'
// 	minLines := 6

//     containsMinLines, err := fileContainsMinLines(filePath, minLines)
//     if err != nil {
//         fmt.Println("Error: 6 lines needed", err)
//         return
//     }
// 		if containsMinLines {
//         fmt.Printf("The file contains a minimum of %d lines.\n", minLines)
//     } else {
//         fmt.Printf("The file does not contain a minimum of %d lines.\n", minLines)
// 				return //exit program
//     }
// 	// 3. Check first line is a positive integer
// 	containsPositiveInteger, err := firstLineContainsPositiveInteger(filePath)
//     if err != nil {
//         fmt.Println("Error: first line requires positive integer", err)
//         return
//     }
// 		if containsPositiveInteger {
//         fmt.Println("The first line contains a positive integer.")
//     } else {
//         fmt.Println("The first line does not contain a positive integer.")
// 				return //exit program
//     }

// 	// 4. Check file contains '##start' and '##end'
// 	containsStartAndEnd, err := fileContainsStartAndEnd(filePath)
//     if err != nil {
//         fmt.Println("Error: need both ##start AND ##end", err)
//         return
//     }
// 		if containsStartAndEnd {
//         fmt.Println("The file contains ##start and ##end.")
//     } else {
//         fmt.Println("File needs to contain both ##start AND ##end.")
// 				return //exit program
//     }

// 	// 5. Check for Duplicate ##start and ##end
// 	noDuplicates, err := noDuplicateStartAndEnd(filePath)
//     if err != nil {
//         fmt.Println("Error: duplicate of ##start ##end", err)
//         return
//     }
// 		if noDuplicates {
//         fmt.Println("There are no duplicates of ##start and ##end in the file.")
//     } else {
//         fmt.Println("Error: duplicates of ##start and/or ##end in the file.")
// 				return //exit program
//     }

// 	// 6. Check if any line begins with L
// 	validFile, err := checkLinesForL(filePath)
//     if err != nil {
//         fmt.Println("Error: line begins with L", err)
//         return //exit program
//     }
// 		if validFile {
//         fmt.Println("All lines in the file do not start with 'L'.")
//     }
// 	// 7. Check if any line begins with hash '#'
// 	validPrefix, err := checkLinesForHash(filePath)
// 		if err != nil {
// 			fmt.Println("Error: line begins with hash '#'", err)
// 			return //exit program
// 		}
// 		if validPrefix {
// 			fmt.Println("All lines in file do not begin with hash '#'")
// 		}
// 	// 8.) Check lines for Valid Room data 3 fields
// 	{
// 		filePath = filename
// 		validRoom(filePath)
// 	}
// }