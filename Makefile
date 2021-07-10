.PHONY: update-testdata
update-testdata:
	protoc --go_out=./cmd/testdata/ ./cmd/testdata/clan.proto
	go test ./... -update

.PHONY: check
check: fmt test lint

.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: test
test:
	go test -cover ./...

.PHONY: lint
lint:
	@go install honnef.co/go/tools/cmd/staticcheck@latest
	@staticcheck -f stylish ./...