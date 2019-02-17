[![CircleCI](https://circleci.com/gh/cybervagabound2/go-redis-pubsub/tree/master.svg?style=svg)](https://circleci.com/gh/cybervagabound2/go-redis-pubsub/tree/master)
# go-redis-pubsub
Simple application that implement Publisher/Subcriber pattern using Golang and Redis.

## Installation

### Manual Install
- You need Golang installed on your machine:
https://golang.org
- Clone repository (all dependicies will be installed)
`$ go get github.com/cybervagabound2/go-redis-pubsub`
- Install Redis: https://redis.io/topics/quickstart
- Run redis server:
`$ redis-server &`
- In project directory, build application:
`$ go build main.go`
- Run application:
`$ ./main`

### Install vis .sh script
`$ chmod +x install.sh`  
`$ ./install.sh`

### Install via Docker

## Screencast
- running install.sh
https://asciinema.org/a/QcpbXSM8RPEZ3wtC3AhGLca4r
## TODO
- [ ] Enable Docker support
- [ ] Setup CI/CD
- [ ] Create websocket implementation
