package main

import (
	"encoding/binary"
	"fmt"
	"unsafe"
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

	fmt.Printf("Total size of struct: %d bytes\n", unsafe.Sizeof(data))
	fmt.Printf("Size of field ID: %d bytes\n", unsafe.Sizeof(data.ID))
	fmt.Printf("Size of field Timestamp: %d bytes\n", unsafe.Sizeof(data.Timestamp))
	fmt.Printf("Size of field Value: %d bytes\n", unsafe.Sizeof(data.Value))

	// ENCODING
	buf := make([]byte, 14)
	binary.BigEndian.PutUint32(buf[0:], data.ID)
	binary.BigEndian.PutUint64(buf[4:], data.Timestamp)
	binary.BigEndian.PutUint16(buf[12:], uint16(data.Value))
	fmt.Printf("encoded binary length: %d\n", len(buf))
	fmt.Printf("encoded binary: %v\n", buf)

	// DECODING
	decData := &Data{}
	decData.ID = binary.BigEndian.Uint32(buf[:4])
	decData.Timestamp = binary.BigEndian.Uint64(buf[4:12])
	decData.Value = int16(binary.BigEndian.Uint16(buf[12:]))
	fmt.Printf("ID: %d, Timestamp: %d, Value: %d\n", decData.ID, decData.Timestamp, decData.Value)
	// ID: 100, Timestamp: 1600000000, Value: 10

	// Total size of struct: 24 bytes
	// Size of field ID: 4 bytes
	// Size of field Timestamp: 8 bytes
	// Size of field Value: 2 bytes
	// encoded binary length: 14
	// encoded binary: [0 0 0 100 0 0 0 0 95 94 16 0 0 10]
}
