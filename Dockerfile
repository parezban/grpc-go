FROM golang:1.8 as builder
WORKDIR /go/src/github.com/parezban/grpc-go/
ADD ./main.go /go/src/github.com/parezban/grpc-go/main.go
RUN CGO_ENABLED=0 GOOS=linux go build -o ./grpc-go /go/src/github.com/parezban/grpc-go/main.go
FROM alpine:latest
RUN apk --no-cache add ca-certificates
RUN apk add --no-cache curl
COPY --from=builder /go/src/github.com/parezban/grpc-go /app/
RUN addgroup --gid 3000 grpc-go
RUN adduser -h /app -s /bin/sh -G grpc-go -u 3000 -D grpc-go
RUN chown grpc-go:grpc-go -R /app
USER grpc-go
WORKDIR /app
ENTRYPOINT ["/app/grpc-go"]