# Builds protoc-gen-markdown
.PHONY: build
build:
	go build -o protoc-gen-markdown ./main.go

# Builds the docker image
.PHONY: docker-image
docker-image:
	TAG="$(shell git describe --tags --abbrev=0 || echo "v0.0.1")" && \
	docker build --platform linux/amd64 -t gchr.io/omnipeak/protoc-gen-markdown:$(TAG) .
