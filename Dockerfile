FROM golang:1.21.1-alpine

RUN apk update && apk add bash

RUN mkdir /go/src/app
WORKDIR /go/src/app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .

CMD ["./cmd.sh"]
