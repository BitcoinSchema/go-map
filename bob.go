package magic

import (
	"fmt"

	"github.com/bitcoinschema/go-bpu"
)

// NewFromTape takes a tape and returns a new MAP
func NewFromTape(tape *bpu.Tape) (magicTx MAP, err error) {
	magicTx = make(MAP)
	err = magicTx.FromTape(tape)
	return
}

// FromTape sets a MAP object from a BOB Tape
func (m MAP) FromTape(tape *bpu.Tape) error {

	if len(tape.Cell) < 3 {
		return fmt.Errorf("invalid MAP record - missing required parameters %d", len(tape.Cell))
	}

	if len(tape.Cell) > 0 && tape.Cell[0].S != nil && *tape.Cell[0].S == Prefix {
		m[Cmd] = *tape.Cell[1].S

		switch m[Cmd] {
		case Delete:
			m.delete(tape.Cell)
		case Add:
			m.add(tape.Cell)
		case Remove:
			m.remove(tape.Cell)
		case Set:
			m.set(tape.Cell)
		case Select:
			m.selecter(tape.Cell)
		}
	}
	return nil
}
