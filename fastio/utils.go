package fastio

import (
	"bytes"

	"github.com/lukechampine/fastxor"
)

func concatBytes(s [][]byte) []byte {
	return bytes.Join(s, nil)
}

func buildEPart1(flag byte, id []byte) []byte {
	return concatBytes([][]byte{{flag}, id})
}

func parseEPart1(p1 []byte) (byte, []byte) {
	return p1[0], p1[1:]
}

func xor32Bytes(part1, part2 []byte) []byte {
	e := make([]byte, eSize)
	fastxor.Block(e[0:16], part1[0:16], part2[0:16])
	fastxor.Block(e[16:32], part1[16:32], part2[16:32])

	return e
}

const paddedIdSize = MaxIDSize + 1

func padId(id []byte) []byte {
	idLen := len(id)

	padded := make([]byte, 0, paddedIdSize)
	padded = append(padded, byte(idLen))
	padded = append(padded, id...)
	padded = append(padded, make([]byte, paddedIdSize-idLen-1)...)

	return padded
}

func unpadId(id []byte) []byte {
	return id[1 : id[0]+1]
}
