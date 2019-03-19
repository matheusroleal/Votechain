package blockchain

import (
  "time"
  "github.com/davecgh/go-spew/spew"
)

type Blockchain struct {
	Chain []Block
}

func NewBlockchain() *Blockchain {
  var blockchain Blockchain

  genesisBlock := GenesisBlock()
  blockchain.Chain = append(blockchain.Chain, genesisBlock)

  return &blockchain
}

func (blockchain *Blockchain) ReplaceChain(newBlock []Block) {
	if len(blockchain.Chain) < len(newBlock) {
		blockchain.Chain = newBlock
	}
}

func GenesisBlock() Block{
	var genesisBlock Block

  genesisBlock.Index = 0
  genesisBlock.Timestamp = time.Now().String()
  genesisBlock.Vote = ""
  genesisBlock.PrevHash = ""
  genesisBlock.Difficulty = difficulty
  genesisBlock.Nonce = ""
  genesisBlock.Address = ""
  genesisBlock.Hash = calculateHash(genesisBlock)

  spew.Dump(genesisBlock)

  return genesisBlock
}
