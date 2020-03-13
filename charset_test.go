package negotiate // nolint

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCharset(t *testing.T) {
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
			req.Header.Add("Accept-Charset", testCase.acceptCharset)

			charset := Charset(req, testCase.offerCharsets...)

			assert.Equal(t, testCase.expCharset, charset)
		})
	}
}
