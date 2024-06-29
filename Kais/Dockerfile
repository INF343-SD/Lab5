FROM golang:latest

WORKDIR /app

COPY ../proto/ /app/proto/
COPY ../Kais/ .
COPY ../go.mod .
COPY ../go.sum .

RUN go build kais.go

EXPOSE 5002
CMD ["./kais"]