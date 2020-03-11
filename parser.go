package negotiate

import (
	"strconv"
	"strings"
)

// parseContentNegotiation parses and the value of a content negotiation header.
// The following request headers are sent by a user agent to engage in proactive negotiation
// of the response content: Accept, Accept-Charset, Accept-Encoding, Accept-Language.
func parseContentNegotiation(values []string) map[string]float64 {
	accepts := make(map[string]float64)

	for _, value := range values {
		specs := strings.Split(value, ",")

		for _, spec := range specs {
			name, qvalue := parseSpec(spec)
			accepts[name] = qvalue
		}
	}

	return accepts
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
