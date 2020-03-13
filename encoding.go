package negotiate

import "net/http"

const encodingIdentity = "identity"

// ContentEncoding returns the best offered charset for the request's Accept-Encoding header.
// If no offers are acceptable, then identity encoding is returned or "" if it is explicitly excluded.
func ContentEncoding(req *http.Request, offerEncodings ...string) string {
	bestQvalue := 0.0
	bestEncoding := ""

	acceptEncodings, exists := parseAccept(req.Header, "Accept-Encoding")
	if !exists && len(offerEncodings) > 0 {
		return offerEncodings[0]
	}

	for _, offer := range offerEncodings {
		if qvalue, exists := acceptEncodings.qvalue(offer); exists && qvalue > bestQvalue {
			bestEncoding = offer
			bestQvalue = qvalue
		}
	}

	if qvalue, exists := acceptEncodings.qvalue(encodingIdentity); bestEncoding == "" && (!exists || qvalue > 0.0) {
		return encodingIdentity
	}

	return bestEncoding
}
