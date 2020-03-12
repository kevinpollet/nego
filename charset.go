package negotiate

import (
	"net/http"
)

type acceptCharset map[string]float64

func (a acceptCharset) qvalue(charset string) (qvalue float64, exists bool) {
	if qvalue, exists = a[charset]; !exists {
		qvalue, exists = a["*"]
	}

	return
}

func Charset(req *http.Request, offerCharsets ...string) string {
	bestQvalue := 0.0
	bestCharset := ""

	values := req.Header["Accept-Charset"]
	acceptCharsets := acceptCharset(parseContentNegotiation(values))

	if len(acceptCharsets) == 0 {
		return offerCharsets[0]
	}

	for _, charset := range offerCharsets {
		if qvalue, exists := acceptCharsets.qvalue(charset); exists && qvalue > bestQvalue {
			bestCharset = charset
			bestQvalue = qvalue
		}
	}

	return bestCharset
}