package main

import (
	"flag"
	"github.com/unrolled/logger"
	"log"
	"net/http"
)

var port string
var message string
var verbose bool

var app http.Handler = http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte(message))
})

func main() {
	flag.BoolVar(&verbose, "verbose", false, "Print request details")
	flag.StringVar(&port, "port", ":8080", "Port to listen (prepended by colon), i.e. :8080")
	flag.StringVar(&message, "message", "HTTP OK", "Server response")
	flag.Parse()

	if verbose {
		 app = logger.New(logger.Options{
			Prefix:               "httpserv",
			RemoteAddressHeaders: []string{"X-Forwarded-For"},
			OutputFlags:          log.LstdFlags,
		}).Handler(app)
	}

	log.Printf("HTTP server is listening on port %s, verbose = %v\n", port, verbose)
	log.Fatalln("ListenAndServe:", http.ListenAndServe(port, app))
}
