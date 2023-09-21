package lemin

import (
	"errors"
	"strconv"
	"strings"
)

// vr (valid room)
// ser (start, end, regular room)

// IsValidRoom validates and processes a room entry
func IsValidRoom(vr, ser string, data *Input) error {
	node := strings.Split(vr, " ")
	if len(node) != 3 { //checks for 3 fields of data name, x, y
		return errors.New("ERR: invalid format of the room, IsValidRoom")
	}
	if strings.ContainsAny(node[0], " #") || node[0][0] == 'L' {
		return errors.New("ERR: invalid data format, the room can't start with L, # or space")
	}
	x, err := strconv.Atoi(node[1]) //convert field 2 to x coord
	if err != nil || x < 0 {
		return errors.New("ERR: invalid data format, the coord X must be a non-negative number")
	}
	y, err := strconv.Atoi(node[2]) //convert field 3 to y coord
	if err != nil || y < 0 {
		return errors.New("ERR: invalid data format, the coord Y must be a non-negative number")
	}
	nodeCoord := []string{node[1], node[2]} //node coord slice to hold x & y values
	if !IsValidCoords(nodeCoord, data) {
		return errors.New("ERR: invalid data format, room with such coordinates already exists")
	}
	switch ser {
	case "start":
		data.StartR = node[0]
	case "end":
		data.EndR = node[0]
	default:
		data.Rooms = append(data.Rooms, node[0])
		data.Coords = append(data.Coords, []string{node[1], node[2]})
	}
	return nil
}
//StartR = StartRm start room
//EndR = EndRm end room
//Rooms = GenRm general room
