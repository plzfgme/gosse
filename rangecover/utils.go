package rangecover

import (
	"bytes"
	"encoding/binary"
)

func makeKeyword(prefix uint64, suffixLen int) ([]byte, error) {
	buf := &bytes.Buffer{}
	err := binary.Write(buf, binary.BigEndian, prefix)
	if err != nil {
		return nil, err
	}
	err = binary.Write(buf, binary.BigEndian, uint8(suffixLen))
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
