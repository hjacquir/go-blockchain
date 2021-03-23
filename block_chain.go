package go_blockchain

import (
	"errors"
	"fmt"
)

type BlockChain struct {
	BlockIndex    int
	Blocks        []Block
	HashGenerator StringHashGenerator
	currentHash   string
}

func NewBlockChain(generator UUIDHashGenerator) BlockChain {
	blockChain := BlockChain{
		BlockIndex:    0,
		Blocks:        []Block{},
		HashGenerator: generator,
		currentHash:   generator.generate(),
	}
	// @todo SYNCHRONIZE THIS NEW BLOCKCHAIN WITH THE OTHERS
	// @todo get the other blockChain from the blockChainCollection
	// @todo add all blocks from other blockChain and add all blocks to the new blockchain
	// @todo lock the other blockChain to avoid adding data in it when operation is not finished.

	return blockChain
}

func (b *BlockChain) AddData(data Data) {
	// the previous block hash is
	// empty for the first block
	previousBlockHash := ""

	// when previous block exist we
	// get the previous block hash
	if b.BlockIndex != 0 {
		previousBlock := b.Blocks[b.BlockIndex-1]
		previousBlockHash = previousBlock.CurrentHash
	}

	block := Block{
		Data:        data,
		PrevHash:    previousBlockHash,
		CurrentHash: b.HashGenerator.generate(),
	}

	b.Blocks = append(b.Blocks, block)

	// @todo SYNCHRONIZE OTHERS BLOCKCHAIN WITH THIS
	// @todo add the current block in all other blockChain.
	// @todo lock the other blockChain to avoid adding data in it when operation is not finished.

	b.BlockIndex++
}

func (b BlockChain) CountBlocks() int {
	return len(b.Blocks)
}

func (b BlockChain) GetBlock(index int) (*Block, error) {
	if index > (b.CountBlocks() - 1) {
		return nil, errors.New(fmt.Sprintf("the block with index %v does not exist in the blockchain.", index))
	}

	block := b.Blocks[index]

	return &block, nil
}
