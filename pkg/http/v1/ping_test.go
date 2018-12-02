package v1

import (
	"github.com/gin-gonic/gin"
	"net/http/httptest"
	"testing"
)

func TestPing(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(r)

	Ping(c)

	if r.Body.String() != "{\"message\":\"pong\"}" {
		t.Error("Message body is different")
	}
}
