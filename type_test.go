package nego // nolint

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

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
