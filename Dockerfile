FROM golang:alpine AS builder

RUN mkdir -p $GOPATH/src/jonasmartinez.com/docker-node-hostname

COPY main.go $GOPATH/src/jonasmartinez.com/docker-node-hostname
WORKDIR $GOPATH/src/jonasmartinez.com/docker-node-hostname

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags="-w -s" -o /usr/local/bin/docker-node-hostname

FROM scratch
COPY --from=builder /usr/local/bin/docker-node-hostname /usr/local/bin/docker-node-hostname

EXPOSE 5000

ENTRYPOINT ["docker-node-hostname"]