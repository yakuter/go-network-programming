package main

import (
	"bytes"
	"encoding/binary"
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

	buf := &bytes.Buffer{}
	_ = binary.Write(buf, binary.BigEndian, data)
	fmt.Printf("encoded binary length: %d\n", buf.Len())
	fmt.Printf("encoded binary: %v\n", buf.Bytes())
	// encoded binary length: 14
	// encoded binary: [0 0 0 100 0 0 0 0 95 94 16 0 0 10]

	dst := &Data{}
	_ = binary.Read(buf, binary.BigEndian, dst)
	fmt.Printf("ID: %d, Timestamp: %d, Value: %d\n", dst.ID, dst.Timestamp, dst.Value)
	// ID: 100, Timestamp: 1600000000, Value: 10
}
