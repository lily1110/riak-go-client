package riak

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
)

func rpbEnsureCode(expected byte, actual byte) (err error) {
	if expected != actual {
		err = errors.New(fmt.Sprintf("expected response code %d, got: %d", expected, actual))
	}
	return
}

func rpbWrite(code byte, data []byte) []byte {
	buf := new(bytes.Buffer)
	// write total message length, including one byte for msg code
	binary.Write(buf, binary.BigEndian, int32(len(data)+1))
	// write the message code
	binary.Write(buf, binary.BigEndian, int8(code))
	// write the protobuf data
	buf.Write(data)
	return buf.Bytes()
}
