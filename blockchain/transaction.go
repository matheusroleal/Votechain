package blockchain

import (
  "fmt"
  "time"
  "sync"
  "github.com/davecgh/go-spew/spew"
)

type Transaction struct {
	Data  string `json:"data"`
	Token int    `json:"token"`
}

var Blockchain []Block
var mutex = &sync.Mutex{}

func replaceChain(newBlock []Block) {
	if len(Blockchain) < len(newBlock) {
		Blockchain = newBlock
	}
}

func CreateChain(t *Transaction){
  last_index_block := len(Blockchain) - 1

  mutex.Lock()
  newBlock,e := generateBlock(Blockchain[last_index_block], t.Data)
  mutex.Unlock()

  if e != nil {
    fmt.Printf("ERROR: Could not generate block")
  }

  if checkBlockValidation(newBlock, Blockchain[last_index_block]) {
    Blockchain = append(Blockchain, newBlock)
    spew.Dump(newBlock)
  }
}

func GenesisBlock() {
	genesisBlock := Block{0, time.Now().String(), "", "", "", difficulty, ""}
  spew.Dump(genesisBlock)

  mutex.Lock()
	Blockchain = append(Blockchain, genesisBlock)
  mutex.Unlock()
}
