TEST_REGEX ?= '.'
VERSION ?= $(shell cat VERSION))
APP_NAME ?= "mateusduboli/searchzin"

install:
	which dep || brew install dep
	which golint || go get -u github.com/golang/lint/golint
	mkdir -p .git/hooks
	ln -s -f ${PWD}/scripts/pre-commit .git/hooks
	dep ensure

readme:
	scripts/readme.py

lint:
	go fmt ./cmd/* ./pkg/*
	golint ./cmd/* ./pkg/*

test:
	go test -v -run ${TEST_REGEX} ./...

build: build-searchzin

build-searchzin:
	mkdir -p dist/web
	env GOOS=linux go build -o dist/searchzin cmd/searchzin/main.go
	cp -R web/templates dist/web/

run-searchzin:
	go run cmd/searchzin/main.go

run: run-searchzin

run-dev: build
	docker build \
		--tag "${APP_NAME}:dev" \
		.
	docker run \
		-p 8080:8080 \
		-v "${PWD}/dist:/opt/searchzin" \
		"${APP_NAME}:dev"

release: clean build
	docker build \
		--force-rm \
		--compress \
		--pull \
		--no-cache \
		--tag "${APP_NAME}:${VERSION}" \
		.

clean: clean-dist clean-docker

clean-dist:
	go clean
	rm -rf dist/

clean-docker:
	docker images -q "${APP_NAME}" | xargs docker rmi -f

publish: clean release
	docker push "${APP_NAME}:$(VERSION)"

publish-latest: publish
	docker tag "${APP_NAME}:${VERSION}" "${APP_NAME}:latest"
	docker push "${APP_NAME}:latest"
