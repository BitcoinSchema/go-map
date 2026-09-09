package magic

import (
	"encoding/base64"
	"encoding/json"
	"github.com/bitcoinschema/go-bpu"
	"testing"
)

func TestBinaryMAPValuesPreserveBytes(t *testing.T) {
	key := "msg"
	binary := []byte{0xff, 0, 0x80, 0x3d, 0xd8}
	encoded := base64.StdEncoding.EncodeToString(binary)
	lossy := "replacement text must not win over raw bytes"
	tape := bpu.Tape{Cell: []bpu.Cell{{S: &Prefix}, {S: &Set}, {S: &key}, {B: &encoded, S: &lossy}}}
	m, err := NewFromTape(&tape)
	if err != nil {
		t.Fatal(err)
	}
	value, ok := m[key].(BinaryValue)
	if !ok || value.B != encoded {
		t.Fatalf("lost binary value: %#v", m[key])
	}
	data, err := json.Marshal(m)
	if err != nil {
		t.Fatal(err)
	}
	var decoded map[string]interface{}
	if err = json.Unmarshal(data, &decoded); err != nil {
		t.Fatal(err)
	}
	if decoded[key].(map[string]interface{})["b"] != encoded {
		t.Fatal("JSON did not preserve base64")
	}
	for _, command := range []string{Add, Delete} {
		tape.Cell[1].S = &command
		if _, err = NewFromTape(&tape); err != nil {
			t.Fatal(err)
		}
	}
	text := "hello 👋"
	text64 := base64.StdEncoding.EncodeToString([]byte(text))
	tape.Cell[1].S = &Set
	tape.Cell[3] = bpu.Cell{B: &text64}
	m, err = NewFromTape(&tape)
	if err != nil || m[key] != text {
		t.Fatalf("text changed: %#v %v", m, err)
	}
	tape.Cell[3] = bpu.Cell{LB: &encoded}
	m, err = NewFromTape(&tape)
	if err != nil || m[key].(BinaryValue).B != encoded {
		t.Fatal("large binary cell lost")
	}
	tape.Cell = tape.Cell[:3]
	if _, err = NewFromTape(&tape); err == nil {
		t.Fatal("accepted missing SET value")
	}
}
