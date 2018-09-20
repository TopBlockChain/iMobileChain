.PHONY: build cluster test

default: build

build:
	go build -o bin/swarm    github.com/blockchain/iMobileChain/cmd/swarm
	go build -o bin/geth     github.com/blockchain/iMobileChain/cmd/geth
	go build -o bin/bootnode github.com/blockchain/iMobileChain/cmd/bootnode

cluster: build
	scripts/boot-cluster.sh

test:
	go test -v github.com/blockchain/iMobileChain/swarm/...
