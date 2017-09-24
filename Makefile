test:
	go test

dependencies:
	dep ensure

run: dependencies
	go run main.go

readme:
	scripts/readme.sh
