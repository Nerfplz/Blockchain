package block

import (
	"fmt"

	"github.com/Nerfplz/blockchain/data"
)

// Blockchain is a glorified slice of blocks.
type Blockchain struct {
	blocks      map[Hash]*Block
	last        **Block
	genesisHash Hash
}

// NewBlockchain creates a blockchain that is initialized with a genesis Block.
func NewBlockchain(d data.Data) Blockchain {
	genesis := NewGenesis(d)
	blockchain := Blockchain{make(map[Hash]*Block), &genesis, genesis.Hash}
	blockchain.blocks[genesis.Hash] = genesis
	return blockchain
}

//Last returns a pointer to the last Block in the Blockchain.
func (b Blockchain) Last() *Block {
	return *b.last
}

// ByHash returns the block with the given hash
func (b Blockchain) ByHash(hash Hash) *Block {
	return b.blocks[hash]
}

// AddBlock adds the next block to the Blockchain.
func (b *Blockchain) AddBlock(d data.Data) {
	block := New(d, b.Last())
	b.blocks[block.Hash] = block
	b.last = &block
}

// IsValid checks if the Blockhain is valid, i.e. the hashes are lining up.
func (b Blockchain) IsValid() bool {
	block := b.Last()
	for block != nil {
		previous := b.ByHash(block.PreviousHash)
		if (previous == nil || block.PreviousHash != previous.Hash) && block.Hash != b.genesisHash {
			return false
		}
		block = previous
	}
	return true
}

func (b Blockchain) String() string {
	str := "Timestamp\t\t\tHash\t\t\t\t\t\t\t\t\tPrevious Hash\n"
	for _, bl := range b.blocks {
		str = str + fmt.Sprintf(bl.String())
	}
	return str
}
