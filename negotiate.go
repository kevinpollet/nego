package negotiate

import "net/http"

type acceptEncoding map[string]float64

func (a acceptEncoding) qvalue(contentCoding string) (qvalue float64, exists bool) {
	if qvalue, exists = a[contentCoding]; !exists {
		qvalue, exists = a["*"]
	}

	return
}

func ContentEncoding(req *http.Request, offerEncodings ...string) string {
	bestQvalue := 0.0
	bestEncoding := ""
	identity := "identity"

	values := req.Header["Accept-Encoding"]
	acceptEncodings := acceptEncoding(parseContentNegotiation(values))

	if len(acceptEncodings) == 0 {
		return identity
	}

	for _, contentCoding := range offerEncodings {
		if qvalue, exists := acceptEncodings.qvalue(contentCoding); exists && qvalue > bestQvalue {
			bestEncoding = contentCoding
			bestQvalue = qvalue
		}
	}

	if qvalue, exists := acceptEncodings.qvalue(identity); bestEncoding == "" && (!exists || qvalue > 0.0) {
		return identity
	}

	return bestEncoding
}
