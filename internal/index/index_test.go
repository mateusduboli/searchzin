package index

import (
	"github.com/mateusduboli/searchzin/internal/document"
	"testing"
)

func TestIndexDocument(t *testing.T) {
	document := document.NewDocument(map[string]interface{}{

		"id":   4,
		"name": "joão",
	})

	IndexDocument(document)
}
