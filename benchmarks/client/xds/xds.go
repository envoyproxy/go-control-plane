package xds

import (
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// Sess holds a grpc connection as well as config options to use during the simulation
type Sess struct {
	Session *grpc.ClientConn
	Opts    Options
}

// NewSession will dial a new benchmarking session with the configured options
func NewSession(url string, opts ...Option) (*Sess, error) {
	var options Options
	for _, o := range opts {
		o(&options)
	}

	conn, err := grpc.Dial(url, grpc.WithBlock(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return &Sess{
		Session: conn,
		Opts:    options,
	}, nil
}

// Simulate will start an xDS stream which provides simulated clients communicating with an xDS server
func (s *Sess) Simulate(target time.Duration) error {
	// Create a loop that will continually do work until the elapsed time as passed
	for timeout := time.After(target); ; {
		select {
		case <-timeout:
			return nil
		default:
			// TODO(alec): implement me
			// Do some work
		}
	}
}
