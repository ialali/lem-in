package lemin

import (
	"errors"
	"strconv"
	"strings"
)

// GetData extracts relevant data from the input file and populates the Input struct
func GetData(data *Input, dataFile []string) error {
	for i, v := range dataFile {
		switch {
		case i == 0:
			ants, err := strconv.Atoi(v)
			if err != nil || ants < 1 { // positive int for ants
				return errors.New("ERR: invalid data format, invalid number of Ants")
			}
			data.Ants = ants
		case v == "##start" && i != len(dataFile)-1: //chk initalisation of start-room
			if err := IsValidRoom(dataFile[i+1], "start", data); err != nil {
				return err
			}
		case v == "##end" && i != len(dataFile)-1: //chk initalisation of end-room
			if err := IsValidRoom(dataFile[i+1], "end", data); err != nil {
				return err
			}
		case strings.Contains(v, " "): //chk room for 3 components name, x, y
			if err := IsValidRoom(v, "", data); err != nil {
				return err
			}
		case strings.Contains(v, "-"): //chk link for 2 components and (hyphen/dash) between with rooms
			node := strings.Split(v, "-")
			if len(node) != 2 {
				return errors.New("ERR: invalid data format, invalid link between rooms")
			}
			data.Links = append(data.Links, node)
		}
	}
	return nil //no errors during data extraction
}
