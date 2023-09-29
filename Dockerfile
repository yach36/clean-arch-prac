FROM golang:1.21.1-alpine

RUN apk update && apk --no-cache add tzdata bash curl unzip
RUN curl -LO https://github.com/protocolbuffers/protobuf/releases/download/v24.3/protoc-24.3-linux-x86_64.zip && \
    unzip protoc-24.3-linux-x86_64.zip -d /root/.local && rm -rf protoc-24.3-linux-x86_64.zip
ENV PATH "$PATH:/root/.local/bin"

WORKDIR /go/src/github.com/yach36/clean-arch-prac

COPY go.mod go.sum ./
RUN go mod download && go mod verify
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest && \
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

COPY . .
RUN go build -o /server main.go

CMD ["/server"]
