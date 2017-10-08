package indexer

import (
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"
)

const (
	INDEX_FOLDER = "data/index"
)

func IndexDocumentFields(document Document) {
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

func indexFolder() string {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}

	dir := path.Dir(ex)
	indexFolder := path.Join(dir, INDEX_FOLDER)
	return indexFolder
}
