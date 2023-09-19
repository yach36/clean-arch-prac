FROM golang:1.21.1-alpine

RUN apk update && apk --no-cache add tzdata bash

WORKDIR /go/src/github.com/yach36/clean-arch-prac

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -o /server main.go

CMD ["/server"]
