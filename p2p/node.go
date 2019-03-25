package p2p

import (
  "context"
  "log"

  floodsub "github.com/libp2p/go-libp2p-pubsub"
  host "github.com/libp2p/go-libp2p-host"
  blockchain "github.com/matheusroleal/Votechain/blockchain"
  mdns "github.com/matheusroleal/Votechain/mdns"
)

type Node struct {
	p2pNode    host.Host
	blockchain *blockchain.Blockchain
  pubsub     *floodsub.PubSub
}

func CreateNewNode(ctx context.Context, listenF int) *Node {
  var node Node
  blkch := blockchain.NewBlockchain()

  if listenF == 0 {
    log.Println("Please provide a port to bind on with -l")
  }

  ha, err := MakeHost(listenF)
  if err != nil {
    log.Println(err)
  }

  pubsub, err := floodsub.NewFloodSub(ctx, ha)
	if err != nil {
		panic(err)
	}

	node.blockchain = blkch
  node.pubsub = pubsub

	node.ListenBlocks(ctx)

  peerChan := mdns.InitMDNS(context.Background(), ha, "meetmehere")

  peer := <-peerChan // will block untill we discover a peer
  log.Println("Found peer:", peer, ", connecting")

  err = ha.Connect(context.Background(), peer)
  if err != nil {
    log.Println("Connection failed:", err)
  }

	node.p2pNode = ha

  return &node
}
