TEST_REGEX ?= '.'
VERSION ?= $(shell git show -s --format=%h))
APP_NAME ?= "mateusduboli/searchzin"

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
	go test -v -run ${TEST_REGEX} ./...

run:
	go run main.go

run-dev:
	docker run -p 8080:8080 "${APP_NAME}:dev"

release:
	mkdir -p dist
	env GOOS=linux go build -o dist/searchzin .
	docker build \
		--force-rm \
		--compress \
		--pull \
		--no-cache \
		--tag "${APP_NAME}:${VERSION}" \
		.

release-dev:
	mkdir -p dist
	env GOOS=linux go build -o dist/searchzin .
	docker build \
		--tag "${APP_NAME}:dev" \
		.

clean:
	go clean
	rm -rf dist/
	docker images -q "${APP_NAME}" | xargs docker rmi -f

publish: clean release
	docker push "${APP_NAME}:$(VERSION)"

publish-latest: publish
	docker tag "${APP_NAME}:${VERSION}" "${APP_NAME}:latest"
	docker push "${APP_NAME}:latest"
