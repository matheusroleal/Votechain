package blockchain

import (
  "sync"
  "github.com/davecgh/go-spew/spew"
)

var mutex = &sync.Mutex{}

func CreateChain(blockchain *Blockchain,newBlock Block){
  last_index_block := len(blockchain.Chain) - 1
  if checkBlockValidation(newBlock, blockchain.Chain[last_index_block]) {
    mutex.Lock()
    blockchain.Chain = append(blockchain.Chain, newBlock)
    mutex.Unlock()
    spew.Dump(newBlock)
  }
}
