package block

import "testing"

var (
	emptyTestHash = Hash{}
)

func TestCalculateHashDifficulty(t *testing.T) {
	if calculateHashDifficulty(emptyTestHash) != 0 {
		t.Error("calculateHash miscalculated empty hash value.")
	}
	testHash1 := Hash{}
	testHash1[0] = 255
	if calculateHashDifficulty(testHash1) != 255 {
		t.Error("calculateHash miscalculated hash value in the first byte.")
	}

	testHash2 := Hash{}
	testHash2[1] = 255
	if calculateHashDifficulty(testHash2) != 65280 {
		t.Error("calculateHash miscalculated hash value in the second byte.")
	}

	testHash3 := Hash{}
	testHash3[0] = 42
	testHash3[1] = 255
	if calculateHashDifficulty(testHash3) != 65280+42 {
		t.Error("calculateHash miscalculated hash value in the first and second byte.")
	}
}
