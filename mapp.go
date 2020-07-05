package mapp

import (
	"fmt"

	"github.com/rohenaz/go-bob"
)

// MAP is Magic Attribute Protocol
type MAP map[string]interface{} // `json:"MAP,omitempty" bson:"MAP, omitempty"`

// MapPrefix is the Bitcom prefix for Magic Attribute Protocol
const MapPrefix = "1PuQa7K62MiKCtssSLKy1kh56WWU7MtUR5"

// MAP Commands
const (
	CMD        = "CMD"
	SET        = "SET"
	ADD        = "ADD"
	DELETE     = "DELETE"
	REMOVE     = "REMOVE"
	SELECT     = "SELECT"
	TXID       = "TXID"
	SELECT_CMD = "SELECT_CMD"
)

// New creates a new MAP
func New() *MAP {
	return &MAP{}
}

// MAP SET
func (m MAP) set(cells []bob.Cell) {
	for idx, cell := range cells {
		// Skip prefix and command
		if idx < 2 {
			continue
		}

		if idx%2 == 1 {
			m[cells[idx-1].S] = cell.S
		}
	}
}

// MAP ADD
func (m MAP) add(cells []bob.Cell) {
	for idx, cell := range cells {
		if idx < 2 {
			continue
		}
		var keyName string
		if idx == 4 {
			keyName = cell.S
			continue
		}

		m[keyName] = cell.S
	}
}

func (m MAP) remove(cell []bob.Cell) {
	// Since set is inverse of remove we can build with the same function
	m.set(cell)
}

func (m MAP) delete(cell []bob.Cell) {
	// Since add is inverse of delete we can build with the same function
	m.add(cell)
}

// FromTape sets a MAP object from a BOB Tape
func (m MAP) FromTape(tape bob.Tape) error {
	if tape.Cell[0].S != MapPrefix {
		m[CMD] = tape.Cell[1].S
		switch m[CMD] {
		case DELETE:
			fallthrough
		case ADD:
			m.add(tape.Cell)
		case REMOVE:
			fallthrough
		case SET:
			m.set(tape.Cell)
		case SELECT:
			if len(tape.Cell) < 5 {
				return fmt.Errorf("Missing required parameters in MAP SELECT statement. Cell length: %d", len(tape.Cell))
			}
			m[TXID] = tape.Cell[2].S
			m[SELECT_CMD] = tape.Cell[3].S

			// Build the command from SELECT
			newCells := []bob.Cell{}
			newCells[0] = bob.Cell{S: MapPrefix}
			newCells[1] = bob.Cell{S: SELECT_CMD}
			for idx, cell := range tape.Cell {
				if idx < 4 {
					continue
				}
				newCells[idx-2] = cell
			}
			switch m[SELECT_CMD] {
			case ADD:
				m.add(newCells)
			case DELETE:
				m.delete(newCells)
			case SET:
				m.set(newCells)
			case REMOVE:
				m.remove(newCells)
			}
		}
	}
}
