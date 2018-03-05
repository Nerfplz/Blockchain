package block

import (
	"blockchain/data"
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
	"time"
)

// HashSize is the size of a hash in bytes.
const HashSize = sha256.Size

// HashFunction is an alias for the hash function that is used for the calculation of the blocks.
var HashFunction = sha256.Sum256

// Hash is a slice of bytes that a size of HashSize.
type Hash = [HashSize]byte

// Block is the smallest unit in the blockchain.
type Block struct {
	Timestamp    int64
	Data         data.Data
	PreviousHash Hash
	Hash         Hash
}

func newBlock(data data.Data, PreviousHash Hash) *Block {
	block := &Block{
		Timestamp:    time.Now().Unix(),
		Data:         data,
		PreviousHash: PreviousHash,
	}
	block.generateHash()
	return block
}

// New return an apropriatley construcet block.
func New(d data.Data, previousBlock *Block) *Block {
	return newBlock(d, previousBlock.Hash)
}

// NewGenesis returns a block that can be used as the first block of the blockchain.
func NewGenesis(d data.Data) *Block {
	return newBlock(d, Hash{})
}

// Only called by the constructor function of Block
func (b *Block) generateHash() {
	bytes := bytes.Join([][]byte{[]byte(strconv.FormatInt(b.Timestamp, 10)), b.Data.Bytes(), b.PreviousHash[:]}, []byte{})
	b.Hash = HashFunction(bytes)
}

func (b Block) String() string {
	return fmt.Sprintf("%s\t%s\t%s\n", time.Unix(b.Timestamp, 0).String(), toHex(b.Hash), toHex(b.PreviousHash))
}

func toHex(hash Hash) string {
	str := hex.EncodeToString(hash[:])
	return str
}
