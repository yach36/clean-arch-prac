#!/bin/bash

protoc -I delivery/grpc/proto \
    --go_out delivery/grpc/proto \
    --go_opt paths=source_relative \
    --go-grpc_out delivery/grpc/proto \
    --go-grpc_opt paths=source_relative \
    delivery/grpc/proto/user.proto
