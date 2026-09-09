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

	if tape == nil {
		return fmt.Errorf("missing MAP tape")
	}
	if len(tape.Cell) < 3 {
		return fmt.Errorf("invalid MAP record - missing required parameters %d", len(tape.Cell))
	}

	if len(tape.Cell) > 0 && tape.Cell[0].S != nil && *tape.Cell[0].S == Prefix {
		command, err := cellKey(tape.Cell[1])
		if err != nil {
			return err
		}
		m[Cmd] = command

		switch m[Cmd] {
		case Delete:
			return m.delete(tape.Cell)
		case Add:
			return m.add(tape.Cell)
		case Remove:
			return m.remove(tape.Cell)
		case Set:
			return m.set(tape.Cell)
		case Select:
			return m.selecter(tape.Cell)
		}
	}
	return nil
}
