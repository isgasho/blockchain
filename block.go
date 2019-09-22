package blockchain

import (
	"bytes"
	"crypto/sha256"
	"time"
)

// Block represents a single block in a chain of blocks (blockchain) and holds
// the position in the chain (`Index`), the time the block was created (`Timestamp1)
// the user-defined data (`Data`), the previous block's hash (`PrevHash`) and the
// current block's own hash (`Hash`).
type Block struct {
	Index     int64
	Timestamp time.Time
	Data      []byte
	PrevHash  []byte
	Hash      []byte
}

// NewBlock creates and returns a new empty `Block` suitable for use as a
// "Genesis Block". All values except the creation time of the block (`Timestamp`)
// and the block's hash itself (`Hash`) are zero-value.
func NewBlock() Block {
	b := Block{}

	b.Timestamp = time.Now()
	b.Hash = hashBlock(b)

	return b
}

// Generate creates a new block from the current block which is assumed to be the
// last block in the chain. THe new block has an index equal to the last block
// incremented by one, as well as the hash of the previous block. Data can be
// any slice of bytes to store in the new block.
func (b Block) Generate(data []byte) Block {
	n := Block{
		Index:     b.Index + 1,
		Timestamp: time.Now(),
		Data:      data,
		PrevHash:  b.Hash,
	}

	n.Hash = hashBlock(n)

	return n
}

// Validate validates the current block (`b`) with a previous block in a chain
// (`o`) by ensuring that the new block (`b`) has an index one greater than the
// previous block in the chain (`o`), the hash of the new block is valid and
// the `PrevHash` field of the new block matches the hash of the previous block.
func (b Block) Validate(o Block) bool {
	if (bytes.Compare(b.Hash, hashBlock(b)) != 0) ||
		(b.Index != (o.Index + 1)) ||
		(bytes.Compare(b.PrevHash, o.Hash) != 0) {
		return false
	}
	return true
}

func hashBlock(b Block) []byte {
	h := sha256.New()

	h.Write(Int64Bytes(b.Index))
	h.Write(Int64Bytes(b.Timestamp.Unix()))
	h.Write(b.Data)
	h.Write(b.PrevHash)

	return h.Sum(nil)
}
