package explorer

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/chorockuin/chorocoin/blockchain"
)

var templates *template.Template

const (
	port         string = ":4000"
	template_dir string = "explorer/templates/"
)

type home_data struct {
	PageTitle string
	Blocks    []*blockchain.Block
}

func home(rw http.ResponseWriter, r *http.Request) {
	data := home_data{"Chorockuin Home!", blockchain.GetBlockchain().AllBlocks()}
	templates.ExecuteTemplate(rw, "home", data)
}

func add(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		templates.ExecuteTemplate(rw, "add", nil)
	case "POST":
		r.ParseForm()
		data := r.Form.Get("blockData")
		blockchain.GetBlockchain().Add_block(data)
		http.Redirect(rw, r, "/", http.StatusPermanentRedirect)
	}
}

func Start() {
	templates = template.Must(template.ParseGlob(template_dir + "pages/*.gohtml"))
	templates = template.Must(templates.ParseGlob(template_dir + "partials/*.gohtml"))
	http.HandleFunc("/", home)
	http.HandleFunc("/add", add)
	fmt.Printf("Listening on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
