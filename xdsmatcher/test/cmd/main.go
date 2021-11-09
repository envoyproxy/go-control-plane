package main

import (
	"github.com/lyft/xdsmatcher/cmd/matcher/cmd"
	_ "github.com/lyft/xdsmatcher/test"
)

func main() {
	cmd.Execute()
}
