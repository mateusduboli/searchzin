package api

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestDocumentSave(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(r)

	sample := map[string]interface{}{
		"Test": "test",
	}

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(sample)

	c.Request, _ = http.NewRequest("POST", "/", b)

	documentsSave(c)

	decoder := json.NewDecoder(r.Body)

	var actual map[string]interface{}

	decoder.Decode(&actual)

	if actual["Test"] != "test" {
		t.Error("Failed to receive message")
	}
}
