FROM golang:alpine

WORKDIR /cloud-go
COPY . .

RUN go build -o ./bin/api ./cmd/api

CMD ["/cloud-go/bin/api"]
EXPOSE 8080