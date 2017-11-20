package indexer

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path"
)

const (
	DATA_FOLDER = "data"
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
}

func GetDocument(id string) (Document, error) {
	dataFolder := dataFolder()
	filename := path.Join(dataFolder, id)

	var doc Document
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		return doc, err
	}
	err = json.Unmarshal(contents, &doc)
	if err != nil {
		return doc, err
	}
	return doc, nil
}

func GetDocuments(ids []string) ([]Document, error) {
	result := []Document{}
	for _, id := range ids {
		document, err := GetDocument(id)
		if err != nil {
			return nil, err
		}
		result = append(result, document)
	}
	return result, nil
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

func dataFolder() string {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}

	dir := path.Dir(ex)
	dataFolder := path.Join(dir, DATA_FOLDER)
	return dataFolder
}
