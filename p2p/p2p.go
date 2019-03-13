package p2p

import (
  "bufio"
  "io"
  "fmt"
  "crypto/rand"
  "context"
  "log"
  mrand "math/rand"

  host "github.com/libp2p/go-libp2p-host"
  crypto "github.com/libp2p/go-libp2p-crypto"
  libp2p "github.com/libp2p/go-libp2p"
  ma "github.com/multiformats/go-multiaddr"
  net "github.com/libp2p/go-libp2p-net"
)

func MakeHost(listenPort int, secio bool, randseed int64) (host.Host, error){
  var r io.Reader
  if randseed == 0 {
  	r = rand.Reader
  } else {
  	r = mrand.New(mrand.NewSource(randseed))
  }

	priv, _, err := crypto.GenerateKeyPairWithReader(crypto.RSA, 2048, r)
	if err != nil {
		return nil, err
	}

	opts := []libp2p.Option{
		libp2p.ListenAddrStrings(fmt.Sprintf("/ip4/127.0.0.1/tcp/%d", listenPort)),
		libp2p.Identity(priv),
    libp2p.DisableRelay(),
	}

	basicHost, err := libp2p.New(context.Background(), opts...)
	if err != nil {
		return nil, err
	}

	hostAddr, _ := ma.NewMultiaddr(fmt.Sprintf("/ipfs/%s", basicHost.ID().Pretty()))

	addr := basicHost.Addrs()[0]
	fullAddr := addr.Encapsulate(hostAddr)
	log.Printf("I am %s\n", fullAddr)

	return basicHost, nil
}

func HandleStream(stream net.Stream) {
  log.Println("Got a new stream!")

	rw := bufio.NewReadWriter(bufio.NewReader(stream), bufio.NewWriter(stream))

	go ReadData(rw)
	go WriteData(rw)
}
