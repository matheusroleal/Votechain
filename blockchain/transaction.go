package blockchain

import (
  "fmt"
  "time"
)

type Transaction struct {
	Data  string `json:"data"`
	Token int    `json:"token"`
}

var Blockchain []Block
const difficulty = 1

func replaceChain(newBlock []Block) {
	if len(Blockchain) < len(newBlock) {
		Blockchain = newBlock
	}
}

func CreateChain(t *Transaction){
  last_index_block := len(Blockchain) - 1
  newBlock,e := generateBlock(Blockchain[last_index_block], t.Data)

  if e != nil {
    fmt.Printf("ERROR: Could not generate block")
  }

  if checkBlockValidation(newBlock, Blockchain[last_index_block]) {
    Blockchain = append(Blockchain, newBlock)
    fmt.Print(newBlock)
  }
}

func GenesisBlock() {
	genesisBlock := Block{0, time.Now().String(), "", "", "", difficulty, ""}
	Blockchain = append(Blockchain, genesisBlock)
}
