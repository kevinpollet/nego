package nego

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseAccept(t *testing.T) {
	testCases := []struct {
		desc   string
		accept string
		expLen int
	}{
		{
			desc:   "should return an empty map if the given values are empty",
			accept: "",
			expLen: 0,
		},
		{
			desc:   "should return a map with one element",
			accept: "gzip",
			expLen: 1,
		},
		{
			desc:   "should return a map with the given number of elements",
			accept: "gzip,deflate",
			expLen: 2,
		},
		{
			desc:   "should return a map with the given number of elements ignoring spaces",
			accept: "gzip , deflate",
			expLen: 2,
		},
		{
			desc:   "should return a map with the given number of elements in the given values",
			accept: "gzip, deflate, br",
			expLen: 3,
		},
	}

	for _, testCase := range testCases {
		testCase := testCase

		t.Run(testCase.desc, func(t *testing.T) {
			t.Parallel()

			header := make(http.Header)
			header.Add("Accept", testCase.accept)

			specs, _ := parseAccept(header, "Accept")

			assert.Equal(t, testCase.expLen, len(specs))
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
			expQ:  1.0,
		},
		{
			desc:  "should return the parsed name with the given quality",
			value: "test;q=0.1",
			expN:  "test",
			expQ:  0.1,
		},
		{
			desc:  "should return the parsed name with the given quality ignoring whitespaces",
			value: "test ; q=0.1",
			expN:  "test",
			expQ:  0.1,
		},
		{
			desc:  "should return the parsed name with the given quality ignoring extra params",
			value: "test ; format=foo; q=0.1; format=bar",
			expN:  "test",
			expQ:  0.1,
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
			expQ:  1.0,
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
