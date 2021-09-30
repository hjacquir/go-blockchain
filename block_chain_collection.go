package main

type BlockChainCollection struct {
	BlockChains []BlockChain
}

func (c *BlockChainCollection) AddBlockChain(chain BlockChain) {
	c.BlockChains = append(c.BlockChains, chain)
}
