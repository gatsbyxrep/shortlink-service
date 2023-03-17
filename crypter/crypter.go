package crypter

type ICrypter interface {
	GetHashSumFromString(str string) []byte
}
