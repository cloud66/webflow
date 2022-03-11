# Webflow

## Development
To generate the go output, run the following command:

```bash
protoc webflow.proto --go_out=. --go_opt=paths=source_relative -I . -I ${GOPATH}/src -I ${GOPATH}/src/github.com/envoyproxy/protoc-gen-validate
```

### NOTE
Webflow uses Protobuf validations (https://github.com/envoyproxy/protoc-gen-validate) which needs to be installed before running the conversion command above:

```bash
go get -d github.com/envoyproxy/protoc-gen-validate
```

You also need to clone https://github.com/envoyproxy/protoc-gen-validate in your `$GOPATH` and run `make build` there once as part of installation.