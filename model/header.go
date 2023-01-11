package model

import (
	"math/rand"
	"time"
)

type Hash []byte

const DefaultVersion = "1.0"

type Header struct {
	Version    string
	PrevHash   Hash
	MerkleRoot Hash
	Timestamp  int64
	Target     []byte
	Nonce      []byte
}

var GenesisBlockHeader = Header{
	Version:   DefaultVersion,
	PrevHash:  []byte("Genesis Header"),
	Timestamp: time.Now().Unix(),
	// todo: what's the other fields?
}

// NewHeader returns a new header of block.
// Input merkleRoot represents the data of Txs.
func NewHeader(merkleRoot Hash) *Header {
	return &Header{
		Version: DefaultVersion,
		MerkleRoot: merkleRoot,
		Timestamp: time.Now().Unix(),
		Nonce: []byte(rand.Int()),
	}
}
