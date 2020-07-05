package mapp

import (
	"testing"

	"github.com/rohenaz/go-bob"
)

func TestFromTape(t *testing.T) {
	tape := bob.Tape{
		Cell: []bob.Cell{
			{S: "testing"},
		},
	}
	m := &MAP{}
	m.FromTape(tape)
	if m == nil {
		t.Error("MAP was nil")
	}
}

func TestAdd(t *testing.T) {
	tape := bob.Tape{
		Cell: []bob.Cell{
			{S: Prefix},
			{S: ADD},
			{S: "keyName"},
			{S: "something"},
			{S: "something else"},
		},
	}
	m := MAP{}
	m.FromTape(tape)
	// expectedResult := []string{"something", "something else"}
	if m["keyName"] == nil {
		t.Errorf("ADD Failed %s", m["keyName1"])
	}
}

func TestSet(t *testing.T) {
	tape := bob.Tape{
		Cell: []bob.Cell{
			{S: Prefix},
			{S: SET},
			{S: "keyName1"},
			{S: "something"},
			{S: "keyName2"},
			{S: "something else"},
		},
	}
	m := MAP{}
	m.FromTape(tape)
	if m["keyName1"] != "something" {
		t.Errorf("SET Failed %s", m["keyName1"])
	}
}
