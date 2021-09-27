package nego

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNegotiateContentCharset(t *testing.T) {
	testCases := []struct {
		desc       string
		offers     []string
		accept     string
		expCharset string
	}{
		{
			desc:       "should return the first offer if the request has no Accept-Charset header",
			offers:     []string{"utf-8"},
			accept:     "",
			expCharset: "utf-8",
		},
		{
			desc:       "should return an empty string if no offer is acceptable",
			offers:     []string{"utf-8"},
			accept:     "utf-16",
			expCharset: "",
		},
		{
			desc:       "should return an empty string if the offer is explicitly discared",
			offers:     []string{"utf-8"},
			accept:     "utf-8;q=0",
			expCharset: "",
		},
		{
			desc:       "should return an empty string if no offer is defined",
			offers:     []string{},
			accept:     "utf-8",
			expCharset: "",
		},
		{
			desc:       "should return the acceptable offer with the best qvalue if multiple offers are acceptable",
			offers:     []string{"utf-8", "utf-16"},
			accept:     "utf-8;q=0.8, utf-16",
			expCharset: "utf-16",
		},
	}

	for _, testCase := range testCases {
		testCase := testCase

		t.Run(testCase.desc, func(t *testing.T) {
			t.Parallel()

			req := httptest.NewRequest(http.MethodGet, "http://dummy.com", nil)
			if testCase.accept != "" {
				req.Header.Add("Accept-Charset", testCase.accept)
			}

			contentCharset := NegotiateContentCharset(req, testCase.offers...)

			assert.Equal(t, testCase.expCharset, contentCharset)
		})
	}
}

func TestNegotiateContentEncoding(t *testing.T) {
	testCases := []struct {
		desc        string
		offers      []string
		accept      string
		expEncoding string
	}{
		{
			desc:        "should return the first offer if the request has no Accept-Encoding header",
			offers:      []string{"gzip", "deflate"},
			accept:      "",
			expEncoding: "gzip",
		},
		{
			desc:        "should return identity if no offer is defined",
			offers:      []string{},
			accept:      "",
			expEncoding: "identity",
		},
		{
			desc:        "should return identity if no offer is acceptable",
			offers:      []string{"gzip", "br"},
			accept:      "deflate",
			expEncoding: "identity",
		},
		{
			desc:        "should return identity if the request has an empty Accept-Encoding header",
			offers:      []string{"gzip", "br"},
			accept:      " ",
			expEncoding: "identity",
		},
		{
			desc:        "should return an empty string if no offer is defined and identity is explicitly discared",
			offers:      []string{},
			accept:      "identity;q=0",
			expEncoding: "",
		},
		{
			desc:        "should return an empty string if no offer is defined and identity is discared",
			offers:      []string{},
			accept:      "*;q=0",
			expEncoding: "",
		},
		{
			desc:        "should return the acceptable offer",
			offers:      []string{"gzip", "br"},
			accept:      "br",
			expEncoding: "br",
		},
		{
			desc:        "should return the acceptable offer with the best qvalue if multiple offers are acceptable",
			offers:      []string{"gzip", "br"},
			accept:      "br;q=0.5, gzip",
			expEncoding: "gzip",
		},
	}

	for _, testCase := range testCases {
		testCase := testCase

		t.Run(testCase.desc, func(t *testing.T) {
			t.Parallel()

			req := httptest.NewRequest(http.MethodGet, "http://dummy.com", nil)
			if testCase.accept != "" {
				req.Header.Add("Accept-Encoding", testCase.accept)
			}

			contentEncoding := NegotiateContentEncoding(req, testCase.offers...)

			assert.Equal(t, testCase.expEncoding, contentEncoding)
		})
	}
}

func TestNegotiateContentLanguage(t *testing.T) {
	testCases := []struct {
		desc        string
		offers      []string
		accept      string
		expLanguage string
	}{
		{
			desc:        "should return the first offer if request has no Accept-Language header",
			offers:      []string{"en", "en-us"},
			accept:      "",
			expLanguage: "en",
		},
		{
			desc:        "should return an empty string if no offers is acceptable",
			offers:      []string{"en", "en-us"},
			accept:      "fr",
			expLanguage: "",
		},
		{
			desc:        "should return an empty string if an offer is explicitly discared",
			offers:      []string{"en", "en-us"},
			accept:      "en;q=0",
			expLanguage: "",
		},
		{
			desc:        "should return an empty string if no offer is defined",
			offers:      []string{},
			accept:      "en",
			expLanguage: "",
		},
		{
			desc:        "should return the acceptable offer with the best qvalue if multiple offers are acceptable",
			offers:      []string{"en", "en-us"},
			accept:      "en;q=0.8, en-us",
			expLanguage: "en-us",
		},
	}

	for _, testCase := range testCases {
		testCase := testCase

		t.Run(testCase.desc, func(t *testing.T) {
			t.Parallel()

			req := httptest.NewRequest(http.MethodGet, "http://dummy.com", nil)
			if testCase.accept != "" {
				req.Header.Add("Accept-Language", testCase.accept)
			}

			contentLanguage := NegotiateContentLanguage(req, testCase.offers...)

			assert.Equal(t, testCase.expLanguage, contentLanguage)
		})
	}
}

func TestNegotiateContentType(t *testing.T) {
	testCases := []struct {
		desc         string
		offers       []string
		accept       string
		expMediaType string
	}{
		{
			desc:         "should return the first offer if the request has no Accept header",
			offers:       []string{"text/html"},
			accept:       "",
			expMediaType: "text/html",
		},
		{
			desc:         "should return an empty string if no offer is acceptable",
			offers:       []string{"text/html", "text/plain"},
			accept:       "application/json",
			expMediaType: "",
		},
		{
			desc:         "should return an empty string if no offer is acceptable",
			offers:       []string{"text/html", "text/plain"},
			accept:       "application/json",
			expMediaType: "",
		},
		{
			desc:         "should return an empty string if offer is explicitly discared",
			offers:       []string{"text/html"},
			accept:       "text/html;q=0",
			expMediaType: "",
		},
		{
			desc:         "should return an empty string if offer is discared",
			offers:       []string{"text/html"},
			accept:       "text/*;q=0",
			expMediaType: "",
		},
		{
			desc:         "should return an empty string if no offer is defined",
			offers:       []string{},
			accept:       "application/json",
			expMediaType: "",
		},
		{
			desc:         "should return the acceptable offer with the best qvalue if multiple offers are acceptable",
			offers:       []string{"text/html", "text/plain"},
			accept:       "text/html;q=0.8, text/plain",
			expMediaType: "text/plain",
		},
		// {
		// 	desc:         "should return the most specific acceptable offer if multiple offers are acceptable",
		// 	offers:       []string{"text/html", "text/plain"},
		// 	accept:       "text/plain, text/*",
		// 	expMediaType: "text/plain",
		// },
	}

	for _, testCase := range testCases {
		testCase := testCase

		t.Run(testCase.desc, func(t *testing.T) {
			t.Parallel()

			req := httptest.NewRequest(http.MethodGet, "http://dummy.com", nil)
			if testCase.accept != "" {
				req.Header.Add("Accept", testCase.accept)
			}

			contentType := NegotiateContentType(req, testCase.offers...)

			assert.Equal(t, testCase.expMediaType, contentType)
		})
	}
}
