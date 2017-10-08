package indexer

import (
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"
)

const (
	INDEX_FOLDER  = "data/index"
	KEY_SEPARATOR = ":"
	ID_SEPARATOR  = "\n"
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
		entry := []byte(document.Data[key].(string) + KEY_SEPARATOR + document.Id + ID_SEPARATOR)
		log.Printf("Writting to file [%s] contents [%b]\n", filename, entry)
		if _, err := f.Write(entry); err != nil {
			panic(err)
		}
	}
}

func ListIndices() map[string]map[string][]string {
	indexFolder := indexFolder()
	if err := os.MkdirAll(indexFolder, 0755); err != nil {
		panic(err)
	}

	files, err := ioutil.ReadDir(indexFolder)
	if err != nil {
		panic(err)
	}

	indexes := make(map[string]map[string][]string)

	for _, file := range files {
		if file.IsDir() {
			continue
		}
		field := file.Name()
		filepath := path.Join(indexFolder, field)
		contents, err := ioutil.ReadFile(filepath)
		if err != nil {
			panic(err)
		}
		if indexes[field] == nil {
			indexes[field] = make(map[string][]string)
		}
		entries := strings.Split(string(contents), ID_SEPARATOR)
		for _, entry := range entries {
			if entry == "" {
				continue
			}
			split := strings.Split(entry, KEY_SEPARATOR)
			key, id := split[0], split[1]
			if indexes[field][key] == nil {
				indexes[field][key] = []string{id}
			} else {
				indexes[field][key] = append(indexes[field][key], id)
			}
		}
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
