FROM golang:1.17 as builder

WORKDIR /workspace

COPY go.mod go.mod
COPY go.sum go.sum

RUN go mod download
COPY ./main.go /workspace/

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on go build -a -o httpserver main.go

FROM scratch
COPY --from=builder /workspace/httpserver /
EXPOSE 80
ENTRYPOINT ["/httpserver"]