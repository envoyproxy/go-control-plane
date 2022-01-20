.PHONY: test
test:
	go test ./... -race
coverage:
	go test ./... -race -covermode=atomic -coverprofile=coverage.out
coverage_html: coverage
	go tool cover -html=coverage.out
proto: test/proto/*
	protoc --proto_path=test/proto --go_out=test/proto --go_opt=paths=source_relative test.proto


