package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"testing"
)

func request(path string) *httptest.ResponseRecorder {
	r := gin.Default()
	req, _ := http.NewRequest("GET", path, nil)
	ping(r)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	return w
}

func TestPing(t *testing.T) {
	res := request("/ping")

	if res.Body.String() != "{\"message\":\"pong\"}" {
		t.Error("Message body is different")
	}
}
