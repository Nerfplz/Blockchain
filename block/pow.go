package block

import (
	"encoding/binary"
)

func i32ToBytes(i int32) [4]byte {
	var bs [4]byte
	binary.LittleEndian.PutUint32(bs[:], uint32(i))
	return bs
}

func calculateHashDifficulty(hash Hash) uint32 {
	return uint32(hash[0]) +
		uint32(hash[1])<<8*1 +
		uint32(hash[2])<<8*2 +
		uint32(hash[3])<<8*3
}

// ConcatNonce concatinates the nonce with the data.
func ConcatNonce(nonce int32, bytes []byte) []byte {
	byteSlice := i32ToBytes(nonce)
	return append(bytes, byteSlice[:]...)
}

// IsProofValid validates a proof of work.
func IsProofValid(nonce int32, bytes []byte, difficulty uint32) bool {
	hash := HashFunction(ConcatNonce(nonce, bytes))
	hashNumber := calculateHashDifficulty(hash)
	if hashNumber < difficulty {
		return true
	}
	return false
}
