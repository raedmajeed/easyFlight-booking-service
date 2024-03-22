FROM golang:latest

WORKDIR go/src/app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

WORKDIR ./cmd/api

RUN CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o main

CMD ["./main"]