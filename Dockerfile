FROM golang:1.13.1-alpine as builder

WORKDIR /app
COPY . .

RUN CGO_ENABLED=0 GOOS=`go env GOHOSTOS` GOARCH=`go env GOHOSTARCH` go build -o plugin


FROM alpine:latest

WORKDIR /

RUN apk add --no-cache ca-certificates
COPY --from=builder /app/plugin .

ENTRYPOINT ["/plugin"]