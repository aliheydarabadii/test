FROM golang:1.19 AS builder

WORKDIR /go/src/test
COPY . .

RUN GO111MODULE=on CGO_ENABLED=1 GOOS=linux go build -ldflags="-extldflags=-static" -a -installsuffix nocgo -tags=nomsgpack -o /app main.go

FROM debian:buster-slim

# http server and profiler and prometheus
EXPOSE 3000
COPY --from=builder /app ./

ENTRYPOINT ["./app"]
