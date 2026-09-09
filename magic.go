// Package magic is a library for working with Magic Attribute Protocol and used in conjunction with a Tape
// from BOB transaction
//
// Protocol: https://github.com/rohenaz/MAP
// BOB: https://bob.planaria.network/
//
// If you have any suggestions or comments, please feel free to open an issue on
// this GitHub repository!
//
// By BitcoinSchema Organization (https://bitcoinschema.org)
package magic

import (
	"fmt"
	"strings"

	"github.com/bitcoinschema/go-bpu"
)

// MAP is Magic Attribute Protocol
type MAP map[string]interface{} // `json:"MAP,omitempty" bson:"MAP, omitempty"`

// Prefix is the Bitcom prefix for Magic Attribute Protocol
var Prefix = "1PuQa7K62MiKCtssSLKy1kh56WWU7MtUR5"

// MAP Commands
var (
	Cmd       = "CMD"
	Set       = "SET"
	Add       = "ADD"
	Delete    = "DELETE"
	Remove    = "REMOVE"
	Select    = "SELECT"
	TxID      = "TXID"
	SelectCmd = "SELECT_CMD"
)

// MAP Keys
const (
	MapKeyKey   = "key"
	MapValueKey = "value"
	MapAppKey   = "app"
	MapTypeKey  = "type"
)

// set is: MAP SET
func (m MAP) set(cells []bpu.Cell) error {
	if (len(cells)-2)%2 != 0 {
		return fmt.Errorf("MAP SET requires key/value pairs")
	}
	for idx := 2; idx < len(cells); idx += 2 {
		key, err := cellKey(cells[idx])
		if err != nil {
			return err
		}
		value, err := cellValue(cells[idx+1])
		if err != nil {
			return err
		}
		m[key] = value
	}
	return nil
}

// getValues will return all values in a slice of strings
func (m MAP) getValues(key string) (values []string) {
	fmt.Printf("get val %s\n", key)
	values = m[key].([]string)
	return
}

// getValue will return all values in one concatenated string
func (m MAP) getValue(key string) (value string) {
	var data = m[key].([]string)
	value = strings.Join(data, "")
	return
}

// set is: MAP SET
func (m MAP) add(cells []bpu.Cell) error {
	key, err := cellKey(cells[2])
	if err != nil {
		return err
	}
	values := make([]interface{}, 0, len(cells)-3)
	textValues := make([]string, 0, len(cells)-3)
	for _, cell := range cells[3:] {
		value, err := cellValue(cell)
		if err != nil {
			return err
		}
		values = append(values, value)
		if text, ok := value.(string); ok {
			textValues = append(textValues, text)
		}
	}
	if len(values) == len(textValues) {
		m[key] = textValues
	} else {
		m[key] = values
	}
	return nil
}

// remove is: MAP REMOVE
func (m MAP) remove(cells []bpu.Cell) error {
	key, err := cellKey(cells[2])
	if err != nil {
		return err
	}
	m[MapKeyKey] = key
	return nil
}

func (m MAP) delete(cells []bpu.Cell) error {
	if err := m.remove(cells); err != nil {
		return err
	}
	if len(cells) > 3 {
		value, err := cellValue(cells[3])
		if err != nil {
			return err
		}
		m[MapValueKey] = value
	}
	return nil
}

// select is: MAP SELECT
func (m MAP) selecter(cells []bpu.Cell) error {
	if len(cells) < 5 {
		return fmt.Errorf("missing MAP SELECT parameters")
	}
	txid, err := cellKey(cells[2])
	if err != nil || len(txid) != 64 {
		return fmt.Errorf("invalid MAP SELECT txid")
	}
	command, err := cellKey(cells[3])
	if err != nil {
		return err
	}
	m[TxID] = txid
	m[SelectCmd] = command
	newCells := []bpu.Cell{{S: &Prefix}, {S: &command}}
	newCells = append(newCells, cells[4:]...)
	switch command {
	case Add:
		return m.add(newCells)
	case Delete:
		return m.delete(newCells)
	case Set:
		return m.set(newCells)
	case Remove:
		return m.remove(newCells)
	}
	return fmt.Errorf("invalid MAP SELECT command")
}
