package main

import (
  "bufio"
	"context"
	"flag"
	"fmt"
	"log"

	golog "github.com/ipfs/go-log"
	peer "github.com/libp2p/go-libp2p-peer"
	pstore "github.com/libp2p/go-libp2p-peerstore"
	ma "github.com/multiformats/go-multiaddr"
	gologging "github.com/whyrusleeping/go-logging"
  blockchain "github.com/matheusroleal/Votechain/blockchain"
  p2p "github.com/matheusroleal/Votechain/p2p"
)

func main() {

	golog.SetAllLoggers(gologging.INFO) // Change to DEBUG for extra info

	listenF := flag.Int("l", 0, "wait for incoming connections")
	target := flag.String("d", "", "target peer to dial")
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

	if *target == "" {
    blockchain.GenesisBlock()

		fmt.Println("listening for connections")
		ha.SetStreamHandler("/p2p/1.0.0", p2p.HandleStream)

		select {}

	} else {
		ha.SetStreamHandler("/p2p/1.0.0", p2p.HandleStream)

		ipfsaddr, err := ma.NewMultiaddr(*target)
		if err != nil {
			fmt.Println(err)
		}

		pid, err := ipfsaddr.ValueForProtocol(ma.P_IPFS)
		if err != nil {
			fmt.Println(err)
		}

		peerid, err := peer.IDB58Decode(pid)
		if err != nil {
			fmt.Println(err)
		}

		targetPeerAddr, _ := ma.NewMultiaddr(
			fmt.Sprintf("/ipfs/%s", peer.IDB58Encode(peerid)))
		targetAddr := ipfsaddr.Decapsulate(targetPeerAddr)

		ha.Peerstore().AddAddr(peerid, targetAddr, pstore.PermanentAddrTTL)

		log.Println("opening stream")
		s, err := ha.NewStream(context.Background(), peerid, "/p2p/1.0.0")
		if err != nil {
			fmt.Println(err)
		}

		rw := bufio.NewReadWriter(bufio.NewReader(s), bufio.NewWriter(s))

		go p2p.WriteData(rw)
		go p2p.ReadData(rw)

		select {}
	}
}
