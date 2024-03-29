package magic

import (
	"testing"

	"github.com/bitcoinschema/go-bob"
	"github.com/bitcoinschema/go-bpu"
)

var mapKey = "key"
var mapValue = "value"
var keyNameStr = "keyName"

var mapTestKey = "keyName1"
var mapTestValue = "something"

func contains(arr []string, str string) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}
	return false
}

func TestSelectDelete(t *testing.T) {
	txVal := "a9a4387d2baa2edcc53ec040b3affbc38778e9dd876f9a47e5c767c785aacf76"
	tape := &bpu.Tape{
		Cell: []bpu.Cell{
			{S: &Prefix},
			{S: &Select},
			{S: &txVal},
			{S: &Delete},
			{S: &mapTestKey},
			{S: &mapTestValue},
		},
	}

	m, err := NewFromTape(tape)
	if err != nil {
		t.Fatalf("Failed to create magicTx from tape %s", err)
	}

	if m[Cmd] != Select || m[mapValue] != mapTestValue {
		t.Fatalf("SELECT + DELETE Failed. command: %s, full %+v, key: %s, value: %s", m[Cmd], m, m[mapTestKey], mapTestValue)
	}
}

func TestAdd(t *testing.T) {
	somethingElse := "something else"
	tape := bpu.Tape{
		Cell: []bpu.Cell{
			{S: &Prefix},
			{S: &Add},
			{S: &keyNameStr},
			{S: &mapTestValue},
			{S: &somethingElse},
		},
	}
	m, err := NewFromTape(&tape)
	if err != nil {
		t.Fatalf("error occurred: %s", err.Error())
	}

	switch m[keyNameStr].(type) {
	case []string:
		if !contains(m[keyNameStr].([]string), mapTestValue) ||
			!contains(m[keyNameStr].([]string), "something else") {
			t.Fatalf("ADD Failed %s", m["keyName1"])
		}
	default:
		t.Fatalf("ADD Failed %s", m[mapTestKey])
	}
}

func TestGetValue(t *testing.T) {

	tape := bpu.Tape{
		Cell: []bpu.Cell{
			{S: &Prefix},
			{S: &Add},
			{S: &keyNameStr},
			{S: &mapTestValue},
		},
	}
	m, err := NewFromTape(&tape)
	if err != nil {
		t.Fatalf("error occurred: %s", err.Error())
	}

	if val := m.getValue(keyNameStr); val != "something" {
		t.Fatalf("expected: [%v] but got: [%v]", "something", val)
	}
}

func TestGetValues(t *testing.T) {

	anotherValue := "another value"
	thirdValue := "third value"
	tape := bpu.Tape{
		Cell: []bpu.Cell{
			{S: &Prefix},
			{S: &Add},
			{S: &keyNameStr},
			{S: &mapTestValue},
			{S: &anotherValue},
			{S: &thirdValue},
		},
	}
	m, err := NewFromTape(&tape)
	if err != nil {
		t.Fatalf("error occurred: %s", err.Error())
	}
	if val := m.getValues(keyNameStr); len(val) > 0 && val[0] != "something" {
		t.Fatalf("expected: [%v] but got: [%v]", "something", val)
	} else if val[1] != "another value" {
		t.Fatalf("expected: [%v] but got: [%v]", "another value", val)
	}
}

func TestDelete(t *testing.T) {

	tape := bpu.Tape{
		Cell: []bpu.Cell{
			{S: &Prefix},
			{S: &Delete},
			{S: &keyNameStr},
			{S: &mapTestValue},
		},
	}
	m, err := NewFromTape(&tape)
	if err != nil {
		t.Fatalf("error occurred: %s", err.Error())
	}

	if m[mapValue] != mapTestValue {
		t.Errorf("DELETE Failed %+v %s", m, mapTestValue)
	}

}

func TestSet(t *testing.T) {
	keyName2 := "keyName2"
	somethingElse := "something else"
	tape := bpu.Tape{
		Cell: []bpu.Cell{
			{S: &Prefix},
			{S: &Set},
			{S: &mapTestKey},
			{S: &mapTestValue},
			{S: &keyName2},
			{S: &somethingElse},
		},
	}
	m, err := NewFromTape(&tape)
	if err != nil {
		t.Fatalf("error occurred: %s", err.Error())
	}
	if m[mapTestKey] != mapTestValue {
		t.Errorf("SET Failed %s", m[mapTestKey])
	}
}

func TestRemove(t *testing.T) {
	keyName1 := "keyName1"
	tape := bpu.Tape{
		Cell: []bpu.Cell{
			{S: &Prefix},
			{S: &Remove},
			{S: &keyName1},
		},
	}
	m, err := NewFromTape(&tape)
	if err != nil {
		t.Fatalf("error occurred: %s", err.Error())
	}
	if m[mapKey] != keyName1 {
		t.Errorf("REMOVE Failed %s", m[mapKey])
	}
}

func TestInvalid(t *testing.T) {
	t.Run("invalid tx that used to crash the parser", func(t *testing.T) {
		// tx b947f1d8a88d171dff70f4d6845cee04c05f5cedd1d9f04030eead164a7d4846
		tx, err := bob.NewFromRawTxString("0100000001bcbe8916daf12b8474cf3c2a8027be748b87f819cf54ee2eeb4a6703ba254c18010000006a473044022075af94ad6076d020db32d29a2368031d366a519b6d7880ff7efe231c7ec5a2c902202f90f0e8cc1ea149b82b297e628a7820587e90895a6119fd80b4b4bd0398f14041210346cdb7f53c7656b9e4f76cf359091498810c8ff9864f49b63c10ed165e5e9dcbffffffff0200000000000000009a006a2231394878696756345179427633744870515663554551797131707a5a56646f41757424e381aae38293e381a7e38282e38184e38184efbc9f44454c455445e381a0e3818be382890a746578742f706c61696e055554462d38106d61705f746573743030312e68746d6c017c223150755161374b36324d694b43747373534c4b79316b683536575755374d745552350644454c4554450134d9230000000000001976a914a806d9a4b98d5893fdc7ae12bdf41e75174696c288ac00000000")
		if err != nil {
			t.Fatalf("error occurred: %s", err.Error())
		}

		_, err = NewFromTape(&tx.Out[0].Tape[2])
		if err != nil {
			t.Fatalf("error occurred: %s", err.Error())
		}
	})
}
