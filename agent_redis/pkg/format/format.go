package format

import (
	"bytes"
	"bufio"
	"encoding/binary"
)

func MakeFormat(seq uint64, data []byte) ([]byte, error) {
	temp := &bytes.Buffer{}
	buf := bufio.NewWriter(temp)

	l := len(data)
	lErr := binary.Write(buf, binary.LittleEndian, uint32(l))
	if lErr != nil {
		return nil, lErr
	}
	buf.WriteByte(':')
	err := binary.Write(buf, binary.LittleEndian, seq)
	if err != nil {
		return nil ,err
	}
	buf.WriteByte(':')
	buf.WriteString("REDIS")
	buf.WriteByte(':')
	buf.Write(data)

	flushErr := buf.Flush()
	if flushErr != nil {
		return nil, flushErr
	}

	return temp.Bytes(), nil
} 