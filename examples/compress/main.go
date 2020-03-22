package main

import (
	"compress/gzip"
	"io"
	"log"
	"net/http"

	"github.com/kevinpollet/nego"
)

type compressResponseWriter struct {
	io.Writer
	http.ResponseWriter
}

func (w *compressResponseWriter) Write(p []byte) (n int, err error) {
	return w.Writer.Write(p)
}

func main() {
	compressHandler := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			contentEncoding := nego.ContentEncoding(req, "gzip")

			if contentEncoding == "" {
				rw.WriteHeader(http.StatusNotAcceptable)
				return
			}

			if contentEncoding == "gzip" {
				cw := gzip.NewWriter(rw)
				defer cw.Close()

				rw.Header().Add("Content-Encoding", contentEncoding)
				rw = &compressResponseWriter{cw, rw}
			}

			next.ServeHTTP(rw, req)
		})
	}

	handler := http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.Header().Add("Content-Type", "text/plain")
		rw.WriteHeader(http.StatusOK)
		rw.Write([]byte("Hello, World!!")) // nolint
	})

	log.Fatal(http.ListenAndServe(":8080", compressHandler(handler)))
}
