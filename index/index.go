package index

import (
	"encoding/json"
	. "github.com/mateusduboli/searchzin/document"
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
	DATA_FOLDER   = "data"
)

func IndexDocument(document Document) {
	filename := documentFile(document.Id)
	dataFolder := dataFolder()

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

func GetDocument(id string) (Document, error) {
	filename := documentFile(id)

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

func DeleteDocument(id string) (bool, error) {
	_, err := GetDocument(id)
	if err != nil {
		return false, err
	}
	documentFilename := documentFile(id)

	err = os.Remove(documentFilename)

	refreshIndex()

	return true, nil
}

func refreshIndex() {
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

func documentFile(id string) string {
	dataFolder := dataFolder()
	return path.Join(dataFolder, id)
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

func GetIndex(field string) map[string][]string {
	indexes := ListIndices()
	return indexes[field]
}

func GetIndexTerm(field string, term string) []string {
	indexes := ListIndices()
	return indexes[field][term]
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
