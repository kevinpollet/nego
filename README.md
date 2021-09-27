# nego

[![Build Status](https://github.com/kevinpollet/nego/workflows/build/badge.svg)](https://github.com/kevinpollet/nego/actions)
[![GoDoc](https://godoc.org/github.com/kevinpollet/nego?status.svg)](https://pkg.go.dev/github.com/kevinpollet/nego)
[![Go Report Card](https://goreportcard.com/badge/github.com/kevinpollet/nego?burst=)](https://goreportcard.com/report/github.com/kevinpollet/nego)

Go package providing an implementation of [HTTP Content Negotiation](https://en.wikipedia.org/wiki/Content_negotiation) compliant with [RFC 7231](https://tools.ietf.org/html/rfc7231#section-5.3).

As defined in [RFC 7231](https://tools.ietf.org/html/rfc7231#section-5.3) the following request headers are sent by a user agent to engage in a proactive negotiation of the response content: `Accept`, `Accept-Charset`, `Accept-Language` and `Accept-Encoding`.
This package provides convenient functions to negotiate the best and acceptable response content `type`, `charset`, `language` and `encoding`.

## Installation

Install using `go get github.com/kevinpollet/nego`.

## Usage

```go
package main

import (
	"log"
	"net/http"
	"github.com/kevinpollet/nego"
)

func main() {
	handler := http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		contentCharset := nego.NegotiateContentCharset(req, "utf-8")
		contentEncoding := nego.NegotiateContentEncoding(req, "gzip", "deflate")
		contentLanguage := nego.NegotiateContentLanguage(req, "fr", "en")
		contentType := nego.NegotiateContentType(req, "text/html", "text/plain")
		...
	})
}
```

## Contributing

Contributions are welcome!

Want to file a bug, request a feature or contribute some code?

1. Check out the [Code of Conduct](./CODE_OF_CONDUCT.md).
2. Check for an existing [issue](https://github.com/kevinpollet/nego/issues) corresponding to your bug or feature request.
3. Open an issue to describe your bug or feature request.

## License

[MIT](./LICENSE.md)
