package magic

import (
	"fmt"
	"log"

	"github.com/bitcoinschema/go-bob"
)

// MAP is Magic Attribute Protocol
type MAP map[string]interface{} // `json:"MAP,omitempty" bson:"MAP, omitempty"`

// Prefix is the Bitcom prefix for Magic Attribute Protocol
const Prefix = "1PuQa7K62MiKCtssSLKy1kh56WWU7MtUR5"

// MAP Commands
const (
	Cmd       = "CMD"
	Set       = "SET"
	Add       = "ADD"
	Delete    = "DELETE"
	Remove    = "REMOVE"
	Select    = "SELECT"
	TxID      = "TXID"
	SelectCmd = "SELECT_CMD"
)

// MAP SET
func (m MAP) set(cells []bob.Cell) {
	for idx, cell := range cells {
		// Skip prefix (0) and command (1)
		if idx < 2 {
			continue
		}

		if idx%2 == 1 {
			m.setValue(cells[idx-1].S, cell.S)
		}
	}
}

func (m MAP) setValue(key string, val string) {
	log.Printf("Setting value %s instead of %s in %+v", val, m[key], m)
	m[key] = val
	return
}

func (m MAP) setValues(key string, val []string) {
	log.Println("Setting values", val)
	m[key] = val
	return
}

// value getter
func (m MAP) getValue(key string) (val string) {
	return fmt.Sprintf("%s", m[key])
}

// MAP ADD
func (m MAP) add(cells []bob.Cell) {
	var keyValues []string
	keyName := cells[2].S
	for idx, cell := range cells {
		// Skip prefix (0), command (1) and keyName (2)
		if idx < 3 {
			continue
		}
		keyValues = append(keyValues, cell.S)
	}
	m.setValues(keyName, keyValues)
}

func (m MAP) remove(cells []bob.Cell) {
	// Skip prefix (0) and command (1)
	m.setValue("key", cells[2].S)
}

func (m MAP) delete(cells []bob.Cell) {
	// Skip prefix (0) and command (1)
	m.setValue("key", cells[2].S)
	m.setValue("value", cells[3].S)
}

// NewFromTape takes a tape and returns a new MAP
func NewFromTape(tape *bob.Tape) (magicTx *MAP, err error) {
	magicTx = new(MAP)
	err = magicTx.FromTape(tape)
	return
}

// FromTape sets a MAP object from a BOB Tape
func (m MAP) FromTape(tape *bob.Tape) error {

	if len(tape.Cell) < 3 {
		return fmt.Errorf("Invalid MAP record. Missing require parameters %d", len(tape.Cell))
	}

	if tape.Cell[0].S == Prefix {
		m.setValue(Cmd, tape.Cell[1].S)

		switch m.getValue(Cmd) {
		case Delete:
			m.delete(tape.Cell)
		case Add:
			m.add(tape.Cell)
		case Remove:
			m.remove(tape.Cell)
		case Set:
			m.set(tape.Cell)
		case Select:
			if len(tape.Cell) < 5 {
				return fmt.Errorf("Missing required parameters in MAP SELECT statement. Cell length: %d", len(tape.Cell))
			}
			if len(tape.Cell[2].S) != 64 {
				return fmt.Errorf("MAP syntax error. Invalid Txid in SELECT command: %d", len(tape.Cell))
			}
			m.setValue(TxID, tape.Cell[2].S)
			m.setValue(SelectCmd, tape.Cell[3].S)

			// Build new command from SELECT
			newCells := []bob.Cell{{S: Prefix}, {S: tape.Cell[3].S}}
			newCells = append(newCells, tape.Cell[4:]...)
			switch SelectCmd {
			case Add:
				m.add(newCells)
			case Delete:
				m.delete(newCells)
			case Set:
				m.set(newCells)
			case Remove:
				m.remove(newCells)
			}
		}
	}
	return nil
}
