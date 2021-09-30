package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBlockChainCollection_AddBlockChain(t *testing.T) {
	c := BlockChainCollection{}
	b := NewBlockChain(UUIDHashGenerator{})
	b2 := NewBlockChain(UUIDHashGenerator{})
	c.AddBlockChain(b)
	c.AddBlockChain(b2)

	assert.Equal(t, []BlockChain{b, b2}, c.BlockChains)
}