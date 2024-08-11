package mocks

import (
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"

	"github.com/BrunoPolaski/go-crud/src/configuration/logger"
	"github.com/gin-gonic/gin"
)

func GetTestGinContext(recorder *httptest.ResponseRecorder) *gin.Context {
	gin.SetMode(gin.TestMode)

	ctx, _ := gin.CreateTestContext(recorder)
	ctx.Request = &http.Request{
		Header: make(http.Header),
		URL:    &url.URL{},
	}

	return ctx
}

type BasicAuth struct {
	Username string
	Password string
}

func MakeRequest(
	c *gin.Context,
	params gin.Params,
	u url.Values,
	method string,
	body io.ReadCloser,
	basicAuth *BasicAuth,
) {
	c.Request.Method = method
	c.Request.Header.Set("Content-Type", "application/json")
	c.Params = params
	c.Request.URL.RawQuery = u.Encode()
	c.Request.Body = body
	if basicAuth != nil {
		c.Request.SetBasicAuth(basicAuth.Username, basicAuth.Password)
		logger.Info("Basic auth set")
	}
}
