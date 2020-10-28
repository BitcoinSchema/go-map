# go-map

A library for working with [Magic Attribute Protocol](https://github.com/rohenaz/MAP). Used in conjunction with a Tape from [go-bob](https://github.com/rohenaz/go-bob)

## Usage

```go
    import "github.com/rohenaz/go-bob"
    import "github.com/rohenaz/go-map"

    line := "<BOB formatted json string>"

    bobData := &bob.Tx{}
    if err := json.Unmarshal(line, bobData); err != nil {
      fmt.Println("Error:", err)
      return
    }

    for _, out := range bobData.Out {
      for _, tape := range out.Tape {
        mapData, err := magic.NewFromTape(tape)
        log.Printf("MAP TYPE is %s", mapData["type"])
      }
    }


```
