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
  // addr := "127.0.0.1:8081"
	router := httptreemux.NewContextMux()
	server := &http.Server{
		Addr:         "0.0.0.0:8081",
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  time.Second,
	}
	blockchain.GenesisBlock()
  router.Handler(http.MethodPost, "/newTransaction/", &CreateBlockHandler{})

  // log.Printf("Running web server on: http://%s\n", addr)
  // log.Fatal(http.ListenAndServe(addr, router))
	log.Fatal(server.ListenAndServe())
}
