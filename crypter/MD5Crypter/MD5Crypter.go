package MD5Crypter

import (
	"crypto/md5"
	"encoding/hex"
)

type MD5Crypter struct{}

func (crypter MD5Crypter) GetHashSumFromString(str string) []byte {
	hash := md5.Sum([]byte(str))
	result := make([]byte, hex.EncodedLen(len(hash[:])))
	hex.Encode(result, hash[:])
	return result
}
