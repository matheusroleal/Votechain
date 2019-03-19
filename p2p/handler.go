package p2p

import (
	"context"
	"log"
	"sync"
	"encoding/json"

	"github.com/davecgh/go-spew/spew"
	types "github.com/matheusroleal/Votechain/types"
	blockchain "github.com/matheusroleal/Votechain/blockchain"
)

type Transaction struct {
	Vote  string `json:"vote"`
	Key   string `json:"key"`
}

var mutex = &sync.Mutex{}

func (node *Node) GetNewAddress() *types.GetNewAddressResponse {
	var res types.GetNewAddressResponse
	addr := node.wallet.GetNewAddress()
	res.Address = addr
	return &res
}

func (node *Node) BroadcastBlock(t *Transaction) *types.SendTxResponse{
	var res types.SendTxResponse

	last_index_block := len(node.blockchain.Chain) - 1
	newBlock,e := blockchain.GenerateBlock(node.blockchain.Chain[last_index_block], t.Vote, t.Key)

	mutex.Lock()
	if e != nil {
		log.Printf("ERROR: Could not generate block")
	} else {
		data, err := json.Marshal(newBlock)
		if err != nil{
			log.Printf("ERROR: Could not send block")
		}
		node.pubsub.Publish("blocks", data)
	}
	mutex.Unlock()
	res.Txid = newBlock.Hash
	return &res
}

func (node *Node) ListenBlocks(ctx context.Context) {
	sub, err := node.pubsub.Subscribe("blocks")
	if err != nil {
		panic(err)
	}
	go func() {
		for {
			msg, err := sub.Next(ctx)
			if err != nil {
				panic(err)
			}

			block := blockchain.Block{}

			err = json.Unmarshal(msg.GetData(), &block)
			if err != nil {
				return
			}

			mutex.Lock()
			blockchain.CreateChain(node.blockchain,block)
			mutex.Unlock()

			spew.Dump(node.blockchain.Chain)
		}
	}()
}
