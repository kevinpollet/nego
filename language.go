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

func ContentLanguage(req *http.Request, offerLanguages ...string) string {
	bestQvalue := 0.0
	bestLanguage := ""
	acceptLanguages := acceptLanguage(parseAccept(req.Header, "Accept-Language"))

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
