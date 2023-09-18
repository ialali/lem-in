package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Constants for error messages
const (
	errNoTxtExtension        = "Include an external file with extension .txt"
	errNotEnoughLines        = "The file does not contain a minimum of 6 lines."
	errNotPositiveInteger    = "The first line does not contain a positive integer."
	errMissingStartAndEnd    = "File needs to contain both ##start AND ##end."
	errDuplicateStartAndEnd  = "Error: duplicates of ##start and/or ##end in the file."
	errLineBeginsWithL       = "Error: line begins with 'L'"
	errLineBeginsWithHash    = "Error: line begins with '#'"
	errLineDoesNotContain3   = "Line %d does not contain 3 fields: %s\n"
	errFieldsNotPositiveInts = "Line %d does not contain 3 valid fields with 2 integers\n"
)

// Check if a file has a .txt extension
func hasTXTExtension(filename string) bool {
	return strings.HasSuffix(filename, ".txt")
}

// Check if a file contains a minimum number of lines
func fileContainsMinLines(filePath string, minLines int) (bool, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return false, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineCount := 0

	for scanner.Scan() {
		lineCount++
		if lineCount >= minLines {
			return true, nil
		}
	}

	if err := scanner.Err(); err != nil {
		return false, err
	}

	return false, nil
}

// Check if the first line of a file contains a positive integer
func firstLineContainsPositiveInteger(filePath string) (bool, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return false, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	if scanner.Scan() {
		firstLine := scanner.Text()
		firstLine = strings.TrimSpace(firstLine)

		number, err := strconv.Atoi(firstLine)
		if err != nil || number <= 0 {
			return false, nil
		}
		return true, nil
	}

	if err := scanner.Err(); err != nil {
		return false, err
	}

	return false, nil
}

// Check if a file contains both "##start" and "##end"
func fileContainsStartAndEnd(filePath string) (bool, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return false, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	foundStart := false
	foundEnd := false

	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "##start") {
			foundStart = true
		}
		if strings.Contains(line, "##end") {
			foundEnd = true
		}

		if foundStart && foundEnd {
			return true, nil
		}
	}

	if err := scanner.Err(); err != nil {
		return false, err
	}

	return false, nil
}

// Check for duplicate "##start" and "##end" in a file
func noDuplicateStartAndEnd(filePath string) (bool, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return false, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	startCount := 0
	endCount := 0

	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "##start") {
			startCount++
		}
		if strings.Contains(line, "##end") {
			endCount++
		}
		if startCount > 1 || endCount > 1 {
			return false, nil
		}
	}

	if err := scanner.Err(); err != nil {
		return false, err
	}

	return true, nil
}

// Check if any line in a file begins with 'L'
func checkLinesForL(filePath string) (bool, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return false, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if line != "" && line[0] == 'L' {
			return false, nil
		}
	}

	if err := scanner.Err(); err != nil {
		return false, err
	}

	return true, nil
}

// Check if any line in a file begins with '#' (excluding ##start and ##end)
func checkLinesForHash(filePath string) (bool, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return false, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	ignoreNext := false

	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "##start") || strings.HasPrefix(line, "##end") || strings.HasPrefix(line, "#comment") {
			ignoreNext = true
			continue
		}
		if !ignoreNext && strings.HasPrefix(line, "#") {
			return false, nil
		}
		ignoreNext = false
	}

	if err := scanner.Err(); err != nil {
		return false, err
	}

	return true, nil
}

// Main function
func main() {
	if len(os.Args) < 2 {
		fmt.Println(errNoTxtExtension)
		return
	}

	filename := os.Args[1]

	if !hasTXTExtension(filename) {
		fmt.Printf("%s does not have a .txt extension\n", filename)
		return
	}

	filePath := filename
	minLines := 6

	containsMinLines, err := fileContainsMinLines(filePath, minLines)
	if err != nil {
		fmt.Println(errNotEnoughLines, err)
		return
	}
	if !containsMinLines {
		fmt.Println(errNotEnoughLines)
		return
	}

	containsPositiveInteger, err := firstLineContainsPositiveInteger(filePath)
	if err != nil {
		fmt.Println(errNotPositiveInteger, err)
		return
	}
	if !containsPositiveInteger {
		fmt.Println(errNotPositiveInteger)
		return
	}

	containsStartAndEnd, err := fileContainsStartAndEnd(filePath)
	if err != nil {
		fmt.Println(errMissingStartAndEnd, err)
		return
	}
	if !containsStartAndEnd {
		fmt.Println(errMissingStartAndEnd)
		return
	}

	noDuplicates, err := noDuplicateStartAndEnd(filePath)
	if err != nil {
		fmt.Println(errDuplicateStartAndEnd, err)
		return
	}
	if !noDuplicates {
		fmt.Println(errDuplicateStartAndEnd)
		return
	}

	validFile, err := checkLinesForL(filePath)
	if err != nil {
		fmt.Println(errLineBeginsWithL, err)
		return
	}
	if !validFile {
		fmt.Println(errLineBeginsWithL)
		return
	}

	validPrefix, err := checkLinesForHash(filePath)
	if err != nil {
		fmt.Println(errLineBeginsWithHash, err)
		return
	}
	if !validPrefix {
		fmt.Println(errLineBeginsWithHash)
		return
	}

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineNum := 0

	for scanner.Scan() {
		line := scanner.Text()
		lineNum++

		if strings.TrimSpace(line) == "#comment" {
			continue
		}

		fields := strings.Fields(line)
		if len(fields) != 3 {
			fmt.Printf(errLineDoesNotContain3, lineNum, line)
			continue
		}

		field2, _ := strconv.Atoi(fields[1])
		field3, _ := strconv.Atoi(fields[2])

		if field2 <= 0 || field3 <= 0 {
			fmt.Printf(errFieldsNotPositiveInts, lineNum)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error:", err)
	}
}
