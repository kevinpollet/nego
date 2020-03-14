# negotiate <!-- omit in toc -->

[![Build Status](https://github.com/kevinpollet/negotiate/workflows/build/badge.svg)](https://github.com/kevinpollet/negotiate/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/kevinpollet/negotiate?burst=)](https://goreportcard.com/report/github.com/kevinpollet/negotiate)
[![GoDoc](https://godoc.org/github.com/kevinpollet/negotiate?status.svg)](https://pkg.go.dev/github.com/kevinpollet/negotiate)
[![Conventional Commits](https://img.shields.io/badge/Conventional%20Commits-1.0.0-yellow.svg)](https://conventionalcommits.org)
[![License](https://img.shields.io/github/license/kevinpollet/negotiate)](./LICENSE.md)

Package `negotiate` provides an [RFC 7231](https://tools.ietf.org/html/rfc7231#section-5.3) compliant implementation of [HTTP Content Negotiation](https://en.wikipedia.org/wiki/Content_negotiation).

As defined in [RFC 7231](https://tools.ietf.org/html/rfc7231#section-5.3) the following request headers are sent by a user agent to engage in proactive negotiation of the response content: `Accept`, `Accept-Charset`, `Accept-Language` and `Accept-Encoding`. This package provides convenient functions to negotiate the best and acceptable response content `type`, `charset`, `language` and `encoding` that should be returned by the `HTTP` server.

## Table of Contents <!-- omit in toc -->

- [Install](#install)
- [Usage](#usage)
- [Examples](#examples)
- [Contributing](#contributing)
- [License](#license)

## Install

```shell
go get github.com/kevinpollet/negotiate
```

## Usage

```go
package main

import (
	"log"
	"net/http"

	"github.com/kevinpollet/negotiate"
)

func main() {
	handler := http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		contentCharset := negotiate.ContentCharset(req, "UTF-8")
		contentEncoding := negotiate.ContentEncoding(req, "br", "gzip", "deflate")
		contentLanguage := negotiate.ContentLanguage(req, "en", "en-US")
		contentType := negotiate.ContentType(req, "text/html", "text/plain")

		rw.WriteHeader(http.StatusOK)
	})

	http.ListenAndServe(":8080", handler)
}
```

## Examples

The [examples](./Examples) directory contains an [example](./examples/main.go) which returns the content `charset`, `encoding` and `language` negotiated with the `Accept-Charset`, `Accept-Language` and `Accept-Encoding` headers present in the request.

```shell
$ go run examples/main.go
$ curl localhost:8080 -H "Accept-Charset: utf-8, utf-16" -H "Accept-Language: fr;q=0.3, en" -H "Accept-Encoding: br, gzip"

Content-Charset: utf-8
Content-Encoding: gzip
Content-Language: en
```

## Contributing

Contributions are welcome!

Want to file a bug, request a feature or contribute some code?

1. Check out the [Code of Conduct](./CODE_OF_CONDUCT.md).
2. Check for an existing issue corresponding to your bug or feature request.
3. Open an issue to describe your bug or feature request.

## License

[MIT](./LICENSE.md)
