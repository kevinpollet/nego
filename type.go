package nego

import (
	"net/http"
)

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
