FROM golang:1.23-alpine AS udpz-builder
LABEL builder="true"

WORKDIR /go/src/

COPY cmd/ cmd/
COPY internal/ internal/
COPY main.go go.mod go.sum ./

ENV CGO_ENABLED=0

RUN go mod download
RUN go build -ldflags="-s -w" -o /go/bin/udpz


FROM alpine:3 AS udpz
COPY --from="udpz-builder" /go/bin/udpz /usr/local/bin/udpz

RUN ulimit -n 60000
WORKDIR /output
ENTRYPOINT ["/usr/local/bin/udpz"]
