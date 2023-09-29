#!/bin/bash

protoc -I delivery/grpc/user_grpc \
    --go_out delivery/grpc/user_grpc \
    --go_opt paths=source_relative \
    --go-grpc_out delivery/grpc/user_grpc \
    --go-grpc_opt paths=source_relative \
    delivery/grpc/user_grpc/user.proto
