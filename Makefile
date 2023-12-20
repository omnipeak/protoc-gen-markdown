# Builds protoc-gen-markdown
.PHONY: build
build:
	go build -o protoc-gen-markdown ./main.go

# Builds the docker image
.PHONY: docker-image
docker-image:
	docker build --platform linux/amd64 -t "gchr.io/omnipeak/protoc-gen-markdown:$(shell git describe --tags --abbrev=0 2>/dev/null || echo "0.0.1" | sed "s/^v//")" -t gchr.io/omnipeak/protoc-gen-markdown:latest .

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
