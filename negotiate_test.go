package nego

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContentCharset(t *testing.T) { // nolint
	testCases := []struct {
		desc          string
		offerCharsets []string
		acceptCharset string
		expCharset    string
	}{
		{
			desc:          "first charset offer is acceptable if given request has no Accept-Charset header",
			offerCharsets: []string{"UTF-8", "UTF-16"},
			acceptCharset: "",
			expCharset:    "UTF-8",
		},
		{
			desc:          "no charset offer is acceptable if charset offers are not listed",
			offerCharsets: []string{"UTF-8", "UTF-16"},
			acceptCharset: "ISO-8859-5",
			expCharset:    "",
		},
		{
			desc:          "no charset offer is acceptable if charset offers are excluded",
			offerCharsets: []string{"UTF-8", "UTF-16"},
			acceptCharset: "UTF-8;q=0, UTF-16;q=0",
			expCharset:    "",
		},
		{
			desc:          "charset offer with best quality is acceptable if multiple charset offers are acceptable",
			offerCharsets: []string{"UTF-8", "UTF-16"},
			acceptCharset: "UTF-8;q=0.2, UTF-16;q=0.8",
			expCharset:    "UTF-16",
		},
	}

	for _, testCase := range testCases {
		testCase := testCase

		t.Run(testCase.desc, func(t *testing.T) {
			t.Parallel()

			req := httptest.NewRequest(http.MethodGet, "http://dummy.com", nil)

			if len(testCase.acceptCharset) > 0 {
				req.Header.Add("Accept-Charset", testCase.acceptCharset)
			}

			charset := ContentCharset(req, testCase.offerCharsets...)

			assert.Equal(t, testCase.expCharset, charset)
		})
	}
}

func TestContentEncoding(t *testing.T) {
	testCases := []struct {
		desc           string
		offerEncodings []string
		acceptEncoding string
		expEncoding    string
	}{
		{
			desc:           "identity is acceptable if request has no Accept-Encoding header",
			offerEncodings: []string{"gzip", "deflate"},
			acceptEncoding: "",
			expEncoding:    "identity",
		},
		{
			desc:           "identity is acceptable if the representation has no content-coding",
			offerEncodings: []string{},
			acceptEncoding: "gzip, deflate, br",
			expEncoding:    "identity",
		},
		{
			desc:           "identity is not acceptable if the representation has no content-coding and it is excluded",
			offerEncodings: []string{},
			acceptEncoding: "identity;q=0.0",
			expEncoding:    "",
		},
		{
			desc:           "identity is not acceptable if the representation has no content-coding and it is excluded",
			offerEncodings: []string{},
			acceptEncoding: "*;q=0.0",
			expEncoding:    "",
		},
		{
			desc:           "content-coding with highest qvalue is preferred if multiple content-codings are acceptable",
			offerEncodings: []string{"gzip", "deflate"},
			acceptEncoding: "gzip;q=0.8, deflate",
			expEncoding:    "deflate",
		},
	}

	for _, testCase := range testCases {
		testCase := testCase

		t.Run(testCase.desc, func(t *testing.T) {
			t.Parallel()

			req := httptest.NewRequest(http.MethodGet, "http://dummy.com", nil)
			req.Header.Add("Accept-Encoding", testCase.acceptEncoding)

			encoding := ContentEncoding(req, testCase.offerEncodings...)

			assert.Equal(t, testCase.expEncoding, encoding)
		})
	}
}

func TestContentLanguage(t *testing.T) { // nolint
	testCases := []struct {
		desc           string
		offerLanguages []string
		acceptLanguage string
		expLanguage    string
	}{
		{
			desc:           "first language offer is acceptable if given request has no Accept-Language header",
			offerLanguages: []string{"en", "en-US"},
			acceptLanguage: "",
			expLanguage:    "en",
		},
		{
			desc:           "no language offer is acceptable if language offers are not listed",
			offerLanguages: []string{"en", "en-US"},
			acceptLanguage: "fr",
			expLanguage:    "",
		},
		{
			desc:           "no language offer is acceptable if language offers are excluded",
			offerLanguages: []string{"en", "en-US"},
			acceptLanguage: "en;q=0, en-US;q=0",
			expLanguage:    "",
		},
		{
			desc:           "language offer with best quality is acceptable if multiple language offers are acceptable",
			offerLanguages: []string{"en", "en-US"},
			acceptLanguage: "en;q=0.2, en-US;q=0.8",
			expLanguage:    "en-US",
		},
	}

	for _, testCase := range testCases {
		testCase := testCase

		t.Run(testCase.desc, func(t *testing.T) {
			t.Parallel()

			req := httptest.NewRequest(http.MethodGet, "http://dummy.com", nil)

			if len(testCase.acceptLanguage) > 0 {
				req.Header.Add("Accept-Language", testCase.acceptLanguage)
			}

			language := ContentLanguage(req, testCase.offerLanguages...)

			assert.Equal(t, testCase.expLanguage, language)
		})
	}
}

func TestContentType(t *testing.T) {
	testCases := []struct {
		desc        string
		offerTypes  []string
		acceptType  string
		expLanguage string
	}{
		{
			desc:        "first media type offer is acceptable if given request has no Accept header",
			offerTypes:  []string{"text/html"},
			acceptType:  "",
			expLanguage: "text/html",
		},
		{
			desc:        "media type offer with the most specific reference has precedence",
			offerTypes:  []string{"text/html", "text/plain"},
			acceptType:  "text/html;q=0.2, text/*",
			expLanguage: "text/plain",
		},
		{
			desc:        "no media type offer is acceptable if they are not listed",
			offerTypes:  []string{"image/jpg"},
			acceptType:  "text/html;q=0.2, text/*",
			expLanguage: "",
		},
		{
			desc:        "no media type offer is acceptable if they are explicitly excluded",
			offerTypes:  []string{"image/jpg"},
			acceptType:  "image/jpg;q=0.0, text/*",
			expLanguage: "",
		},
	}

	for _, testCase := range testCases {
		testCase := testCase

		t.Run(testCase.desc, func(t *testing.T) {
			t.Parallel()

			req := httptest.NewRequest(http.MethodGet, "http://dummy.com", nil)

			if len(testCase.acceptType) > 0 {
				req.Header.Add("Accept", testCase.acceptType)
			}

			mediaType := ContentType(req, testCase.offerTypes...)

			assert.Equal(t, testCase.expLanguage, mediaType)
		})
	}
}
