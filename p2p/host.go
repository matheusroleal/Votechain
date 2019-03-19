package p2p

import (
  "fmt"
  "log"
  "context"

  host "github.com/libp2p/go-libp2p-host"
  libp2p "github.com/libp2p/go-libp2p"
  ma "github.com/multiformats/go-multiaddr"
)

func MakeHost(listenPort int) (host.Host, error){

	opts := []libp2p.Option{
		libp2p.ListenAddrStrings(fmt.Sprintf("/ip4/127.0.0.1/tcp/%d", listenPort)),
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
