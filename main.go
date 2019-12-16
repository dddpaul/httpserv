package main

import (
	"bytes"
	"flag"
	"fmt"
	"github.com/unrolled/logger"
	"log"
	"net/http"
)

var prefix string
var port string
var message string
var verbose bool
var headers bool
var l *logger.Logger

var app http.Handler = http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
	var buf bytes.Buffer
	_, ok := req.Header["X-Headers-Logging-Enabled"]
	if headers || ok {
		buf.WriteString("(Headers) ")
		for k, v := range req.Header {
			buf.WriteString(fmt.Sprintf("%s: %s, ", k, v))
		}
		buf.Truncate(buf.Len() - 2)
		l.Print(buf.String())
	}
	if _, err := w.Write([]byte(message)); err != nil {
		l.Println(err)
	}
})

func main() {
	flag.StringVar(&prefix, "prefix", "httpserv", "Logging prefix")
	flag.BoolVar(&verbose, "verbose", false, "Print request details")
	flag.BoolVar(&headers, "headers", false, "Print request headers")
	flag.StringVar(&port, "port", ":8080", "Port to listen (prepended by colon), i.e. :8080")
	flag.StringVar(&message, "message", "HTTP OK", "Server response")
	flag.Parse()

	l = logger.New(logger.Options{
		Prefix:               prefix,
		RemoteAddressHeaders: []string{"X-Forwarded-For"},
		OutputFlags:          log.LstdFlags,
	})

	if verbose {
		app = l.Handler(app)
	}

	l.Printf("HTTP server is listening on port %s, verbose = %v, log headers = %v\n", port, verbose, headers)
	l.Fatalln("ListenAndServe:", http.ListenAndServe(port, app))
}
