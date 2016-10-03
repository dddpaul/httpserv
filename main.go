package main

import (
	"bytes"
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
		w.Write([]byte(fmt.Sprintf("Response from %s%s", name, port)))
	})
	log.Printf("HTTP server is listening on port %s, verbose = %v\n", port, verbose)
	http.ListenAndServe(port, nil)
}
