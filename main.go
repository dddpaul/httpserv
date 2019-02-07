package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {
	var port string
	var message string
	var verbose bool
	flag.BoolVar(&verbose, "verbose", false, "Print all headers")
	flag.StringVar(&port, "port", ":8080", "Port to listen (prepended by colon), i.e. :8080")
	flag.StringVar(&message, "message", "HTTP OK", "Server response")
	flag.Parse()

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		var buf bytes.Buffer
		buf.WriteString(fmt.Sprintf("%s %s from [%s]\n", req.Method, req.RequestURI, req.RemoteAddr))
		_, ok := req.Header["X-Logging-Enabled"]
		if verbose || ok {
			for k, v := range req.Header {
				buf.WriteString(fmt.Sprintf("%s: %s\n", k, v))
			}
			buf.WriteString("\n")
		}
		log.Print(buf.String())
		w.Write([]byte(message))
	})
	log.Printf("HTTP server is listening on port %s, verbose = %v\n", port, verbose)
	log.Fatalln("ListenAndServe:", http.ListenAndServe(port, nil))
}
