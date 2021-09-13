package main

import (
	"github.com/chorockuin/chorocoin/explorer"
	"github.com/chorockuin/chorocoin/rest"
)

func main() {
	go explorer.Start(3000)
	rest.Start(4000)
}
