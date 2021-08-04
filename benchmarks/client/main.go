package main

import (
	"log"
	"os"
	"runtime"
	"time"

	"github.com/envoyproxy/go-control-plane/benchmarks/client/xds"
	"github.com/urfave/cli/v2"
)

func main() {
	// Statically set the max procs
	runtime.GOMAXPROCS(runtime.NumCPU())

	app := &cli.App{
		Name:  "xds-benchmark",
		Usage: "xds benchmarking tool that simulates client workload",
		Commands: []*cli.Command{
			{
				Name:    "run",
				Aliases: []string{"r"},
				Usage:   "run the benchmark with the provided duration",
				Action: func(c *cli.Context) error {
					arg := c.Args().First()
					dur, err := time.ParseDuration(arg)
					if err != nil {
						return err
					}

					sess, err := xds.NewSession("localhost:50000")
					if err != nil {
						return err
					}

					return sess.Simulate(dur)
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
