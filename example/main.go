package main

import (
	"log"
	"net/http"

	"github.com/kevinpollet/nego"
)

func main() {
	handler := http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.Header().Add("Content-Charset", nego.NegotiateContentCharset(req, "utf-8"))
		rw.Header().Add("Content-Language", nego.NegotiateContentLanguage(req, "fr", "en"))
		rw.Header().Add("Content-Type", nego.NegotiateContentType(req, "text/plain"))

		rw.WriteHeader(http.StatusOK)
	})

	log.Fatal(http.ListenAndServe(":8080", handler))
}
