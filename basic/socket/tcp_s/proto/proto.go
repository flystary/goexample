package proto

import (
	"bufio"
	"bytes"
	"encoding/binary"
)

func Encode(message string) ([]byte, error) {
	var length = int32(len(message))
	var pkg = new(bytes.Buffer)

	err := binary.Write(pkg, binary.LittleEndian, length)
	if err != nil {
		return nil, err
	}

	err = binary.Write(pkg, binary.LittleEndian, []byte(message))
	if err != nil {
		return nil, err
	}
	return pkg.Bytes(), nil
}

func Decode(reader *bufio.Reader) (string, error) {
	lengthByte, _ := reader.Peek(4)
	lenghtBuff := bytes.NewBuffer(lengthByte)
	var length int32
	err := binary.Read(lenghtBuff, binary.LittleEndian, &length)
	if err != nil {
		return "", err
	}
	if int32(reader.Buffered()) < length+4 {
		return "", err
	}

	pack := make([]byte, int(4+length))
	_, err = reader.Read(pack)
	if err != nil {
		return "", err
	}
	return string(pack[4:]), nil
}
