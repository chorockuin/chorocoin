package main

import (
	"github.com/chorockuin/chorocoin/blockchain"
	"github.com/chorockuin/chorocoin/cli"
)

func main() {
	blockchain.Blockchain()
	cli.Start()
}
