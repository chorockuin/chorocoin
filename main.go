package main

import (
	"github.com/chorockuin/chorocoin/blockchain"
)

func main() {
	// cli.Start()
	blockchain.Blockchain().AddBlock("First")
	blockchain.Blockchain().AddBlock("Second")
	blockchain.Blockchain().AddBlock("Third")
}
