package tests

import (
	"io"
	"log"
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
)

func SetupTestRoutes() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	return gin.Default()
}

type Request struct {
	method string
	path   string
	body   io.Reader
}

func MockServer(params *Request, controller gin.HandlerFunc) *httptest.ResponseRecorder {
	r := SetupTestRoutes()

	if params.method == http.MethodPost {
		r.POST(params.path, controller)
	} else {
		r.GET(params.path, controller)
	}

	return MockReq(r, params)
}

func MockReq(route *gin.Engine, r *Request) *httptest.ResponseRecorder {
	req, err := http.NewRequest(r.method, r.path, r.body)
	if err != nil {
		log.Panic("this is the error:", err)
	}
	res := httptest.NewRecorder()
	route.ServeHTTP(res, req)

	return res
}
