package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/chorockuin/chorocoin/blockchain"
)

const port string = ":4000"

type home_data struct {
	Page_title string
	Blocks     []*blockchain.Block
}

func home(rw http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("template/home.html"))
	data := home_data{"Chorockuin Home!", blockchain.GetBlockchain().AllBlocks()}
	tmpl.Execute(rw, data)
}

func main() {
	http.HandleFunc("/", home)
	fmt.Printf("Listening on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
