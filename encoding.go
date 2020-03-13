package negotiate

import "net/http"

const encodingIdentity = "identity"

// ContentEncoding returns the best offered charset for the request's Accept-Encoding header.
// If no offers are acceptable, then identity encoding is returned or "" if it is explicitly excluded.
func ContentEncoding(req *http.Request, offerEncodings ...string) string {
	bestQvalue := 0.0
	bestEncoding := ""
	acceptEncodings := parseAccept(req.Header, "Accept-Encoding")

	if len(acceptEncodings) == 0 {
		return encodingIdentity
	}

	for _, encoding := range offerEncodings {
		if qvalue, exists := acceptEncodings.qvalue(encoding); exists && qvalue > bestQvalue {
			bestEncoding = encoding
			bestQvalue = qvalue
		}
	}

	if qvalue, exists := acceptEncodings.qvalue(encodingIdentity); bestEncoding == "" && (!exists || qvalue > 0.0) {
		return encodingIdentity
	}

	return bestEncoding
}
