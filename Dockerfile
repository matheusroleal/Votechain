FROM golang:latest

ADD . /go/src/gitlab.globoi.com/matheusroleal/Votechain

RUN go install gitlab.globoi.com/matheusroleal/Votechain

ENTRYPOINT /go/bin/Votechain

EXPOSE 8080
