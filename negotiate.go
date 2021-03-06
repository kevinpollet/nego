// Package nego implements HTTP Content Negotiation functions compliant with RFC 7231.
//
// See https://tools.ietf.org/html/rfc7231#section-5.3 for more details.
//
// Example
//
// This example shows how to use the negotiation functions.
//
//	import (
//		"net/http"
//		"github.com/kevinpollet/nego"
//	)
//
// 	handler := http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
// 		nego.NegotiateContentCharset(req, "utf-8")
// 		nego.NegotiateContentEncoding(req, "gzip", "deflate")
// 		nego.NegotiateContentLanguage(req, "fr", "en")
// 		nego.NegotiateContentType(req, "text/plain")
// 	})
package nego

import "net/http"

// The identity encoding constant used as a synonym for "no encoding" in order to communicate when
// no encoding is preferred.
//
// See https://tools.ietf.org/html/rfc7231#section-5.3.4 for more details.
const EncodingIdentity = "identity"

// NegotiateContentCharset returns the best acceptable charset offer to use in the response according
// to the Accept-Charset request's header. If the given offer list is empty or no offer is acceptable
// then, an empty string is returned.
//
// See https://tools.ietf.org/html/rfc7231#section-5.3.3 for more details.
func NegotiateContentCharset(req *http.Request, offerCharsets ...string) string {
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

// NegotiateContentEncoding returns the best acceptable encoding offer to use in the response according
// to the Accept-Encoding request's header. If the given offer list is empty or no offer is acceptable
// then, an empty string is returned.
//
// See https://tools.ietf.org/html/rfc7231#section-5.3.4 for more details.
func NegotiateContentEncoding(req *http.Request, offerEncodings ...string) string {
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

	if qvalue, exists := acceptEncodings.qvalue(EncodingIdentity); bestEncoding == "" && (!exists || qvalue > 0.0) {
		return EncodingIdentity
	}

	return bestEncoding
}

// NegotiateContentLanguage returns the best acceptable language offer to use in the response according
// to the Accept-Language request's header. If the given offer list is empty or no offer is acceptable
// then, an empty string is returned.
//
// See https://tools.ietf.org/html/rfc7231#section-5.3.5 for more details.
func NegotiateContentLanguage(req *http.Request, offerLanguages ...string) string {
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

// NegotiateContentType returns the best acceptable media type offer to use in the response according
// to the Accept-Language request's header. If the given offer list is empty or no offer is acceptable
// then, an empty string is returned.
//
// See https://tools.ietf.org/html/rfc7231#section-5.3.2 for more details.
func NegotiateContentType(req *http.Request, offerMediaTypes ...string) string {
	bestMediaType := ""
	bestQvalue := 0.0

	acceptTypes, exists := parseAccept(req.Header, "Accept")
	if !exists && len(offerMediaTypes) > 0 {
		return offerMediaTypes[0]
	}

	for _, offer := range offerMediaTypes {
		if qvalue, exists := acceptTypes.qvalue(offer); exists && qvalue > bestQvalue {
			bestMediaType = offer
			bestQvalue = qvalue
		}
	}

	return bestMediaType
}
