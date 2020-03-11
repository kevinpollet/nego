package negotiate

import (
	"strconv"
	"strings"
)

// parseAccept parses the header value of Accept* HTTP headers.
func parseAccept(value string) (map[string]float64, error) {
	accepts := make(map[string]float64)
	specs := strings.Split(value, ",")

	if len(specs) == 1 && specs[0] == "" {
		return accepts, nil
	}

	for _, spec := range specs {
		qvalue := 1.0
		sSpec := strings.ReplaceAll(spec, " ", "")
		tokens := strings.Split(sSpec, ";")

		for _, param := range tokens[1:] {
			qvStr := strings.TrimLeft(strings.ToLower(param), "q=")
			if len(qvStr) != len(param) {
				qv, err := strconv.ParseFloat(qvStr, 64)
				if err != nil {
					return nil, err
				}

				qvalue = qv

				break
			}
		}

		accepts[tokens[0]] = qvalue
	}

	return accepts, nil
}
