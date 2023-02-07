package magic

import (
	"testing"

	"github.com/bitcoinschema/go-bpu"
)

func TestNewFromTape(t *testing.T) {
	app := "app"
	myapp := "myapp"
	tape := bpu.Tape{
		Cell: []bpu.Cell{
			{
				S: &Prefix,
			}, {
				S: &Set,
			}, {
				S: &app,
			}, {
				S: &myapp,
			},
		},
	}

	tx, err := NewFromTape(&tape)
	if err != nil {
		t.Errorf("Failed to create new magic from tape")
	}

	if tx["app"] != &myapp {
		t.Errorf("Unexpected output %+v %s", tx, tx["app"])
	}
}
