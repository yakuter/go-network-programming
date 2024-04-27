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
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	defer ln.Close()

	fmt.Println("Server is running on port 8080")
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	buf := make([]byte, 14)
	_, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
		return
	}

	decData := &Data{}
	decData.ID = binary.BigEndian.Uint32(buf[:4])
	decData.Timestamp = binary.BigEndian.Uint64(buf[4:12])
	decData.Value = int16(binary.BigEndian.Uint16(buf[12:]))
	fmt.Printf("Received - ID: %d, Timestamp: %d, Value: %d\n",
		decData.ID, decData.Timestamp, decData.Value)
}
