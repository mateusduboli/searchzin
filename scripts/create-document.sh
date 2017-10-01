#!/bin/bash

curl -XPOST -H 'Content-Type: application/json' -d '{"id": 4, "name": "João"}' localhost:8080/v1/documents
