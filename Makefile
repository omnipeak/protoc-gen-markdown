# Builds protoc-gen-markdown
.PHONY: build
build:
	go build -o protoc-gen-markdown ./cmd/protoc-gen-markdown/main.go

# Builds the docker image
.PHONY: build-docker-image
build-docker-image:
	docker build --platform linux/amd64 -t "ghcr.io/omnipeak/protoc-gen-markdown:$(shell git describe --tags --abbrev=0 2>/dev/null || echo "v0.0.1")" -t ghcr.io/omnipeak/protoc-gen-markdown:latest .

# Pushes the docker image
.PHONY: push-docker-image
push-docker-image:
	docker push "ghcr.io/omnipeak/protoc-gen-markdown:$(shell git describe --tags --abbrev=0 2>/dev/null || echo "v0.0.1")"
	docker push ghcr.io/omnipeak/protoc-gen-markdown:latest

# Pushes the image to buf
.PHONY: push-buf-image
push-buf-image:
	npx buf beta registry plugin push \
		--visibility public \
		--image ghcr.io/omnipeak/protoc-gen-markdown:$(shell git describe --tags --abbrev=0 2>/dev/null || echo "v0.0.1")

# Runs buf generate
.PHONY: generate
generate:
	@buf generate \
		--path ./example/protos/models.proto \
		--path ./example/protos/service.proto \
		./example/protos

# Runs the unit tests
.PHONY: coverage
coverage:
	@go clean -testcache && \
	GOEXPERIMENT=nocoverageredesign go test -v -coverpkg=./... -coverprofile=.coverage.out ./... && \
	go tool cover -func=.coverage.out && \
	go tool cover -html=.coverage.out -o .coverage.html
