# nego <!-- omit in toc -->

[![Build Status](https://github.com/kevinpollet/nego/workflows/build/badge.svg)](https://github.com/kevinpollet/nego/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/kevinpollet/nego?burst=)](https://goreportcard.com/report/github.com/kevinpollet/nego)
[![GoDoc](https://godoc.org/github.com/kevinpollet/nego?status.svg)](https://pkg.go.dev/github.com/kevinpollet/nego)
[![Conventional Commits](https://img.shields.io/badge/Conventional%20Commits-1.0.0-yellow.svg)](https://conventionalcommits.org)
[![License](https://img.shields.io/github/license/kevinpollet/nego)](./LICENSE.md)

Package `nego` provides an [RFC 7231](https://tools.ietf.org/html/rfc7231#section-5.3) compliant implementation of [HTTP Content Negotiation](https://en.wikipedia.org/wiki/Content_negotiation).

As defined in [RFC 7231](https://tools.ietf.org/html/rfc7231#section-5.3) the following request headers are sent by a user agent to engage in proactive negotiation of the response content: `Accept`, `Accept-Charset`, `Accept-Language` and `Accept-Encoding`. This package provides convenient functions to negotiate the best and acceptable response content `type`, `charset`, `language` and `encoding`.

## Table of Contents <!-- omit in toc -->

- [Install](#install)
- [Usage](#usage)
- [Examples](#examples)
- [Contributing](#contributing)
- [License](#license)

## Install

```shell
go get github.com/kevinpollet/nego
```

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
		contentCharset := nego.ContentCharset(req, "utf-8")
		contentEncoding := nego.ContentEncoding(req, "gzip", "deflate")
		contentLanguage := nego.ContentLanguage(req, "fr", "en")
		contentType := nego.ContentType(req, "text/html", "text/plain")
		...
	})
}
```

## Examples

The [examples](./examples) directory contains the following examples:

- [echo](./examples/echo) — This example returns the negotiated response content `type`, `charset`, `encoding` and `language`.
- [compress](./examples/compress) — This example negotiates the response content `encoding` and compresses the response body if the client supports the `gzip` encoding.

## Contributing

Contributions are welcome!

Want to file a bug, request a feature or contribute some code?

1. Check out the [Code of Conduct](./CODE_OF_CONDUCT.md).
2. Check for an existing issue corresponding to your bug or feature request.
3. Open an issue to describe your bug or feature request.

## License

[MIT](./LICENSE.md)
