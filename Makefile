run:
	go run main.go -l 10000 -secio

setup: setup-go setup-p2p-go

setup-go:
	go get github.com/dimfeld/httptreemux
	go get github.com/onsi/gomega/...
	go get github.com/davecgh/go-spew/spew

setup-p2p-go:
	go get -u -d github.com/libp2p/go-libp2p/...
	cd $GOPATH/src/github.com/libp2p/go-libp2p && make && make deps
