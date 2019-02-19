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

  Blockchain = append(Blockchain, newBlock)

  fmt.Print("New Block Created: ")
  fmt.Print(newBlock)
}

func GenesisBlock() {
	genesisBlock := Block{0, time.Now().String(), "", "", ""}
	Blockchain = append(Blockchain, genesisBlock)
}
