package server

import (
	"fmt"
	"log"
  "encoding/json"
	"net/http"

	blockchain "github.com/matheusroleal/Votechain/blockchain"

	"github.com/dimfeld/httptreemux"
)

type CreateBlockHandler struct{}

func (g *CreateBlockHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	tb := blockchain.Transaction{}
	transaction := &tb
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(transaction)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	blockchain.CreateChain(transaction)
}

func Run() {
  addr := "127.0.0.1:8081"
  router := httptreemux.NewContextMux()
	blockchain.GenesisBlock()
  router.Handler(http.MethodPost, "/newTransaction/", &CreateBlockHandler{})

  log.Printf("Running web server on: http://%s\n", addr)
  log.Fatal(http.ListenAndServe(addr, router))
}
