package negotiate

import "net/http"

type acceptEncoding map[string]float64

func (a acceptEncoding) qvalue(encoding string) (qvalue float64, exists bool) {
	if qvalue, exists = a[encoding]; !exists {
		qvalue, exists = a["*"]
	}

	return
}

// ContentEncoding returns the best offered charset for the request's Accept-Encoding header.
// If no offers are acceptable, then identity encoding is returned or "" if it is explicitly excluded.
func ContentEncoding(req *http.Request, offerEncodings ...string) string {
	bestQvalue := 0.0
	bestEncoding := ""
	identity := "identity"
	acceptEncodings := acceptEncoding(parseAccept(req.Header, "Accept-Encoding"))

	if len(acceptEncodings) == 0 {
		return identity
	}

	for _, encoding := range offerEncodings {
		if qvalue, exists := acceptEncodings.qvalue(encoding); exists && qvalue > bestQvalue {
			bestEncoding = encoding
			bestQvalue = qvalue
		}
	}

	if qvalue, exists := acceptEncodings.qvalue(identity); bestEncoding == "" && (!exists || qvalue > 0.0) {
		return identity
	}

	return bestEncoding
}
