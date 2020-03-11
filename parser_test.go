package negotiate

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContentNegotiation(t *testing.T) {
	testCases := []struct {
		desc   string
		accept string
		expL   int
	}{
		{
			desc:   "should return an empty map if the given value is empty",
			accept: "",
			expL:   0,
		},
		{
			desc:   "should return a map with one element",
			accept: "gzip",
			expL:   1, //nolint
		},
		{
			desc:   "should return a map with the given number of elements",
			accept: "gzip,deflate",
			expL:   2, //nolint
		},
		{
			desc:   "should return a map with the given number of elements ignoring spaces",
			accept: "gzip , deflate",
			expL:   2, //nolint
		},
	}

	for _, testCase := range testCases {
		testCase := testCase

		t.Run(testCase.desc, func(t *testing.T) {
			t.Parallel()

			specs := parseContentNegotiation(testCase.accept)

			assert.Equal(t, testCase.expL, len(specs))
		})
	}
}

func TestParseSpec(t *testing.T) {
	testCases := []struct {
		desc  string
		value string
		expN  string
		expQ  float64
	}{
		{
			desc:  "should return the parsed name with the default quality",
			value: "test",
			expN:  "test",
			expQ:  1.0, // nolint
		},
		{
			desc:  "should return the parsed name with the given quality",
			value: "test;q=0.1",
			expN:  "test",
			expQ:  0.1, // nolint
		},
		{
			desc:  "should return the parsed name with the given quality ignoring whitespaces",
			value: "test ; q=0.1",
			expN:  "test",
			expQ:  0.1, // nolint
		},
		{
			desc:  "should return the parsed name with the given quality ignoring extra params",
			value: "test ; format=foo; q=0.1; format=bar",
			expN:  "test",
			expQ:  0.1, // nolint
		},
	}

	for _, testCase := range testCases {
		testCase := testCase

		t.Run(testCase.desc, func(t *testing.T) {
			t.Parallel()

			name, quality := parseSpec(testCase.value)

			assert.Equal(t, testCase.expN, name)
			assert.Equal(t, testCase.expQ, quality)
		})
	}
}

func TestParseQuality(t *testing.T) {
	testCases := []struct {
		desc  string
		value string
		expQ  float64
	}{
		{
			desc:  "should return the parsed value",
			value: "1.0",
			expQ:  1.0, // nolint
		},
		{
			desc:  "should return -1 if the value cannot be parsed",
			value: "aa",
			expQ:  -1.0,
		},
	}

	for _, testCase := range testCases {
		testCase := testCase

		t.Run(testCase.desc, func(t *testing.T) {
			t.Parallel()

			quality := parseQuality(testCase.value)

			assert.Equal(t, testCase.expQ, quality)
		})
	}
}
