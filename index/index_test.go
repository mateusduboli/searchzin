package index

import (
	"github.com/mateusduboli/searchzin/document"
	"testing"
)

func TestIndexDocument(t *testing.T) {
	document := NewDocument(map[string]interface{}{
		"id":   4,
		"name": "joão",
	})

	IndexDocument(document)
}
