
FROM golang:1.23 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o go-basic-server .

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/go-basic-server .

EXPOSE 8081

CMD ["./go-basic-server"]
