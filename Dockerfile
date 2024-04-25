FROM golang:alpine AS builder

WORKDIR /app

ADD go.mod .

COPY . .

RUN go build -o binary cmd/main.go

FROM alpine

WORKDIR /app

COPY --from=builder /app/binary /app/binary

CMD ["./binary"]