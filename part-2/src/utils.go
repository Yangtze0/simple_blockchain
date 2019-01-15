package main

import (
	"bytes"
	"encoding/binary"
)

// 将一个 int64 转化为一个字节数组(byte array)
func IntToHex(num int64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num)

	return buff.Bytes()
}
