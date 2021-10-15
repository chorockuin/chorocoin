package main

import (
	"github.com/chorockuin/chorocoin/blockchain"
	"github.com/chorockuin/chorocoin/cli"
	"github.com/chorockuin/chorocoin/db"
)

func main() {
	defer db.Close()
	blockchain.Blockchain()
	cli.Start()
}
