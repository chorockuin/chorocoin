package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/chorockuin/chorocoin/blockchain"
)

const (
	port         string = ":4000"
	template_dir string = "templates/"
)

var templates *template.Template

type home_data struct {
	Page_title string
	Blocks     []*blockchain.Block
}

func home(rw http.ResponseWriter, r *http.Request) {
	data := home_data{"Chorockuin Home!", blockchain.GetBlockchain().AllBlocks()}
	templates.ExecuteTemplate(rw, "home", data)
}

func main() {
	templates = template.Must(template.ParseGlob(template_dir + "pages/*.gohtml"))
	templates = template.Must(templates.ParseGlob(template_dir + "partials/*.gohtml"))
	http.HandleFunc("/", home)
	fmt.Printf("Listening on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
