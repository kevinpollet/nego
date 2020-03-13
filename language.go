package negotiate

import (
	"net/http"
)

// ContentLanguage returns the best offered language for the request's Accept-Language header.
// If no offers are acceptable, then "" is returned.
func ContentLanguage(req *http.Request, offerLanguages ...string) string {
	bestQvalue := 0.0
	bestLanguage := ""
	acceptLanguages := parseAccept(req.Header, "Accept-Language")

	if len(acceptLanguages) == 0 {
		return offerLanguages[0]
	}

	for _, language := range offerLanguages {
		if qvalue, exists := acceptLanguages.qvalue(language); exists && qvalue > bestQvalue {
			bestLanguage = language
			bestQvalue = qvalue
		}
	}

	return bestLanguage
}
