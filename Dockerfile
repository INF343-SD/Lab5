FROM golang:latest

WORKDIR /app

COPY /proto /app/proto/
COPY /Broker .
COPY go.mod .
COPY go.sum .

RUN go build broker.go

EXPOSE 5000
CMD ["./broker"]