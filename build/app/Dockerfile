FROM golang:1.22

ENV GOPATH=/go

RUN apt-get update
RUN apt-get install -y git

WORKDIR ${GOPATH}/src 

RUN go install github.com/cosmtrek/air@v1.52.1

ENTRYPOINT [ "air" ]