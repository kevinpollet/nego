package main

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/kevinpollet/negotiate"
)

func main() {
	negotiateHandler := http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		contentCharset := negotiate.ContentCharset(req, "UTF-8")
		contentEncoding := negotiate.ContentEncoding(req, "gzip", "deflate")
		contentLanguage := negotiate.ContentLanguage(req, "fr", "en")

		rw.WriteHeader(http.StatusOK)
		io.WriteString(rw, fmt.Sprintln("Content-Charset:", contentCharset))   // nolint
		io.WriteString(rw, fmt.Sprintln("Content-Encoding:", contentEncoding)) // nolint
		io.WriteString(rw, fmt.Sprintln("Content-Language:", contentLanguage)) // nolint
	})

	log.Fatal(http.ListenAndServe(":8080", negotiateHandler))
}
