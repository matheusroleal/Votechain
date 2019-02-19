package blockchain

import (
	"crypto/sha256"
	"time"
	"encoding/hex"
)

type Block struct {
	Index      int
	Timestamp  string
	Vote       string
	Hash       string
	PrevHash   string
	Difficulty int
	Nonce      string
}


func calculateHash(block Block) string {
	record := string(block.Index) + block.Timestamp + block.Vote + block.PrevHash + block.Nonce
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

func generateBlock(oldBlock Block, data string) (Block, error) {
  var newBlock Block

  newBlock.Index = oldBlock.Index + 1
  newBlock.Timestamp = time.Now().String()
  newBlock.PrevHash = oldBlock.Hash
  newBlock.Vote = data
  newBlock.PrevHash = oldBlock.Hash
  newBlock.Hash = calculateHash(newBlock)

  return newBlock, nil
}

func checkBlockValidation(newBlock, oldBlock Block) bool {
	if oldBlock.Index+1 != newBlock.Index {
	  return false
	}

	if oldBlock.Hash != newBlock.PrevHash {
	  return false
	}

	if calculateHash(newBlock) != newBlock.Hash {
	  return false
	}

	return true
}
