package cli

import (
	"flag"
	"fmt"
	"os"
	"runtime"

	"github.com/chorockuin/chorocoin/explorer"
	"github.com/chorockuin/chorocoin/rest"
)

func usage() {
	fmt.Printf("Welcome to chorocoin\n\n")
	fmt.Printf("Please use the following flags:\n\n")
	fmt.Printf("-port 		Set the port of the server\n")
	fmt.Printf("-mode 		Choose between 'html' and 'rest'\n")
	runtime.Goexit() // run defer
}

func Start() {
	fmt.Println(os.Args)

	if len(os.Args) < 2 {
		usage()
	}

	port := flag.Int("port", 4000, "Sets the port of the server (default 4000)")
	mode := flag.String("mode", "rest", "Choose between 'html' and 'rest'")
	flag.Parse()
	switch *mode {
	case "rest":
		rest.Start(*port)
	case "html":
		explorer.Start(*port)
	default:
		usage()
	}
}
