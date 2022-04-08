FROM golang:1.17 as builder

WORKDIR /workspace

COPY go.mod go.mod
COPY go.sum go.sum

RUN go mod download
COPY ./main.go /workspace/

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on go build -a -o httpserver main.go

FROM debian:stable

RUN useradd -m httpserver -u 1000 && apt-get update && apt-get install bash curl -y && apt-get autoclean && rm -rf /var/lib/apt/lists/*

USER httpserver:httpserver
RUN mkdir -p /home/httpserver

WORKDIR /home/httpserver
ENV VERSION=v0.1.0
COPY --from=builder /workspace/httpserver /home/httpserver/httpserver
ENTRYPOINT [ "/home/httpserver/httpserver" ]

