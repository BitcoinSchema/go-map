package main

import (
	"log"

	"github.com/bitcoinschema/go-bpu"
	magic "github.com/bitcoinschema/go-map"
)

func main() {
	appKey := "app"
	appVal := "myapp"
	tape := bpu.Tape{
		Cell: []bpu.Cell{
			{S: &magic.Prefix},
			{S: &magic.Set},
			{S: &appKey},
			{S: &appVal},
		},
	}

	tx, err := magic.NewFromTape(&tape)
	if err != nil {
		log.Fatalf("failed to create new MAP from tape")
	}

	log.Printf("cmd: [%s] key: [%s] value: [%s]", tx[magic.Cmd], "app", tx["app"])
}
