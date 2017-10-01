package indexer

import (
	"github.com/satori/go.uuid"
)

type Document struct {
	id   string
	data map[string]interface{}
}

func generateId() string {
	return uuid.NewV4().String()
}

func deepCopy(document map[string]interface{}) map[string]interface{} {
	d := make(map[string]interface{})
	for k := range document {
		d[k] = document[k]
	}
	return d
}

func NewDocument(data map[string]interface{}) Document {
	return Document{
		id:   generateId(),
		data: deepCopy(data),
	}
}
