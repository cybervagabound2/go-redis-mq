#!/bin/bash
# This script will install redis,
# clone repository, build application
# and run it.

# Redis installation part
cd $HOME
wget http://download.redis.io/redis-stable.tar.gz
tar xvzf redis-stable.tar.gz
cd redis-stable
make

sudo cp src/redis-server /usr/local/bin/
sudo cp src/redis-cli /usr/local/bin/

redis-server &

# Application installation part
go get github.com/cybervagabound2/go-redis-pubsub
cd $GOPATH/src/github.com/cybervagabound2/go-redis-pubsub
go build main.go
./main
