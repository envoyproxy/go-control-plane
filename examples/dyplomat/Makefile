.PHONY: test
test:
	go test ./... -race
.PHONY: coverage
coverage:
	go test ./... -race -covermode=atomic -coverprofile=coverage.out
coverage_html: coverage
	go tool cover -html=coverage.out
