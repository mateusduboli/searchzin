install:
	which dep || brew install dep
	which golint || go get -u github.com/golang/lint/golint
	dep ensure

readme:
	scripts/readme.sh

lint:
	go fmt
	golint

test:
	go test

run:
	go run main.go

release:
	mkdir -p dist
	go build -o dist/searchzin .
