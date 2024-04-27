package main

import (
	"encoding/binary"
	"fmt"
	"net"
)

type Data struct {
	ID        uint32
	Timestamp uint64
	Value     int16
}

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	data := &Data{
		ID: 105, Timestamp: 1600000000, Value: 12,
	}
	fmt.Println("Data to send - ID:", data.ID, ", Timestamp:", data.Timestamp, ", Value:", data.Value)

	buf := make([]byte, 14)
	binary.BigEndian.PutUint32(buf[0:], data.ID)
	binary.BigEndian.PutUint64(buf[4:], data.Timestamp)
	binary.BigEndian.PutUint16(buf[12:], uint16(data.Value))

	_, err = conn.Write(buf)
	if err != nil {
		fmt.Println("Error writing to server:", err.Error())
	} else {
		fmt.Println("Sent data to server")
	}
}
