package private

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/mateusduboli/searchzin/pkg/document"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestDocumentList(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(r)

	DocumentList(c)

	decoder := json.NewDecoder(r.Body)

	var actual []map[string]interface{}

	decoder.Decode(&actual)

	if len(actual) != 0 {
		t.Errorf("Actual was not empty: [%s]", actual)
	}
}

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

	DocumentSave(c)

	decoder := json.NewDecoder(r.Body)

	var actual document.Document

	decoder.Decode(&actual)

	if actual.Data["Test"] != "test" {
		t.Error("Failed to receive message")
	}
}
