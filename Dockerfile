FROM golang:1.21-alpine

RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main cmd/main.go

RUN find . -mindepth 1 -not -name 'main' -exec rm -rf {} +

EXPOSE 9000

CMD ["./main"]
