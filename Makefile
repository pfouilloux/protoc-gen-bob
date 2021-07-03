.PHONY: update-testdata
update-testdata:
	protoc --go_out=./test/testdata/input ./test/testdata/input/orgs.proto

.PHONY: check
check: fmt test lint

.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: test
test:
	go test ./...

.PHONY: lint
lint:
	@go install honnef.co/go/tools/cmd/staticcheck@latest
	@staticcheck -f stylish ./...