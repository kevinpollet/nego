package negotiate

import (
	"net/http"
)

type acceptLanguage map[string]float64

func (a acceptLanguage) qvalue(language string) (qvalue float64, exists bool) {
	if qvalue, exists = a[language]; !exists {
		qvalue, exists = a["*"]
	}

	return
}

func Language(req *http.Request, offerLanguages ...string) string {
	bestQvalue := 0.0
	bestLanguage := ""

	values := req.Header["Accept-Language"]
	acceptCharsets := acceptLanguage(parseContentNegotiation(values))

	if len(acceptCharsets) == 0 {
		return offerLanguages[0]
	}

	for _, language := range offerLanguages {
		if qvalue, exists := acceptCharsets.qvalue(language); exists && qvalue > bestQvalue {
			bestLanguage = language
			bestQvalue = qvalue
		}
	}

	return bestLanguage
}
