# `protoc-gen-markdown` - Generates markdown from proto files

`protoc-gen-markdown` is a plugin for `protoc` which generates GitHub-flavored markdown
([GFM](https://github.github.com/gfm/)) from proto files.

By default, it will include validation defined by [protovalidate](https://github.com/bufbuild/protovalidate).

It can also generate [Mermaid diagrams](https://mermaid.js.org/) via additional flags.

## Usage

```shell
# Standard usage
protoc -I./example/protos/ \
  --markdown_out=paths=source_relative:./example/output/ \
  models/models.proto \
  service/service.proto

# Mermaid variant
protoc -I./example/protos/ \
  --markdown_out=mermaid=true,paths=source_relative:./example/output/ \
  models/models.proto \
  service/service.proto

# Exclude validation details
protoc -I./example/protos/ \
  --markdown_out=protovalidate=false,paths=source_relative:./example/output/ \
  models/models.proto \
  service/service.proto
```
