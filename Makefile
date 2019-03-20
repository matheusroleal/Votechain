run:
	go run main.go -l $(shell shuf -i 1-65536 -n 1)

build:
	go build -o votechain-cli ./cli

setup: setup-go setup-p2p-go

setup-go:
	go get github.com/dimfeld/httptreemux
	go get github.com/onsi/gomega/...
	go get github.com/davecgh/go-spew/spew
	go get github.com/libp2p/go-libp2p-pubsub

setup-p2p-go:
	go get -u -d github.com/libp2p/go-libp2p/...
	cd $GOPATH/src/github.com/libp2p/go-libp2p && make && make deps
