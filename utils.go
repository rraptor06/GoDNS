package utils

import (
	"bytes"
	"encoding/binary"
	"net"
)

func CreateDNSResponse(ip string) []byte {
	buffer := new(bytes.Buffer)

	binary.Write(buffer, binary.BigEndian, uint16(0x81a0)) // flags
	binary.Write(buffer, binary.BigEndian, uint16(1))      // questions
	binary.Write(buffer, binary.BigEndian, uint16(1))      // responses
	binary.Write(buffer, binary.BigEndian, uint16(0))
	binary.Write(buffer, binary.BigEndian, uint16(0))

	buffer.Write([]byte{0x07, 'e', 'x', 'a', 'm', 'p', 'l', 'e', 0x03, 'c', 'o', 'm', 0x00})
	binary.Write(buffer, binary.BigEndian, uint16(0x01)) // A type
	binary.Write(buffer, binary.BigEndian, uint16(0x01))

	buffer.Write([]byte{0x07, 'e', 'x', 'a', 'm', 'p', 'l', 'e', 0x03, 'c', 'o', 'm', 0x00})
	binary.Write(buffer, binary.BigEndian, uint16(0x01)) // A type
	binary.Write(buffer, binary.BigEndian, uint16(0x01)) // IN class
	binary.Write(buffer, binary.BigEndian, uint32(3600)) // TTL
	binary.Write(buffer, binary.BigEndian, uint16(4))
	net.IP([]byte(ip)).Write(buffer)

	return buffer.Bytes()
}
