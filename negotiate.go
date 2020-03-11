package negotiate

import "net/http"

func ContentEncoding(req *http.Request, offers ...string) (string, error) {
	acceptEncoding := req.Header.Get("Accept-Encoding")

	acceptSpecs, err := parseAccept(acceptEncoding)
	if err != nil {
		return "", err
	}

	if len(acceptSpecs) == 0 {
		return offers[0], nil
	}

	bestOffer := ""
	bestQValue := 0.0

	for _, offer := range offers {
		qvalue, exists := acceptSpecs[offer]
		if !exists {
			qvalue, exists = acceptSpecs["*"]
		}

		if exists && qvalue > bestQValue {
			bestOffer = offer
			bestQValue = qvalue
		}
	}

	if bestOffer == "" {
		qvalue, exists := acceptSpecs["identity"]
		if !exists {
			qvalue, exists = acceptSpecs["*"]
		}

		if !exists || qvalue > 0.0 {
			bestOffer = "identity"
			bestQValue = qvalue // nolint
		}
	}

	return bestOffer, nil
}
