FROM golang:latest

WORKDIR /app

COPY ../proto/ /app/proto/
COPY ../Ingeniero/ .
COPY ../go.mod .
COPY ../go.sum .

RUN go build ingeniero.go

EXPOSE 5003
CMD ["./ingeniero"]