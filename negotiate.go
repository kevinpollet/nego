package nego

import "net/http"

const encodingIdentity = "identity"

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

// ContentLanguage returns the best offered language for the request's Accept-Language header.
// If no offers are acceptable, then "" is returned.
func ContentLanguage(req *http.Request, offerLanguages ...string) string {
	bestQvalue := 0.0
	bestLanguage := ""

	acceptLanguages, exists := parseAccept(req.Header, "Accept-Language")
	if !exists && len(offerLanguages) > 0 {
		return offerLanguages[0]
	}

	for _, offer := range offerLanguages {
		if qvalue, exists := acceptLanguages.qvalue(offer); exists && qvalue > bestQvalue {
			bestLanguage = offer
			bestQvalue = qvalue
		}
	}

	return bestLanguage
}

// ContentType returns the best offered language for the request's Accept header.
// If no offers are acceptable, then "" is returned.
func ContentType(req *http.Request, offerTypes ...string) string {
	bestType := ""
	bestQvalue := 0.0

	acceptTypes, exists := parseAccept(req.Header, "Accept")
	if !exists && len(offerTypes) > 0 {
		return offerTypes[0]
	}

	for _, offer := range offerTypes {
		if qvalue, exists := acceptTypes.qvalue(offer); exists && qvalue > bestQvalue {
			bestType = offer
			bestQvalue = qvalue
		}
	}

	return bestType
}
