package binctl

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"os"
)

func NewWriter(path string) (*Target, error) {
	ret := new(Target)
	// open file
	file, err := os.Create(path)
	if err != nil {
		return nil, err
	}
	// set member
	ret.file = file
	ret.path = path
	// return instance
	return ret, nil
}

func (tgt Target) Write(data interface{}) error {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.LittleEndian, data)

	if err != nil {
		fmt.Println("err:", err)
		return err
	}
	_, err = tgt.file.Write(buf.Bytes())
	return err
}
