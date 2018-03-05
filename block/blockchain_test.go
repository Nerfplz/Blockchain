package block

import (
	"testing"
)

func TestBlockchainCreation(t *testing.T) {
	blockchain := NewBlockchain(testData)
	checkBlock(t, blockchain.Last(), testData, Hash{})
}

func TestBlockchainIsValid(t *testing.T) {
	blockchain := NewBlockchain(testData)
	blockchain.AddBlock(testData)
	blockchain.AddBlock(testData)
	if !blockchain.IsValid() {
		t.Error("The blockchain was valid, but IsValid failed.")
	}
}

func TestBlockchainCorruptedIsValid(t *testing.T) {
	blockchain := NewBlockchain(testData)
	blockchain.AddBlock(testData)
	block := NewGenesis(testData)
	blockchain.blocks[block.Hash] = block
	if blockchain.IsValid() {
		t.Error("Corrupted blockchain appeared valid.")
	}
}
