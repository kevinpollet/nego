package negotiate // nolint

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLanguage(t *testing.T) {
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
			req.Header.Add("Accept-Language", testCase.acceptLanguage)

			language := Language(req, testCase.offerLanguages...)

			assert.Equal(t, testCase.expLanguage, language)
		})
	}
}
