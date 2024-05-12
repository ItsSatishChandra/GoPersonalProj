package hello

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestHello(t *testing.T) {

	type testCase struct {
		method   string
		url      string
		mimetype string
	}
	testCases := []testCase{
		{method: echo.GET, url: "http://localhost:8000/", mimetype: echo.MIMEApplicationJSON}}
	for _, test := range testCases {
		req := httptest.NewRequest(test.method, test.url, nil)
		req.Header.Set(echo.HeaderContentType, test.mimetype)

		rec := httptest.NewRecorder()
		context := echo.New().NewContext(req, rec)

		if assert.NoError(t, Hello(context)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.Equal(t, "hello world", rec.Body.String())
		}

	}

}
