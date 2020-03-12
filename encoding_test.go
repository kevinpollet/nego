package negotiate

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

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
