package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	var port string
	var verbose bool
	flag.BoolVar(&verbose, "verbose", false, "Print all headers")
	flag.StringVar(&port, "port", ":8080", "Port to listen (prepended by colon), i.e. :8080")
	flag.Parse()

	name, err := os.Hostname()
	if err != nil {
		log.Fatalf("Error: %v\n", err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		if verbose {
			log.Printf("Request from [%s]", req.RemoteAddr)
			for k, v := range req.Header {
				fmt.Printf("%s: %s\n", k, v)
			}
			fmt.Println()
		}
		w.Write([]byte(fmt.Sprintf("Response from %s%s", name, port)))
	})
	log.Printf("HTTP server is listening on port %s, verbose = %v\n", port, verbose)
	http.ListenAndServe(port, nil)
}
