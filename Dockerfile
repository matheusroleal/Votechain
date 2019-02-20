FROM golang:latest

ADD . /go/src/github.com/matheusroleal/Votechain

RUN go get github.com/dimfeld/httptreemux
RUN	go get github.com/onsi/gomega/...
RUN	go get github.com/davecgh/go-spew/spew

RUN go install github.com/matheusroleal/Votechain

ENTRYPOINT /go/bin/Votechain

EXPOSE 8081
