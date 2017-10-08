package indexer

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"
)

const (
	DATA_FOLDER  = "data"
	INDEX_FOLDER = "data/index"
)

func IndexDocument(document Document) {
	log.Printf("Indexing document [%s]\n", document)
	dataFolder := dataFolder()
	log.Printf("Data directory [%s]\n", dataFolder)
	filename := path.Join(dataFolder, document.Id)
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
	log.Printf("Writting to file [%s] contents [%b]\n", filename, b)
	if _, err := f.Write(b); err != nil {
		panic(err)
	}
	indexDocumentFields(document)
}

func indexDocumentFields(document Document) {
	log.Printf("Indexing document fields [%s]\n", document)
	indexFolder := indexFolder()
	log.Printf("Index directory [%s]\n", indexFolder)

	if err := os.MkdirAll(indexFolder, 0755); err != nil {
		panic(err)
	}

	for key := range document.Data {
		filename := path.Join(indexFolder, key)
		log.Printf("Index file [%s]\n", filename)
		f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
		if err != nil {
			panic(err)
		}
		id := []byte(document.Id)
		log.Printf("Writting to file [%s] contents [%b]\n", filename, id)
		if _, err := f.Write(id); err != nil {
			panic(err)
		}
	}
}

func ListDocuments() []Document {
	dataFolder := dataFolder()
	if err := os.MkdirAll(dataFolder, 0755); err != nil {
		panic(err)
	}

	files, err := ioutil.ReadDir(dataFolder)
	if err != nil {
		panic(err)
	}

	var docs []Document

	for _, file := range files {
		if file.IsDir() {
			continue
		}
		var doc Document
		filepath := path.Join(dataFolder, file.Name())
		contents, err := ioutil.ReadFile(filepath)
		if err != nil {
			panic(err)
		}
		err = json.Unmarshal(contents, &doc)
		if err != nil {
			panic(err)
		}
		docs = append(docs, doc)
	}

	return docs
}

func ListIndices() map[string][]string {
	indexFolder := indexFolder()
	if err := os.MkdirAll(indexFolder, 0755); err != nil {
		panic(err)
	}

	files, err := ioutil.ReadDir(indexFolder)
	if err != nil {
		panic(err)
	}

	indexes := make(map[string][]string)

	for _, file := range files {
		if file.IsDir() {
			continue
		}
		filepath := path.Join(indexFolder, file.Name())
		contents, err := ioutil.ReadFile(filepath)
		documentIds := strings.Split(string(contents), "\n")
		if err != nil {
			panic(err)
		}
		indexes[file.Name()] = documentIds
	}

	return indexes
}

func dataFolder() string {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}

	dir := path.Dir(ex)
	dataFolder := path.Join(dir, DATA_FOLDER)
	return dataFolder
}

func indexFolder() string {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}

	dir := path.Dir(ex)
	indexFolder := path.Join(dir, INDEX_FOLDER)
	return indexFolder
}
