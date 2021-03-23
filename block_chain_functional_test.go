package go_blockchain

import (
	"github.com/stretchr/testify/assert"

	"testing"
)

func TestAddDataShouldIterateToTheNextBlockIndex(t *testing.T) {
	blockChain := NewBlockChain(UUIDHashGenerator{})

	// the index is 0 before calling AddData
	assert.Equal(t, 0, blockChain.BlockIndex)

	blockChain.AddData(Data{})

	// the index is 1 after calling AddData
	assert.Equal(t, 1, blockChain.BlockIndex)
}

func TestThatPreviousBlockHashIsEmptyForFirstBlock(t *testing.T) {
	blockChain := NewBlockChain(UUIDHashGenerator{})

	blockChain.AddData(Data{})

	block, _ := blockChain.GetBlock(0)

	assert.Equal(t, "", block.PrevHash)
}

func TestThatNextBlockHashIsEqualToThePreviousBlockCurrentHash(t *testing.T) {
	blockChain := NewBlockChain(UUIDHashGenerator{})

	blockChain.AddData(Data{})
	blockChain.AddData(Data{})

	previousBlock, _ := blockChain.GetBlock(0)
	nextBlock, _ := blockChain.GetBlock(1)

	// we assert that the second block hash is equal to the previous block hash
	assert.Equal(t, previousBlock.CurrentHash, nextBlock.PrevHash)
}

func TestCountBlocks(t *testing.T) {
	blockChain := NewBlockChain(UUIDHashGenerator{})

	// assert that we have 0 block
	assert.Equal(t, 0, blockChain.CountBlocks())

	blockChain.AddData(Data{})

	// assert that we have 1 block
	assert.Equal(t, 1, blockChain.CountBlocks())
}

func TestGetBlockReturnTheAssociatedBlockWhenIndexExist(t *testing.T) {
	blockChain := NewBlockChain(UUIDHashGenerator{})
	data := Data{}

	blockChain.AddData(Data{})
	block, err := blockChain.GetBlock(0)

	assert.Nil(t, err)
	assert.Equal(t, data, block.Data)
}

func TestGetBlockThrowAnExceptionWhenIndexDoesNotExist(t *testing.T) {
	blockChain := NewBlockChain(UUIDHashGenerator{})
	blockChain.AddData(Data{})

	block, err := blockChain.GetBlock(1)

	assert.Nil(t, block)
	assert.EqualErrorf(t, err, "the block with index 1 does not exist in the blockchain.", "the error is not as expected")
}
