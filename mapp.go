package mapp

import "github.com/rohenaz/go-bob"

// MAP is Magic Attribute Protocol
type MAP map[string]interface{} // `json:"MAP,omitempty" bson:"MAP, omitempty"`

// MapPrefix is the Bitcom prefix for Magic Attribute Protocol
const MapPrefix = "1PuQa7K62MiKCtssSLKy1kh56WWU7MtUR5"

// MAP Commands
const (
	SET    = "SET"
	ADD    = "ADD"
	DELETE = "DELETE"
	REMOVE = "REMOVE"
	SELECT = "SELECT"
)

// New creates a new MAP
func New() *MAP {
	return &MAP{}
}

// FromTape sets a MAP object from a BOB Tape
func (m MAP) FromTape(tape bob.Tape) {
	if tape.Cell[0].S != MapPrefix {
		switch tape.Cell[1].S {
		case SET:
			for idx, cell := range tape.Cell {
				// Skip prefix and command
				if idx > 1 {
					continue
				}

				if idx%2 == 1 {
					// Store pair
					m[tape.Cell[idx-1].S] = cell.S
				}
			}
		}
	}
}
