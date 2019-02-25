package main

import (
  "bufio"
	"context"
	"flag"
	"fmt"

	golog "github.com/ipfs/go-log"
	gologging "github.com/whyrusleeping/go-logging"
  blockchain "github.com/matheusroleal/Votechain/blockchain"
  p2p "github.com/matheusroleal/Votechain/p2p"
  mdns "github.com/matheusroleal/Votechain/mdns"
)

func main() {
  blockchain.GenesisBlock()

	golog.SetAllLoggers(gologging.INFO) // Change to DEBUG for extra info

	listenF := flag.Int("l", 0, "wait for incoming connections")
	secio := flag.Bool("secio", false, "enable secio")
	seed := flag.Int64("seed", 0, "set random seed for id generation")
	flag.Parse()

	if *listenF == 0 {
		fmt.Println("Please provide a port to bind on with -l")
	}

	ha, err := p2p.MakeHost(*listenF, *secio, *seed)
	if err != nil {
		fmt.Println(err)
	}

	ha.SetStreamHandler("/p2p/1.0.0", p2p.HandleStream)

	peerChan := mdns.InitMDNS(context.Background(), ha, "meetmehere")

	peer := <-peerChan // will block untill we discover a peer
	fmt.Println("Found peer:", peer, ", connecting")

	err = ha.Connect(context.Background(), peer)
  if err != nil {
		fmt.Println("Connection failed:", err)
	}

	// open a stream, this stream will be handled by handleStream other end
	stream, err := ha.NewStream(context.Background(),peer.ID, "/p2p/1.0.0")

	if err != nil {
		fmt.Println("Stream open failed", err)
	} else {
		rw := bufio.NewReadWriter(bufio.NewReader(stream), bufio.NewWriter(stream))

  	go p2p.WriteData(rw)
		go p2p.ReadData(rw)
	}

	select {}
}
