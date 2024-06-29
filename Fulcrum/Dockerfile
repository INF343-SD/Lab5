FROM golang:latest

WORKDIR /app

COPY ../proto/ /app/proto/
COPY ../Fulcrum/ .
COPY ../go.mod .
COPY ../go.sum .

RUN go build fulcrum.go

EXPOSE 5001
CMD ["./fulcrum"]