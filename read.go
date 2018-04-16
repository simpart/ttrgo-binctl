package binctl

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"os"
)

type Target struct {
	path string
	file *os.File
}

func NewReader(path string) (*Target, error) {
	ret := new(Target)
	// open file
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	// set member
	ret.file = file
	ret.path = path
	// return instance
	return ret, nil
}

func (tgt Target) Read(data interface{}, siz int) error {

	// read binary
	read_dat := make([]byte, siz)
	no, err := tgt.file.Read(read_dat)
	if 0 == no {
		fmt.Println("end of file")
		return nil
	}
	if err != nil {
		return err
	}

	// set header fields
	err = binary.Read(
		bytes.NewBuffer(read_dat),
		binary.LittleEndian,
		data,
	)
	if err != nil {
		return err
	}
	return nil
}
