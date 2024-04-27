package main

import (
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
		ID:        100,        // 1 byte
		Timestamp: 1600000000, // 5 bytes
		Value:     10,         // 1 byte
	}

	buf := make([]byte, binary.MaxVarintLen64)
	n := binary.PutUvarint(buf, uint64(data.ID))
	n += binary.PutUvarint(buf[n:], data.Timestamp)
	n += binary.PutVarint(buf[n:], int64(data.Value))

	fmt.Printf("encoded binary length: %d\n", len(buf[:n]))
	// encoded binary length: 7

	id, idLen := binary.Uvarint(buf)
	ts, tsLen := binary.Uvarint(buf[idLen:])
	value, _ := binary.Varint(buf[idLen+tsLen:])

	decodedData := &Data{
		ID:        uint32(id),
		Timestamp: uint64(ts),
		Value:     int16(value),
	}
	fmt.Printf("ID: %d, Timestamp: %d, Value: %d\n", decodedData.ID, decodedData.Timestamp, decodedData.Value)
	// ID: 100, Timestamp: 1600000000, Value: 10

	// 300 -> in binary 100101100 -> 10010110 10000000 -> 0x96 0x80
	// 10101100 00000010
	// ↑        ↑
	// msb is 1 msb is 0 so it quits
	// 10101100 00000010
	//  0101100  0000010
	//  0000010  0101100
	//  -> 100101100
	//  -> 300
}
