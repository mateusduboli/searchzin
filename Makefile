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
	go test -v ./...

run:
	go run main.go

release:
	mkdir -p dist
	env GOOS=linux go build -o dist/searchzin .
	docker build \
		--force-rm \
		--compress \
		--pull \
		--no-cache \
		--tag "searchzin:dev" \
		.

clean:
	go clean
	rm -rf dist/
	docker images -q "searchzin" | xargs docker rmi -f

publish: clean release
	$(eval version := $(shell git show -s --format=%h))
	docker tag "searchzin:dev" "mateusduboli/searchzin:$(version)"
	docker push "mateusduboli/searchzin:$(version)"

publish-latest: publish
	docker tag "searchzin" "mateusduboli/searchzin:latest"
	docker push "mateusduboli/searchzin:latest"
