package lemin

import (
	"errors"
	"strconv"
	"strings"
)

// lc (line content) data

// InputData extracts relevant data from the input file and populates the Input struct
func InputData(data *Input, dataFile []string) error {
	for i, lc := range dataFile {
		switch {
		case i == 0:
			ants, err := strconv.Atoi(lc)
			if err != nil || ants < 1 { //must be at least 1 ant
				return errors.New("ERR: invalid number of Ants")
			}
			data.Ants = ants
		case lc == "##start" && i != len(dataFile)-1: //check ##start is present and not last line
			if err := IsValidRoom(dataFile[i+1], "start", data); err != nil {
				return err
			}
		case lc == "##end" && i != len(dataFile)-1: //check ##end is present and not last line 
			if err := IsValidRoom(dataFile[i+1], "end", data); err != nil {
				return err
			}
		case strings.Contains(lc, " "): //3 fields of data with separated by a space
			if err := IsValidRoom(lc, "", data); err != nil {
				return err
			}
		case strings.Contains(lc, "-"):
			node := strings.Split(lc, "-") //split line using dash (-) as the delimiter to establish room names
			if len(node) != 2 { //check for exactly 2 rooms only
				return errors.New("ERR: invalid link between rooms")
			}
			data.Links = append(data.Links, node)
		}
	}
	return nil //input data successfully read from file
}
