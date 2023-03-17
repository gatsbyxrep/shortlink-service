package CRC32Crypter

import (
	"fmt"
	"hash/crc32"
)

type CRC32Crypter struct{}

// Its not too secured but short
func (crypter CRC32Crypter) GetHashSumFromString(str string) []byte {
	data := []byte(str)
	return []byte(fmt.Sprintf("%08x", crc32.ChecksumIEEE(data)))
}
