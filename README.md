# clean-arch-prac
clean architecture with golang

## architecture diagram
```mermaid
flowchart LR
  id1(cmd)
  id2(infra)
  id3(usecase)
  id4(delivery)
  id5(domain)
  subgraph s1 [Frameworks & Drivers]
    direction BT
    id6[(PostgreSQL)] -.- id2
  end
  subgraph s2 [Interface Adapters]
    direction TB
    id7([REST API]) -.- id4
    id8([gRPC]) -.- id4
  end
  subgraph s3 [Application Business Rules]
    id3
  end
  subgraph s4 [Enterprise Business Rules]
    id5
  end
  subgraph s5 [Entrypoint]
    id1
  end
  s5 --> s1
  s5 --> s2 --> s3 --> s4
```

## setup
```bash
# clone
git clone https://github.com/yach36/clean-arch-prac.git
cd clean-arch-prac
```
```bash
# set .env file
touch .env
```
```.env
# .envに下記を入力(本リポジトリは実運用しないので記載しているが, 普通は公開してはいけない)
POSTGRES_USER=user
POSTGRES_PASSWORD=password
POSTGRES_DB=test
```
```bash
# start server
docker compose build
docker compose up -d
```

## grpc server
https://github.com/yach36/clean-arch-prac/blob/main/docs/grpc.md

## REST API server
https://github.com/yach36/clean-arch-prac/blob/main/docs/api.md
