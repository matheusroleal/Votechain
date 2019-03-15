package blockchain

import (
  "fmt"
  "time"
  "sync"
  "github.com/davecgh/go-spew/spew"
)

type Transaction struct {
	Vote  string `json:"vote"`
	Key string    `json:"key"`
}

var Chain []Block
var mutex = &sync.Mutex{}

func ReplaceChain(newBlock []Block) {
	if len(Chain) < len(newBlock) {
		Chain = newBlock
	}
}

func CreateChain(t *Transaction){
  last_index_block := len(Chain) - 1

  mutex.Lock()
  newBlock,e := generateBlock(Chain[last_index_block], t.Vote, t.Key)
  mutex.Unlock()

  if e != nil {
    fmt.Printf("ERROR: Could not generate block")
  } else {
    if checkBlockValidation(newBlock, Chain[last_index_block]) {
      mutex.Lock()
      Chain = append(Chain, newBlock)
      mutex.Unlock()
      spew.Dump(newBlock)
    }
  }
}

func GenesisBlock() {
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

  mutex.Lock()
	Chain = append(Chain, genesisBlock)
  mutex.Unlock()
}
