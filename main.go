package main

import (
	"github.com/chorockuin/chorocoin/blockchain"
)

func main() {
	chain := blockchain.GetBlockchain()
	chain.Add_block("second_block")
	chain.Add_block("third_block")
	chain.Add_block("fourth_block")
	chain.List_blocks()
}
