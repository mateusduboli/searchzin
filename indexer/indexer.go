package indexer

import (
	"fmt"
)

func IndexDocument(document map[string]interface{}) {
	for f := range document {
		switch v := document[f].(type) {
		case string:
			fmt.Printf("Document field [%s] with value [%s]\n", f, v)
		case int:
			fmt.Printf("Document field [%s] with value [%d]\n", f, v)
		}
	}
}
