package negotiate

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseAccept(t *testing.T) {
	testCases := []struct {
		desc           string
		accept         string
		expectedResult map[string]float64
	}{
		{
			desc:           "return an empty map if given value is empty",
			accept:         "",
			expectedResult: make(map[string]float64),
		},
		{
			desc:           "return a map with one element and qvalue initialized to 1.0 by default",
			accept:         "gzip",
			expectedResult: map[string]float64{"gzip": 1.0}, //nolint
		},
		{
			desc:           "return a map with all element qvalues initialized to 1.0 by default",
			accept:         "gzip, deflate, *",
			expectedResult: map[string]float64{"gzip": 1.0, "deflate": 1.0, "*": 1.0}, //nolint
		},
		{
			desc:           "return a map with with one element and qvalue initialized to the specified qvalue",
			accept:         "gzip;q=0.2",
			expectedResult: map[string]float64{"gzip": 0.2}, //nolint
		},
		{
			desc:           "return a map with with all element qvalues initialized to the specified qvalue",
			accept:         "gzip;q=0.2, *;q=0.0",
			expectedResult: map[string]float64{"gzip": 0.2, "*": 0.0}, //nolint
		},
		{
			desc:           "return a map with with all element qvalues initialized to the specified qvalue ignoring params",
			accept:         "text/*, text/plain;format=flowed;q=0.8",
			expectedResult: map[string]float64{"text/*": 1.0, "text/plain": 0.8}, //nolint
		},
	}

	for _, testCase := range testCases {
		testCase := testCase

		t.Run(testCase.desc, func(t *testing.T) {
			t.Parallel()

			specs, err := parseAccept(testCase.accept)

			assert.NoError(t, err)
			assert.Equal(t, testCase.expectedResult, specs)
		})
	}
}
