FROM golang:1.23.4-alpine3.21

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .

ENV GIN_MODE=release

RUN go build -o main .

CMD ["/app/main"]
