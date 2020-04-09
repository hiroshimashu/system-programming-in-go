package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

func endien(){
	var i int32
	data := []byte{0x0, 0x0, 0x27, 0x10}
	binary.Read(bytes.NewReader(data), binary.BigEndian, &i)
	fmt.Printf("data: %d\n", i)
}