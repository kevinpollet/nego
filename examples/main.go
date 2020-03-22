package main

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/kevinpollet/nego"
)

func main() {
	handler := http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		contentCharset := nego.ContentCharset(req, "utf-8")
		contentEncoding := nego.ContentEncoding(req, "gzip", "deflate")
		contentLanguage := nego.ContentLanguage(req, "fr", "en")
		contentType := nego.ContentType(req, "text/html", "text/plain")

		rw.WriteHeader(http.StatusOK)

		io.WriteString(rw, fmt.Sprintln("Content-Charset:", contentCharset))   // nolint
		io.WriteString(rw, fmt.Sprintln("Content-Encoding:", contentEncoding)) // nolint
		io.WriteString(rw, fmt.Sprintln("Content-Language:", contentLanguage)) // nolint
		io.WriteString(rw, fmt.Sprintln("Content-Type:", contentType))         // nolint
	})

	log.Fatal(http.ListenAndServe(":8080", handler))
}
