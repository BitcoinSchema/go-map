package magic

import (
	"testing"

	"github.com/bitcoinschema/go-bob"
)

func contains(arr []string, str string) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}
	return false
}

func TestSelectDelete(t *testing.T) {
	tape := bob.Tape{
		Cell: []bob.Cell{
			{S: Prefix},
			{S: Select},
			{S: "a9a4387d2baa2edcc53ec040b3affbc38778e9dd876f9a47e5c767c785aacf76"},
			{S: Delete},
			{S: "keyName1"},
			{S: "something"},
		},
	}
	m := MAP{}
	m.FromTape(&tape)
	if m[Cmd] != Select || m["key"] != "keyName1" || m["value"] != "something" {
		t.Errorf("SELECT + DELETE Failed %s %+v", m[Cmd], m)
	}
}

func TestAdd(t *testing.T) {
	tape := bob.Tape{
		Cell: []bob.Cell{
			{S: Prefix},
			{S: Add},
			{S: "keyName"},
			{S: "something"},
			{S: "something else"},
		},
	}
	m := MAP{}
	m.FromTape(&tape)

	switch m["keyName"].(type) {
	case []string:
		if !contains(m["keyName"].([]string), "something") ||
			!contains(m["keyName"].([]string), "something else") {
			t.Errorf("ADD Failed %s", m["keyName1"])
		}
		break
	default:
		t.Errorf("ADD Failed %s", m["keyName1"])
	}
}

func TestDelete(t *testing.T) {
	tape := bob.Tape{
		Cell: []bob.Cell{
			{S: Prefix},
			{S: Delete},
			{S: "keyName"},
			{S: "something"},
		},
	}
	m := MAP{}
	m.FromTape(&tape)

	if m["key"] != "keyName" || m["value"] != "something" {
		t.Errorf("DELETE Failed %s %s", m["key"], m["value"])
	}

}

func TestSet(t *testing.T) {
	tape := bob.Tape{
		Cell: []bob.Cell{
			{S: Prefix},
			{S: Set},
			{S: "keyName1"},
			{S: "something"},
			{S: "keyName2"},
			{S: "something else"},
		},
	}
	m := MAP{}
	m.FromTape(&tape)
	if m["keyName1"] != "something" {
		t.Errorf("SET Failed %s", m["keyName1"])
	}
}

func TestRemove(t *testing.T) {
	tape := bob.Tape{
		Cell: []bob.Cell{
			{S: Prefix},
			{S: Remove},
			{S: "keyName1"},
		},
	}
	m := MAP{}
	m.FromTape(&tape)
	if m["key"] != "keyName1" {
		t.Errorf("REMOVE Failed %s", m["key"])
	}
}

func TestNewFromTape(t *testing.T) {
	tape := bob.Tape{
		Cell: []bob.Cell{
			{S: Prefix},
			{S: Set},
			{S: "app"},
			{S: "myapp"},
		},
	}

	tx, err := NewFromTape(&tape)
	if err != nil {
		t.Errorf("Failed to create new magic from tape")
	}

	if tx.getValue("app") != "myApp" {
		t.Errorf("Unexpected output %+v %s", tx, tx.getValue("app"))
	}
}
