package blockchain

import (
  "time"
  "sync"
  "github.com/davecgh/go-spew/spew"
)

var mutex = &sync.Mutex{}

type Blockchain struct {
	Chain []Block
}

func NewBlockchain() *Blockchain {
  var blockchain Blockchain

  genesisBlock := GenesisBlock()
  blockchain.Chain = append(blockchain.Chain, genesisBlock)

  return &blockchain
}

func (blockchain *Blockchain) CreateChain(newBlock Block) {
  last_index_block := len(blockchain.Chain) - 1
  if checkBlockValidation(newBlock, blockchain.Chain[last_index_block]) {
    mutex.Lock()
    blockchain.Chain = append(blockchain.Chain, newBlock)
    mutex.Unlock()
  }
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
