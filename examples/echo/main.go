package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/kevinpollet/nego"
)

func main() {
	handler := http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.WriteHeader(http.StatusOK)

		fmt.Fprintln(rw, "Content-Charset:", nego.NegotiateContentCharset(req, "utf-8"))
		fmt.Fprintln(rw, "Content-Encoding:", nego.NegotiateContentEncoding(req, "gzip", "deflate"))
		fmt.Fprintln(rw, "Content-Language:", nego.NegotiateContentLanguage(req, "fr", "en"))
		fmt.Fprintln(rw, "Content-Type:", nego.NegotiateContentType(req, "text/plain"))
	})

	log.Fatal(http.ListenAndServe(":8080", handler))
}
