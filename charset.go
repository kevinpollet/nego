package negotiate

import (
	"net/http"
)

// ContentCharset returns the best offered charset for the request's Accept-Charset header.
// If no offers are acceptable, then "" is returned.
func ContentCharset(req *http.Request, offerCharsets ...string) string {
	bestQvalue := 0.0
	bestCharset := ""
	acceptCharsets := parseAccept(req.Header, "Accept-Charset")

	if len(acceptCharsets) == 0 {
		return offerCharsets[0]
	}

	for _, charset := range offerCharsets {
		if qvalue, exists := acceptCharsets.qvalue(charset); exists && qvalue > bestQvalue {
			bestCharset = charset
			bestQvalue = qvalue
		}
	}

	return bestCharset
}
