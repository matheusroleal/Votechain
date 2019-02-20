package server

import (
	"log"
  "encoding/json"
	"net/http"
	"time"

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
	router := httptreemux.NewContextMux()
	server := &http.Server{
		Addr:         "0.0.0.0:8081",
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  time.Second,
	}

  router.Handler(http.MethodPost, "/newTransaction/", &CreateBlockHandler{})
	log.Fatal(server.ListenAndServe())

	blockchain.GenesisBlock()
}
