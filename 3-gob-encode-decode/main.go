package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

type Data struct {
	ID        uint32 // 4 bytes
	Timestamp uint64 // 8 bytes
	Value     int16  // 2 bytes
}

func main() {
	data := &Data{
		ID: 100, Timestamp: 1600000000, Value: 10,
	}

	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	dec := gob.NewDecoder(&buf)

	_ = enc.Encode(data)
	fmt.Printf("encoded binary length: %d\n", buf.Len())
	fmt.Printf("encoded binary: %v\n", buf.Bytes())
	// encoded binary length: 64
	// encoded binary: [49 255 129 3 1 1 4 68 ...

	var decData *Data
	_ = dec.Decode(&decData)
	fmt.Printf("ID: %d, Timestamp: %d, Value: %d\n", decData.ID, decData.Timestamp, decData.Value)
	// ID: 100, Timestamp: 1600000000, Value: 10
}
