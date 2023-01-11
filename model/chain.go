package model

type BlockChain struct {
	blocks []*Block
}

func NewBlockChain() *BlockChain {
	return &BlockChain{blocks: []*Block{NewGenesisBlock()}}
}

func (bc *BlockChain) AddBlock(header *Header) {
	prevHash := bc.blocks[len(bc.blocks)-1].Hash
	newBlock := NewBlock(header)
	newBlock.Header.PrevHash = prevHash
	bc.blocks = append(bc.blocks, newBlock)
}
