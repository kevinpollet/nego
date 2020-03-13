package negotiate

import (
	"net/http"
)

// ContentCharset returns the best offered charset for the request's Accept-Charset header.
// If no offers are acceptable, then "" is returned.
func ContentCharset(req *http.Request, offerCharsets ...string) string {
	bestQvalue := 0.0
	bestCharset := ""

	acceptCharsets, exists := parseAccept(req.Header, "Accept-Charset")
	if !exists && len(offerCharsets) > 0 {
		return offerCharsets[0]
	}

	for _, offer := range offerCharsets {
		if qvalue, exists := acceptCharsets.qvalue(offer); exists && qvalue > bestQvalue {
			bestCharset = offer
			bestQvalue = qvalue
		}
	}

	return bestCharset
}
