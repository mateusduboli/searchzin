package indexer

import (
	"github.com/satori/go.uuid"
)

type Document struct {
	Id   string                 `json:"id"`
	Data map[string]interface{} `json:"data"`
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
		Id:   generateId(),
		Data: deepCopy(data),
	}
}
