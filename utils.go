package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"reflect"
)

// IntToHex converts an int64 to a byte array
func IntToHex(num int64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)
	}

	return buff.Bytes()
}

func ArrayOfBytes(i int, b byte) (p []byte) {

	for i != 0 {

		p = append(p, b)
		i--
	}
	return
}

func SHAString(data []byte) string {
	return fmt.Sprintf("%x", data)
}

func IsNil(v interface{}) bool {
	return reflect.ValueOf(v).IsNil()
}
