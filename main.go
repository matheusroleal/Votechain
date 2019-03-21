package main

import (
	"context"
	"log"
	"net/http"
	"encoding/json"
	"flag"
	"os"

	golog "github.com/ipfs/go-log"
  gologging "github.com/whyrusleeping/go-logging"
	p2p "github.com/matheusroleal/Votechain/p2p"
)

func main() {
  golog.SetAllLoggers(gologging.INFO) // Change to DEBUG for extra info

	listenF := flag.Int("l", 0, "wait for incoming connections")
	flag.Parse()

	if *listenF == 0 {
		log.Println("Please provide a port to bind on with -l")
	}

	ctx := context.Background()

	node := p2p.CreateNewNode(ctx,*listenF)

	http.HandleFunc("/sendvote", func(w http.ResponseWriter, r *http.Request) {
		from := r.FormValue("from")
		to := r.FormValue("to")

		log.Println("Executing vote", from, to)

		tx := p2p.Transaction{
			Key: from,
			Vote: to,
		}

		err := json.NewEncoder(w).Encode(node.BroadcastBlock(&tx))
		if err != nil {
			panic(err)
		}
	})

	http.HandleFunc("/getnewaddress", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Call getnewaddress")

		err := json.NewEncoder(w).Encode(node.GetNewAddress())
		if err != nil {
			panic(err)
		}
	})

	http.ListenAndServe(":"+os.Getenv("VOTECHAIN_PORT"), nil)
}
