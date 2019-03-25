package blockchain

import (
	"crypto/sha256"
	"time"
	"encoding/hex"
	"fmt"
	"strings"

	wallet "github.com/matheusroleal/Votechain/wallet"
)

type Block struct {
	Index      int
	Timestamp  string
	Vote       string
	Hash       string
	PrevHash   string
	Difficulty int
	Nonce      string
	Address		 string
}

const difficulty = 1

func calculateHash(block Block) string {
	record := string(block.Index) + block.Timestamp + block.Vote + block.PrevHash + block.Nonce
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

func GenerateBlock(oldBlock Block, data string, key string) (Block, error) {
  var newBlock Block

  newBlock.Index = oldBlock.Index + 1
  newBlock.Timestamp = time.Now().String()
	newBlock.Vote = data
	newBlock.Address = wallet.GetPlbAddress(key)
  newBlock.PrevHash = oldBlock.Hash
	newBlock.Difficulty = difficulty

	for i := 0; ; i++ {
		hex := fmt.Sprintf("%x", i)
		newBlock.Nonce = hex
		if !checkHashValidation(calculateHash(newBlock), newBlock.Difficulty) {
			fmt.Println(calculateHash(newBlock))
			time.Sleep(time.Second)
			continue
		} else {
			fmt.Println(calculateHash(newBlock), " work done!")
			newBlock.Hash = calculateHash(newBlock)
			break
		}
	}
  return newBlock, nil
}

func checkHashValidation(hash string, difficulty int) bool {
	prefix := strings.Repeat("0", difficulty)
	return strings.HasPrefix(hash, prefix)
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
