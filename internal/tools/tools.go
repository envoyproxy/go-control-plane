//go:build tools

package tools

import (
	_ "github.com/golangci/golangci-lint/cmd/golangci-lint"
	_ "go.opentelemetry.io/build-tools/multimod"
)
