package magic

import (
	"encoding/base64"
	"fmt"
	"unicode/utf8"

	"github.com/bitcoinschema/go-bpu"
)

// BinaryValue preserves a non-text MAP value using the same base64 field as BOB.
// Text values remain strings; JSON and BSON never contain invalid UTF-8 strings.
type BinaryValue struct {
	B string `json:"b" bson:"b"`
}

func cellValue(cell bpu.Cell) (interface{}, error) {
	var data []byte
	encoded := cell.B
	if encoded == nil {
		encoded = cell.LB
	}
	if encoded != nil {
		var err error
		data, err = base64.StdEncoding.DecodeString(*encoded)
		if err != nil {
			return nil, fmt.Errorf("invalid MAP base64 value: %w", err)
		}
	} else {
		text := cell.S
		if text == nil {
			text = cell.LS
		}
		if text == nil {
			return nil, fmt.Errorf("missing MAP value")
		}
		data = []byte(*text)
	}
	if utf8.Valid(data) {
		return string(data), nil
	}
	return BinaryValue{B: base64.StdEncoding.EncodeToString(data)}, nil
}

func cellKey(cell bpu.Cell) (string, error) {
	value, err := cellValue(cell)
	if err != nil {
		return "", err
	}
	text, ok := value.(string)
	if !ok {
		return "", fmt.Errorf("MAP key must be UTF-8 text")
	}
	return text, nil
}
