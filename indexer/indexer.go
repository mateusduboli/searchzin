package indexer

import (
	"encoding/json"
	"log"
	"os"
	"path"
)

const (
	DATA_FOLDER = "data"
)

func IndexDocument(document Document) {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}

	dir := path.Dir(ex)
	dataFolder := path.Join(dir, DATA_FOLDER)
	log.Printf("Data directory [%s]\n", dataFolder)
	filename := path.Join(dataFolder, document.id)
	log.Printf("Document file [%s]\n", filename)

	if err := os.MkdirAll(dataFolder, 0755); err != nil {
		panic(err)
	}

	f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	b, err := json.Marshal(document)
	if err != nil {
		panic(err)
	}
	if _, err := f.Write(b); err != nil {
		panic(err)
	}
}

func ListDocuments() []Document {
	return make([]Document, 0)
}
