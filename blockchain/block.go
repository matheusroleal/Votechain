package blockchain

import (
	"crypto/sha256"
	"time"
	"encoding/hex"
)

type Block struct {
	Index     int
	Timestamp string
	Vote       string
	Hash      string
	PrevHash  string
}

func calculateHash(block Block) string {
  record := string(block.Index) + block.Timestamp + string(block.Vote) + block.PrevHash

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
  newBlock.Hash = calculateHash(oldBlock)

  return newBlock, nil
}
