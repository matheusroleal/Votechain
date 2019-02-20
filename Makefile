setup-go:
	go get github.com/dimfeld/httptreemux
	go get github.com/onsi/gomega/...
	go get github.com/davecgh/go-spew/spew

run:
	go run main.go
