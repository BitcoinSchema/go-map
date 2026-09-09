package magic

import (
	"encoding/base64"
	"reflect"
	"strings"
	"testing"

	"github.com/bitcoinschema/go-bpu"
)

func textCell(value string) bpu.Cell { return bpu.Cell{S: &value} }
func rawCell(value string) bpu.Cell  { return bpu.Cell{B: &value} }
func mapTape(command string, cells ...bpu.Cell) *bpu.Tape {
	return &bpu.Tape{Cell: append([]bpu.Cell{textCell(Prefix), textCell(command)}, cells...)}
}

func TestCellValueRepresentations(t *testing.T) {
	largeText := "large text 👋"
	empty := ""
	malformed := "!"
	cases := []struct {
		name string
		cell bpu.Cell
		want interface{}
		err  string
	}{
		{"large text", bpu.Cell{LS: &largeText}, largeText, ""},
		{"empty text", textCell(empty), "", ""},
		{"empty raw", rawCell(empty), "", ""},
		{"missing", bpu.Cell{}, nil, "missing MAP value"},
		{"malformed raw", rawCell("!"), nil, "invalid MAP base64 value"},
		{"malformed large raw", bpu.Cell{LB: &malformed}, nil, "invalid MAP base64 value"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got, err := cellValue(c.cell)
			if c.err != "" {
				if err == nil || !strings.Contains(err.Error(), c.err) {
					t.Fatalf("got %v, want %s", err, c.err)
				}
				return
			}
			if err != nil || !reflect.DeepEqual(got, c.want) {
				t.Fatalf("got %#v, %v; want %#v", got, err, c.want)
			}
		})
	}
}

func TestMalformedMAPRecords(t *testing.T) {
	key := textCell("msg")
	bad := rawCell("!")
	binary := rawCell(base64.StdEncoding.EncodeToString([]byte{0xff, 0x80}))
	txid := textCell(strings.Repeat("a", 64))
	cases := []struct {
		name string
		tape *bpu.Tape
		err  string
	}{
		{"nil tape", nil, "missing MAP tape"},
		{"short tape", &bpu.Tape{}, "missing required parameters"},
		{"missing command", &bpu.Tape{Cell: []bpu.Cell{textCell(Prefix), {}, key}}, "missing MAP value"},
		{"binary command", &bpu.Tape{Cell: []bpu.Cell{textCell(Prefix), binary, key}}, "MAP key must be UTF-8 text"},
		{"set binary key", mapTape(Set, binary, textCell("value")), "MAP key must be UTF-8 text"},
		{"set malformed value", mapTape(Set, key, bad), "invalid MAP base64 value"},
		{"set missing value cell", mapTape(Set, key, bpu.Cell{}), "missing MAP value"},
		{"add malformed key", mapTape(Add, bad), "invalid MAP base64 value"},
		{"add malformed value", mapTape(Add, key, bad), "invalid MAP base64 value"},
		{"remove binary key", mapTape(Remove, binary), "MAP key must be UTF-8 text"},
		{"delete malformed key", mapTape(Delete, bad), "invalid MAP base64 value"},
		{"delete malformed value", mapTape(Delete, key, bad), "invalid MAP base64 value"},
		{"select short", mapTape(Select, txid), "missing MAP SELECT parameters"},
		{"select short txid", mapTape(Select, textCell("short"), textCell(Remove), key), "invalid MAP SELECT txid"},
		{"select malformed txid", mapTape(Select, bad, textCell(Remove), key), "invalid MAP SELECT txid"},
		{"select malformed command", mapTape(Select, txid, bad, key), "invalid MAP base64 value"},
		{"select unknown command", mapTape(Select, txid, textCell("UNKNOWN"), key), "invalid MAP SELECT command"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			_, err := NewFromTape(c.tape)
			if err == nil || !strings.Contains(err.Error(), c.err) {
				t.Fatalf("got %v, want %s", err, c.err)
			}
		})
	}
}

func TestSELECTPreservesCommandValues(t *testing.T) {
	txid := strings.Repeat("a", 64)
	encoded := base64.StdEncoding.EncodeToString([]byte{0xff, 0, 0x80})
	value := BinaryValue{B: encoded}
	cases := []struct {
		command string
		cells   []bpu.Cell
		want    MAP
	}{
		{Set, []bpu.Cell{textCell("msg"), rawCell(encoded)}, MAP{"msg": value}},
		{Add, []bpu.Cell{textCell("msg"), textCell("hello"), rawCell(encoded)}, MAP{"msg": []interface{}{"hello", value}}},
		{Delete, []bpu.Cell{textCell("msg"), rawCell(encoded)}, MAP{MapKeyKey: "msg", MapValueKey: value}},
		{Remove, []bpu.Cell{textCell("msg")}, MAP{MapKeyKey: "msg"}},
	}
	for _, c := range cases {
		t.Run(c.command, func(t *testing.T) {
			cells := append([]bpu.Cell{textCell(txid), textCell(c.command)}, c.cells...)
			got, err := NewFromTape(mapTape(Select, cells...))
			c.want[Cmd] = Select
			c.want[TxID] = txid
			c.want[SelectCmd] = c.command
			if err != nil || !reflect.DeepEqual(got, c.want) {
				t.Fatalf("got %#v, %v; want %#v", got, err, c.want)
			}
		})
	}
}

func TestFromTapeIgnoresOtherProtocols(t *testing.T) {
	for _, prefix := range []bpu.Cell{textCell("other-protocol"), {}} {
		got, err := NewFromTape(&bpu.Tape{Cell: []bpu.Cell{prefix, textCell(Set), textCell("key"), textCell("value")}})
		if err != nil || len(got) != 0 {
			t.Fatalf("non-MAP tape produced %#v, %v", got, err)
		}
	}
}
