package mapp

import (
	"testing"

	"github.com/rohenaz/go-bob"
)

func contains(arr []string, str string) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}
	return false
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
			{S: DELETE},
			{S: "keyName"},
			{S: "something"},
		},
	}
	m := MAP{}
	m.FromTape(tape)

	switch m["keyName"].(type) {
	case []string:
		if !contains(m["keyName"].([]string), "something") {
			t.Errorf("DELETE Failed %s", m["keyName1"])
		}
		break
	default:
		t.Errorf("DELETE Failed %s", m["keyName1"])
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

func TestRemove(t *testing.T) {
	tape := bob.Tape{
		Cell: []bob.Cell{
			{S: Prefix},
			{S: REMOVE},
			{S: "keyName1"},
		},
	}
	m := MAP{}
	m.FromTape(tape)
	if m["key"] != "keyName1" {
		t.Errorf("REMOVE Failed %s", m["key"])
	}
}
