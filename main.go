package main

import (
	"fmt"

	"github.com/envoyproxy/go-control-plane/pkg/version"
)

func main() {
	fmt.Printf("Hello Envoy %s!\n", version.Build.Version)
}
