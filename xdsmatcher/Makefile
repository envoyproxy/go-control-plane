.PHONY: test
test:
	go test ./... -race
.PHONY: coverage
coverage:
	go test ./... -race -covermode=atomic -coverprofile=coverage.out
coverage_html: coverage
	go tool cover -html=coverage.out
.PHONY: proto
proto: test/proto/*
	protoc --proto_path=test/proto --go_out=test/proto --go_opt=paths=source_relative test.proto

.PHONY: tidy
tidy:
	go mod tidy
