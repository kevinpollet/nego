package negotiate

import (
	"net/http"
)

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
