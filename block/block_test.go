package block

import (
	"testing"
	"time"

	"github.com/Nerfplz/blockchain/data"
)

var (
	testData = data.String("Test Data")
)

func checkBlock(t *testing.T, block *Block, d data.Data, prevHash Hash) {
	if block.Data != d {
		t.Error("Data got changed during creation.")
	}
	if block.Timestamp == time.Unix(0, 0).Unix() {
		t.Error("Zero timestamp has been initialized")
	}
	if block.PreviousHash != prevHash {
		t.Error("Previous hash has not been correctley accessed.")
	}
}

func TestNewBlock(t *testing.T) {
	block := newBlock(testData, Hash{})
	checkBlock(t, block, testData, Hash{})
}

func TestNew(t *testing.T) {
	block := New(testData, &Block{})
	checkBlock(t, block, testData, Hash{})
}
