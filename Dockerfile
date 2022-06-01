FROM golang:1.8 as builder
WORKDIR /go/src/github.com/gangsta/goport/
ADD ./main.go /go/src/github.com/gangsta/goport/main.go
RUN CGO_ENABLED=0 GOOS=linux go build -o ./goport /go/src/github.com/gangsta/goport/main.go
FROM alpine:latest
RUN apk --no-cache add ca-certificates
RUN apk add --no-cache curl
COPY --from=builder /go/src/github.com/gangsta/goport/goport /app/
RUN addgroup --gid 3033 goport
RUN adduser -h /app -s /bin/sh -G goport -u 3033 -D goport
RUN chown goport:goport -R /app
USER goport
WORKDIR /app
ENTRYPOINT ["/app/goport"]