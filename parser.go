package nego

import (
	"net/http"
	"strconv"
	"strings"
)

type accept map[string]float64

func (a accept) qvalue(offer string) (float64, bool) {
	if qvalue, exists := a[offer]; exists {
		return qvalue, exists
	}

	if !strings.Contains(offer, "/") {
		qvalue, exists := a["*"]
		return qvalue, exists
	}

	slashIndex := strings.Index(offer, "/")

	if qvalue, exists := a[offer[:slashIndex]+"/*"]; exists {
		return qvalue, exists
	}

	if qvalue, exists := a["*/*"]; exists {
		return qvalue, exists
	}

	return 0.0, false // nolint
}

// parseAccept parses the values of a content negotiation header. The following request headers are sent
// by a user agent to engage in proactive negotiation: Accept, Accept-Charset, Accept-Encoding, Accept-Language.
func parseAccept(header http.Header, key string) (accept, bool) {
	values, exists := header[key]
	accepts := make(map[string]float64)

	for _, value := range values {
		if len(value) == 0 {
			continue
		}

		for _, spec := range strings.Split(value, ",") {
			name, qvalue := parseSpec(spec)
			accepts[name] = qvalue
		}
	}

	return accepts, exists
}

func parseSpec(spec string) (string, float64) {
	qvalue := 1.0
	sToken := strings.ReplaceAll(spec, " ", "")
	parts := strings.Split(sToken, ";")

	for _, param := range parts[1:] {
		lowerParam := strings.ToLower(param)
		qvalueStr := strings.TrimPrefix(lowerParam, "q=")

		if qvalueStr != lowerParam {
			qvalue = parseQuality(qvalueStr)
		}
	}

	return parts[0], qvalue
}

func parseQuality(value string) float64 {
	float, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return -1
	}

	return float
}
