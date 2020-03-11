package negotiate

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContentEncoding(t *testing.T) {
	testCases := []struct {
		desc            string
		offers          []string
		acceptEncoding  string
		contentEncoding string
	}{
		{
			desc:            "first offer is acceptable if no Accept-Encoding header is in the request",
			offers:          []string{"gzip", "deflate"},
			acceptEncoding:  "",
			contentEncoding: "gzip",
		},
		{
			desc:            "identity is acceptable if the representation has no content-coding",
			offers:          []string{},
			acceptEncoding:  "gzip, deflate, br",
			contentEncoding: "identity",
		},
		{
			desc:            "identity is not acceptable if identity content-coding is explicitly disallowed",
			offers:          []string{},
			acceptEncoding:  "identity;q=0.0",
			contentEncoding: "",
		},
		{
			desc:            "identity is not acceptable if identity content-coding is explicitly disallowed",
			offers:          []string{},
			acceptEncoding:  "*;q=0.0",
			contentEncoding: "",
		},
		{
			desc:            "offer with the highest non-zero qvalue is preferred if multiple content-codings are acceptable",
			offers:          []string{"gzip", "deflate"},
			acceptEncoding:  "gzip;q=0.8, deflate",
			contentEncoding: "deflate",
		},
	}

	for _, testCase := range testCases {
		testCase := testCase

		t.Run(testCase.desc, func(t *testing.T) {
			t.Parallel()

			req := httptest.NewRequest(http.MethodGet, "http://dummy.com", nil)
			req.Header.Add("Accept-Encoding", testCase.acceptEncoding)

			contentEncoding, err := ContentEncoding(req, testCase.offers...)

			assert.NoError(t, err)
			assert.Equal(t, testCase.contentEncoding, contentEncoding)
		})
	}
}
