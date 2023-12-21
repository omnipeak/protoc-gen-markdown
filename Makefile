# Builds protoc-gen-markdown
.PHONY: build
build:
	go build -o protoc-gen-markdown ./main.go

# Builds the docker image
.PHONY: build-docker-image
build-docker-image:
	docker build --platform linux/amd64 -t "ghcr.io/omnipeak/protoc-gen-markdown:v$(shell git describe --tags --abbrev=0 2>/dev/null || echo "0.0.1" | sed "s/^v//")" -t ghcr.io/omnipeak/protoc-gen-markdown:latest .

# Pushes the docker image
.PHONY: push-docker-image
push-docker-image:
	docker push "ghcr.io/omnipeak/protoc-gen-markdown:v$(shell git describe --tags --abbrev=0 2>/dev/null || echo "0.0.1" | sed "s/^v//")"
	docker push ghcr.io/omnipeak/protoc-gen-markdown:latest

# Pushes the image to buf
.PHONY: push-buf-image
push-buf-image:
	npx buf beta registry plugin push \
		--visibility public \
		--image ghcr.io/omnipeak/protoc-gen-markdown:v$(shell git describe --tags --abbrev=0 2>/dev/null || echo "0.0.1" | sed "s/^v//")

# Runs buf generate
.PHONY: generate
generate:
	buf generate ./example/protos

# Runs the unit tests
.PHONY: coverage
coverage:
	@go clean -testcache && \
	GOEXPERIMENT=nocoverageredesign go test -v -coverpkg=./... -coverprofile=.coverage.out ./... && \
	go tool cover -func=.coverage.out && \
	go tool cover -html=.coverage.out -o .coverage.html
