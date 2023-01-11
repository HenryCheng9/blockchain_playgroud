package model

import (
	"bytes"
	"crypto/sha256"
	"my_blockchain/common"
	"strconv"
)

type Block struct {
	Header    *Header
	Hash      []byte
	TxCounter int
	TxList    []*Tx
}

func (b *Block) SetHash() error {
	if b.Header == nil {
		return common.ErrHeaderEmpty
	}
	timestamp := []byte(strconv.FormatInt(b.Header.Timestamp, 10))
	headerBytes := bytes.Join([][]byte{b.Header.MerkleRoot, b.Header.PrevHash, timestamp}, []byte{})
	hash := sha256.Sum256(headerBytes)
	b.Hash = hash[:]
	return nil
}

func NewGenesisBlock() *Block {
	genesisBlock := &Block{
		Header:    &GenesisBlockHeader,
		TxCounter: 0,
		TxList:    make([]*Tx, 2048),
	}
	genesisBlock.SetHash()
	return genesisBlock
}

func NewBlock(header *Header) *Block {
	block := &Block{
		Header:    header,
		TxCounter: 0,
		TxList:    make([]*Tx, 2048),
	}
	block.SetHash()
	return block
}
